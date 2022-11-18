// onvif.go
package onvif_device

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	sdkMedia "github.com/use-go/onvif/sdk/media"

	goonvif "github.com/use-go/onvif"
	"github.com/use-go/onvif/media"
)

func InitIpc(config VideoConfig) error {
	ctx := context.Background()
	dev, err := goonvif.NewDevice(goonvif.DeviceParams{
		Xaddr:      fmt.Sprintf("%s:%d", config.IP, config.Port),
		Username:   config.UserName,
		Password:   config.Password,
		HttpClient: new(http.Client),
	})
	if err != nil {
		return err
	}
	// register snap shor uri
	resp, err := sdkMedia.Call_GetSnapshotUri(ctx, dev, media.GetSnapshotUri{})
	if err != nil {
		return err
	}
	u := SnapShotUri{
		Ip:   config.IP,
		Port: config.Port,
		Uri:  strings.ReplaceAll(string(resp.MediaUri.Uri), "://", fmt.Sprintf("://%s:%d@", config.IP, config.Port)),
	}
	registerSnapShotUri(u)
	// register pzt
	return nil
}
