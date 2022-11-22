package main

import (
	"context"

	"github.com/anazh/onvif_webrtc/onvif_device"
	"github.com/gogf/gf/v2/frame/g"
)

// main
func main() {
	v := onvif_device.VideoConfig{
		IP:       "b.nps.kyunmao.com",
		Port:     82,
		UserName: "admin",
		Password: "mdzh12345",
	}
	dev, err := onvif_device.InitIpc(v)
	if err != nil {
		g.Log().Warning(context.Background(), err)
	}
	vc, err := onvif_device.DeviceAllPoints(dev, "Profile_1")
	if err != nil {
		g.Log().Warning(context.Background(), err)
	}
	g.Log().Debug(context.Background(), vc)
}
