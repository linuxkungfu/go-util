package util

type HeartbeatInfo struct {
	Topic        string  `json:"topic"`
	Timestamp    int64   `json:"timestamp"`
	DeviceNumber int     `json:"device_number"`
	CpuLoad      float64 `json:"cpu_load"`
	CpuUsage     float64 `json:"cpu_usage"`
}
type NotifyMessage struct {
	TestId    uint32      `json:"testId"`
	PeerId    string      `json:"peerId"`
	EventType string      `json:"eventType"`
	Data      interface{} `json:"data"`
}
