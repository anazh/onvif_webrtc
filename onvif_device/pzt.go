package onvif_device

import (
	"encoding/xml"
	"io"

	"github.com/use-go/onvif"
	"github.com/use-go/onvif/ptz"
	onvif2 "github.com/use-go/onvif/xsd/onvif"
)

// ptz controller

type move struct {
	dev         *onvif.Device
	profileName string
	x           float64
	y           float64
	z           float64
}

// 速度转换
// 1-7 => 0.1-0.7
func ptzSpeed(speed int) float64 {
	if speed < 1 {
		speed = 1
	} else if speed > 7 {
		speed = 7
	}
	return float64(speed) / 10
}

// 目前位置
func PztStatus(dev *onvif.Device, profileName string) (ptz.GetStatusResponse, error) {
	resp, err := dev.CallMethod(ptz.GetStatus{ //相对移动
		ProfileToken: onvif2.ReferenceToken(profileName),
	})
	if err != nil {
		return ptz.GetStatusResponse{}, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ptz.GetStatusResponse{}, err
	}
	type Envelope struct {
		Body struct {
			GetStatusResponse ptz.GetStatusResponse
		}
	}
	var reply Envelope
	err = xml.Unmarshal(data, &reply)
	return reply.Body.GetStatusResponse, err
}

// 停止移动
func StopMove(dev *onvif.Device, profileName string) error {
	resp, err := dev.CallMethod(ptz.Stop{ //相对移动
		ProfileToken: onvif2.ReferenceToken(profileName),
		PanTilt:      false,
		Zoom:         false,
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	return err
}

// 当前位置

// ----------------------------------------------------
// 持续移动
func continueMove(m move) error {
	resp, err := m.dev.CallMethod(ptz.ContinuousMove{ //相对移动
		ProfileToken: onvif2.ReferenceToken(m.profileName),
		Velocity: onvif2.PTZSpeed{
			PanTilt: onvif2.Vector2D{
				X: m.x,
				Y: m.y,
			},
			Zoom: onvif2.Vector1D{
				X: m.z,
			},
		},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	// gfile.PutBytes("relativeMove", data)
	return err
}

func ContinuousMoveUp(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           0,
		y:           ptzSpeed(speed),
		z:           0,
	}
	return continueMove(m)
}

func ContinuousMoveLow(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           0,
		y:           -ptzSpeed(speed),
		z:           0,
	}
	return continueMove(m)
}

func ContinuousMoveRight(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           ptzSpeed(speed),
		y:           0,
		z:           0,
	}
	return continueMove(m)
}

func ContinuousMoveLeft(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           -ptzSpeed(speed),
		y:           0,
		z:           0,
	}
	return continueMove(m)
}
func ContinuousMoveClose(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           0,
		y:           0,
		z:           ptzSpeed(speed),
	}
	return continueMove(m)
}
func ContinuousMoveFar(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           0,
		y:           0,
		z:           -ptzSpeed(speed),
	}
	return continueMove(m)
}

// 在目前的位置上，相对移动
func relativeMove(m move) error {
	resp, err := m.dev.CallMethod(ptz.RelativeMove{ //相对移动
		ProfileToken: onvif2.ReferenceToken(m.profileName),
		Translation: onvif2.PTZVector{
			PanTilt: onvif2.Vector2D{
				X:     m.x,
				Y:     m.y,
				Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/TranslationGenericSpace",
			},
			Zoom: onvif2.Vector1D{
				X:     m.z,
				Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/TranslationGenericSpace",
			},
		},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	// gfile.PutBytes("relativeMove", data)
	return err
}

// 向上移动
func RelativeMoveUp(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           0,
		y:           ptzSpeed(speed),
		z:           0,
	}
	return relativeMove(m)
}

// 向下移动
func RelativeMoveLow(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           0,
		y:           -ptzSpeed(speed),
		z:           0,
	}
	return relativeMove(m)
}

// 向右移动
func RelativeMoveRight(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           ptzSpeed(speed),
		y:           0,
		z:           0,
	}
	return relativeMove(m)
}

// 向左移动
func RelativeMoveLeft(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           -ptzSpeed(speed),
		y:           0,
		z:           0,
	}
	return relativeMove(m)
}

// 放大焦距
func RelativeMoveClose(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           0,
		y:           0,
		z:           ptzSpeed(speed),
	}
	return relativeMove(m)
}

// 缩小焦距
func RelativeMoveFar(dev *onvif.Device, profileName string, speed int) error {
	m := move{
		dev:         dev,
		profileName: profileName,
		x:           0,
		y:           0,
		z:           -ptzSpeed(speed),
	}
	return relativeMove(m)
}
