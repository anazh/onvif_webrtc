package main

import "github.com/anazh/onvif_webrtc/onvif_device"

// main
func main() {
	v := onvif_device.VideoConfig{
		IP:       "192.168.0.64",
		Port:     80,
		UserName: "admin",
		Password: "mdzh12345",
	}
	onvif_device.InitIpc(v)
	onvif_device.DoShot(v.IP, v.Port, "y.jpeg")
}
