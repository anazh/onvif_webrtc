package onvif_device

import (
	"github.com/use-go/onvif"
)

func getTestConfig() *onvif.Device {
	v := VideoConfig{
		IP:       "b.nps.kyunmao.com",
		Port:     82,
		UserName: "admin",
		Password: "mdzh12345",
	}
	dev, err := InitIpc(v)
	if err != nil {
		panic(err)
	}
	return dev
}
