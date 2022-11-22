// onvif.go
package onvif_device

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	sdkMedia "github.com/use-go/onvif/sdk/media"

	goonvif "github.com/use-go/onvif"
	"github.com/use-go/onvif/media"
)

func InitIpc(config VideoConfig) (*goonvif.Device, error) {
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      fmt.Sprintf("%s:%d", config.IP, config.Port),
		Username:   config.UserName,
		Password:   config.Password,
		HttpClient: new(http.Client),
	})
	if err != nil {
		return nil, err
	}
	// register pzt
	return dev, nil
}

func InitShotUri(dev *goonvif.Device, config VideoConfig) error {
	ctx := context.Background()
	resp, err := sdkMedia.Call_GetSnapshotUri(ctx, dev, media.GetSnapshotUri{})
	if err != nil {
		return err
	}
	u := SnapShotUri{
		Ip:   config.IP,
		Port: config.Port,
		Uri:  strings.ReplaceAll(string(resp.MediaUri.Uri), "://", fmt.Sprintf("://%s:%s@", config.UserName, config.Password)),
	}
	g.Log().Debug(context.Background(), "shoturi", u)
	registerSnapShotUri(u)
	return nil
}
