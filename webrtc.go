package util

type DataPackType int

const (
	UnknownPack DataPackType = 0
	StunPack    DataPackType = 1
	DtlsPack    DataPackType = 2
	RtpPack     DataPackType = 3
	RtcpPack    DataPackType = 4
)
const (
	RtpVersion         byte = 2
	RtpHeaderSize      byte = 12
	RtpversionShift    byte = 6
	RtpversionMask     byte = 0x3
	RtcpCommHeaderSize byte = 4
)

func IsRtp(data []byte) bool {
	if len(data) < int(RtpHeaderSize) {
		return false
	}
	// https://tools.ietf.org/html/draft-ietf-avtcore-rfc5764-mux-fixes
	if data[0] <= 127 || data[0] >= 192 {
		return false
	}
	version := data[0] >> RtpversionShift & RtpversionMask
	return version == RtpVersion
}
func IsRtcp(data []byte) bool {
	if len(data) < int(RtcpCommHeaderSize) {
		return false
	}
	if data[0] <= 127 || data[0] >= 192 {
		return false
	}
	version := data[0] >> RtpversionShift & RtpversionMask
	if version != RtpVersion {
		return false
	}
	if data[1] < 192 || data[1] > 223 {
		return false
	}
	return true
}
