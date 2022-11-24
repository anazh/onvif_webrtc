package onvif_device

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

// get img by onvif
// var deviceImages map[string]SnapShotUri{}
var deviceSnapShotUri sync.Map

func registerSnapShotUri(u SnapShotUri) {
	deviceSnapShotUri.Store(deviceKey(u.Ip, u.Port), u)
}

func getSnapShotUri(ip string, port int) string {
	v, ok := deviceSnapShotUri.Load(deviceKey(ip, port))
	if ok {
		return v.(SnapShotUri).Uri
	}
	return ""
}

// when this return "" , you need to init shot url
func DoShot(ip string, port int, localFile string) string {
	url := getSnapShotUri(ip, port)
	if url != "" {
		data := g.Client().SetTimeout(5*time.Second).GetBytes(context.Background(), url)
		gfile.PutBytes(localFile, data)
	}
	return url
}
