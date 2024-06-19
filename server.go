package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	utilString "github.com/linuxkungfu/go-util/string"
	logger "github.com/sirupsen/logrus"
	"github.com/twmb/murmur3"
)

var (
	ConnId            int    = 0
	serverId          string = ""
	serverName        string = ""
	timezoneOffsetStr string = ""
	timezoneOffset    int    = 0
)

// 需要重新设计
type NetworkConn struct {
	id                string
	ConnProto         string
	LocalIp           string
	LocalPort         int
	LocalAddr         net.Addr
	RemoteIp          string
	RemotePort        int
	RemoteAddr        net.Addr
	ReadBytes         int64
	WriteBytes        int64
	FailedCount       int64
	IsBind            bool
	conn              interface{}
	heartbeatTS       time.Time
	heartbeatInterval time.Duration
	owner             interface{}
	data              interface{}
}

type NetworkConnOwner interface {
	OnConnDead(conn *NetworkConn)
}

type ServerListener interface {
	OnAcceptConn(conn *NetworkConn)
	OnReceiveData(data []byte, conn *NetworkConn)
	OnClosed(ls interface{})
}

type NetworkServer struct {
	id           string
	running      bool
	listenAddr   NetworkAddr
	addrs        sync.Map
	Listener     ServerListener
	listenServer interface{}
}

// CreateServerId 生成服务器id hostname+ip+port
func CreateServerId(idString string) string {
	hostname, _ := os.Hostname()
	serverId = fmt.Sprintf("%d", murmur3.Sum64([]byte(fmt.Sprintf("%s_%s", hostname, idString))))
	return serverId
}

func generateId() int {
	ConnId++
	return ConnId
}

func generateStrId() string {
	return fmt.Sprintf("%d", generateId())
}

func (conn NetworkConn) Id() string {
	return conn.id
}

func (conn *NetworkConn) UpdateHeatbeat(ts time.Time) {
	conn.heartbeatTS = ts
}

func (conn *NetworkConn) SetHeatbeatInterval(dur time.Duration) {
	conn.heartbeatInterval = dur
}

func (conn *NetworkConn) IsDead(checkTime time.Time) bool {
	// logger.Debugf("[networkConn][%s]NetworkConn IsDead tm:%s interval:%f", conn.id, conn.heartbeatTS.String(), conn.heartbeatInterval.Seconds())
	return conn.heartbeatTS.Add(conn.heartbeatInterval).Before(checkTime)
}

func (conn *NetworkConn) HeartbeatInterval() time.Duration {
	return time.Since(conn.heartbeatTS)
}

func (conn *NetworkConn) OnSendData(data []byte) {
	conn.Write(data)
}

func (conn *NetworkConn) OnClosed(interface{}) {
	logger.Infof("[networkConn][%s]OnClosed", conn.id)
	conn.Close()
}

func (conn *NetworkConn) Close() {
	logger.Infof("[networkConn][%s]local addr %s:%d remote addr %s:%d close read bytes:%s write bytes:%s, failed count:%d", conn.id, conn.LocalIp, conn.LocalPort, conn.RemoteIp, conn.RemotePort, utilString.FormatFileSize(conn.ReadBytes), utilString.FormatFileSize(conn.WriteBytes), conn.FailedCount)
	if strings.ToLower(conn.ConnProto) == "udp" {
		conn.closeUdpConn()
	} else if strings.ToLower(conn.ConnProto) == "tcp" {
		conn.closeTcpConn()
	}
	conn.LocalAddr = nil
	conn.RemoteAddr = nil
	conn.conn = nil
}

func (conn *NetworkConn) RegisterConnOwner(owner interface{}) {
	conn.owner = owner
}

func (conn *NetworkConn) closeUdpConn() {
	if conn.conn != nil {
		conn.conn = nil
	}
}

func (conn *NetworkConn) closeTcpConn() {
	conn.conn.(*net.TCPConn).Close()
}

func (conn *NetworkConn) LocalAddrStr() string {
	return fmt.Sprintf("%s:%d", conn.LocalIp, conn.LocalPort)
}

func (conn *NetworkConn) RemoteAddrStr() string {
	return fmt.Sprintf("%s:%d", conn.RemoteIp, conn.RemotePort)
}

func (conn *NetworkConn) Write(data []byte) bool {
	switch conn.ConnProto {
	case "udp":
		return conn.WriteUdp(data)
	case "tcp":
		return conn.WriteTcp(data)
	default:
		logger.Warnf("[networkConn][%s]NetworkConn Write unknown proto:%v", conn.id, conn.ConnProto)
		return false
	}
}

func (conn *NetworkConn) WriteUdp(data []byte) bool {
	if conn.conn != nil {
		bytes, err := conn.conn.(*net.UDPConn).WriteTo(data, conn.RemoteAddr)
		if err != nil {
			logger.Warnf("[networkConn][%s]WriteUdp addr:%s, error:%s", conn.id, conn.RemoteAddr.String(), err.Error())
			conn.FailedCount++
			return false
		}
		conn.WriteBytes += int64(bytes)
		return err == nil
	}
	return false
}

func (conn *NetworkConn) WriteTcp(data []byte) bool {
	if conn.conn != nil {
		bytes, err := conn.conn.(*net.TCPConn).Write(data)
		if err != nil {
			logger.Warnf("[networkConn][%s]WriteTcp addr:%s, error:%s", conn.id, conn.RemoteAddr.String(), err.Error())
			conn.FailedCount++
			return false
		}
		conn.WriteBytes += int64(bytes)
		return err == nil
	}
	return false
}

func (conn *NetworkConn) Dead() {
	if conn.owner != nil {
		conn.owner.(NetworkConnOwner).OnConnDead(conn)
	}
}

func (conn *NetworkConn) SetData(intf interface{}) {
	conn.data = intf
}

func (conn *NetworkConn) GetData() interface{} {
	return conn.data
}

func (srv *NetworkServer) Start(addr NetworkAddr) {
	srv.listenAddr = addr
	switch strings.ToLower(addr.Protocol) {
	case "udp":
		srv.startUdp()
	case "":
		srv.startTcp()
	default:
		logger.Warnf("[networkServer][%s]NetworkServer Start unknown protocol:%v", srv.id, addr.Protocol)
	}
}

func (srv *NetworkServer) Id() string {
	return srv.id
}

func (srv *NetworkServer) IsRunning() bool {
	return srv.running
}

func (srv *NetworkServer) SetId(id string) {
	srv.id = id
}

func (srv *NetworkServer) startUdp() {
	srv.running = true
	ip := "0.0.0.0"
	addr, ex := net.ResolveUDPAddr("udp", fmt.Sprintf("%v:%v", ip, srv.listenAddr.Port))
	if ex != nil {
		// panic(fmt.Sprintf("解析IP地址异常:%s", e))
		logger.Warnf("[networkServer][%s]NetworkServer resolve raw udp address:%s  error:%v", srv.id, fmt.Sprintf("%s:%d", ip, srv.listenAddr.Port), ex)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		logger.Warnf("[networkServer][%s]NetworkServer listen udp:%s failed:%v", srv.id, fmt.Sprintf("%s:%d", ip, srv.listenAddr.Port), err)
		return
	}
	logger.Infof("[networkServer][%s]NetworkServer is started at:%s", srv.id, fmt.Sprintf("%s://%s:%d", srv.listenAddr.Protocol, ip, srv.listenAddr.Port))
	srv.listenServer = conn
	go srv.handleConn(conn)
}

func (srv *NetworkServer) startTcp() {
	logger.Warnf("[networkServer][%s]NetworkServer startTcp not implemented", srv.id)
}

func (srv *NetworkServer) Stop() {
	srv.running = false
	switch strings.ToLower(srv.listenAddr.Protocol) {
	case "udp":
		srv.stopUdp()
	case "":
		srv.stopTcp()
	default:
		logger.Warnf("[networkServer][%s]NetworkServer Start unknown protocol:%v", srv.id, srv.listenAddr.Protocol)
	}
	srv.addrs = sync.Map{}
	srv.Listener = nil
}

func (srv *NetworkServer) stopUdp() {
	if srv.listenServer != nil {
		srv.listenServer.(*net.UDPConn).Close()
		srv.listenServer = nil
	}

}

func (srv *NetworkServer) stopTcp() {

}

func (srv *NetworkServer) handleConn(conn *net.UDPConn) {
	defer func() {
		conn.Close()
		if r := recover(); r != nil {
			logger.Errorf("[networkServer][%s]NetworkServer handleConn panic:%v", srv.id, r)
			debug.PrintStack()
			srv.Start(srv.listenAddr)
		}
	}()
	// srv.conn = conn
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	for srv.running {
		data := make([]byte, 1500)
		pos, addr, err := conn.ReadFromUDP(data)
		if !srv.running {
			break
		}
		if err != nil {
			if e, ok := err.(net.Error); ok && e.Timeout() {
				conn.SetReadDeadline(time.Now().Add(10 * time.Second))
				continue
			}
			if srv.running {
				if errors.Is(err, net.ErrClosed) {
					logger.Infof("[networkServer][%s]NetworkServer handleConn addr:%s has been closed", srv.id, srv.listenAddr.String())
				} else {
					logger.Warnf("[networkServer][%s]NetworkServer handleConn addr:%s err:%s", srv.id, srv.listenAddr.String(), err.Error())
				}
			}
			break
		}
		connIf, ok := srv.addrs.Load(addr.String())
		var networkConn *NetworkConn
		if !ok {
			remoteAddr := &NetworkAddr{}
			remoteAddr.ParseString(fmt.Sprintf("udp://%s", conn.LocalAddr().String()))
			networkConn = &NetworkConn{
				id:         generateStrId(),
				ConnProto:  "udp",
				LocalIp:    srv.listenAddr.Ip,
				LocalPort:  srv.listenAddr.Port,
				LocalAddr:  conn.LocalAddr(),
				RemoteIp:   addr.IP.String(),
				RemotePort: addr.Port,
				RemoteAddr: addr,
				ReadBytes:  int64(pos),
				IsBind:     false,
				conn:       conn,
			}
			logger.Infof("[networkServer][%s]accept a new connect, id: %s, LocalAddr: %s, RemoteAddr: %s", srv.id, networkConn.id, networkConn.LocalAddr.String(), networkConn.RemoteAddr.String())
			if srv.Listener != nil {
				srv.Listener.OnAcceptConn(networkConn)
			}

			srv.addrs.Store(addr.String(), networkConn)
		} else {
			networkConn = connIf.(*NetworkConn)
		}
		networkConn.ReadBytes += int64(pos)
		if srv.Listener != nil {
			srv.Listener.OnReceiveData(data[:pos], networkConn)
		}
	}
	logger.Infof("[networkServer][%s]NetworkServer:%s exit", srv.id, srv.listenAddr.String())
}

func (srv *NetworkServer) CloseConn(networkConn *NetworkConn) {
	srv.RemoveConn(networkConn)
	networkConn.Close()
}

func (srv *NetworkServer) RemoveConn(networkConn *NetworkConn) bool {
	if networkConn.RemoteAddr != nil {
		srv.addrs.Delete(networkConn.RemoteAddr.String())
		return true
	}
	return false
}

func SetServerInfo(name string) {
	serverName = name
	cur := time.Now()
	_, timezoneOffset = cur.Local().Zone()
	timezoneOffset = timezoneOffset / 3600
	if timezoneOffset > 0 {
		timezoneOffsetStr = fmt.Sprintf("T+%d", timezoneOffset)
	} else {
		timezoneOffsetStr = fmt.Sprintf("T%d", timezoneOffset)
	}
}

func GetServerId() string {
	return serverId
}

func PasswordPlainToMd5(userId uint64, password string) string {
	h := md5.New()
	io.WriteString(h, "kungapp")
	io.WriteString(h, password)
	io.WriteString(h, fmt.Sprintf("%d", userId))

	return fmt.Sprintf("%x", h.Sum(nil))

}
