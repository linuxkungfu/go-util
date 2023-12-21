package util

type MediaServerInfo struct {
	ServerId string      `json:"server_id"`
	Address  NetworkAddr `json:"address"`
}
