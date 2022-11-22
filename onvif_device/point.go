package onvif_device

import (
	"encoding/xml"
	"io"

	"github.com/use-go/onvif"
	"github.com/use-go/onvif/ptz"
	"github.com/use-go/onvif/xsd"
	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

// 预置点列表专用的json解析
type PTZPresetList struct {
	Token       string `xml:"token,attr"`
	Name        string
	PTZPosition PTZVectorList
}

type PTZVectorList struct {
	PanTilt onvif2.Vector2D `xml:"PanTilt"`
	Zoom    onvif2.Vector1D `xml:"Zoom"`
}

// 获取设备的所有的预置点
func DeviceAllPoints(dev *onvif.Device, profileName string) ([]PTZPresetList, error) {
	resp, err := dev.CallMethod(ptz.GetPresets{ //获取预置点
		ProfileToken: onvif2.ReferenceToken(profileName),
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	type Envelope struct {
		Body struct {
			GetPresetsResponse struct {
				Preset []PTZPresetList
			}
		}
	}
	var reply Envelope
	err = xml.Unmarshal(data, &reply)
	return reply.Body.GetPresetsResponse.Preset, err
}

// 设置预置点为当前点位
// pointName 是点位的名称必须有
// pointToken 可以为空，为空的话会顺位设置点位
// pointToken 不为空，表示设置一个已经存在的点位，这个token务必要存在 DeviceAllPoints 中
func SetPoint(dev *onvif.Device, profileName, pointName, pointToken string) (string, error) {
	p := ptz.SetPreset{
		ProfileToken: onvif2.ReferenceToken(profileName), //代表哪路流
		PresetName:   xsd.String(pointName),
	}
	if pointToken != "" {
		p.PresetToken = onvif2.ReferenceToken(pointToken)
	}
	resp, err := dev.CallMethod(p)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type Envelope struct {
		Body struct {
			SetPresetResponse struct {
				PresetToken string
			}
		}
	}
	var reply Envelope
	err = xml.Unmarshal(data, &reply)
	if err != nil {
		return "", err
	}
	return reply.Body.SetPresetResponse.PresetToken, nil
}

// 移动到某个预置点
func GoToPoint(dev *onvif.Device, profileName, pointToken string) error {
	_, err := dev.CallMethod(ptz.GotoPreset{ //去到预置点
		ProfileToken: onvif2.ReferenceToken(profileName), //代表哪路流
		PresetToken:  onvif2.ReferenceToken(pointToken),
	})
	return err
}

// 删除预置点 => 不建议使用删除操作
// 删除此预置点后，不能通过 传入 pointToken，来设置此token
// 所以建议，不要使用删除操作
// 所以建议，不要使用删除操作
// 所以建议，不要使用删除操作
func DelPoint(dev *onvif.Device, profileName, pointToken string) error {
	resp, err := dev.CallMethod(ptz.RemovePreset{
		ProfileToken: onvif2.ReferenceToken(profileName), //代表哪路流
		PresetToken:  onvif2.ReferenceToken(pointToken),
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	type Envelope struct {
		Body struct {
			RemovePresetResponse struct {
			}
		}
	}
	var reply Envelope
	err = xml.Unmarshal(data, &reply)
	if err != nil {
		return err
	}
	return nil
}
