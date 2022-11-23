package onvif_device

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// take this function to init all ipc for onvif
func StartOnvif(ipcs []VideoConfig) {
	for _, v := range ipcs {
		dev, err := InitIpc(v)
		if err != nil {
			g.Log().Warning(context.Background(), "StartOnvif", err)
			continue
		}
		InitShotUri(dev, v)
	}
}
