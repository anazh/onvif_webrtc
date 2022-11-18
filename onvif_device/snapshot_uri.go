package onvif_device

import (
	"sync"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
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

func DoShot(ip string, port int, localFile string) {
	url := getSnapShotUri(ip, port)
	if url != "" {
		data := ghttp.GetBytes(url)
		gfile.PutBytes(localFile, data)
	}
}
