package onvif_device

import "fmt"

type VideoConfig struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	UserName string `json:"user_name"` //to be use for rtsp and onvif
	Password string `json:"password"`
}

// snapshoturi cache
type SnapShotUri struct {
	Ip   string
	Port int
	Uri  string
}

func deviceKey(ip string, port int) string {
	return fmt.Sprintf("%s%d", ip, port)
}
