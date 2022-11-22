package onvif_device

import "testing"

func TestDeviceAllPoints(t *testing.T) {
	v := VideoConfig{
		IP:       "b.nps.kyunmao.com",
		Port:     82,
		UserName: "admin",
		Password: "mdzh12345",
	}
	dev, err := InitIpc(v)
	if err != nil {
		t.Error(err)
	}
	points, err := DeviceAllPoints(dev, "Profile_1")
	if err != nil {
		t.Error(err)
	}
	if len(points) != 300 && points[0].Token != "1" {
		t.Fail()
	}
}

func TestDelPoint(t *testing.T) {
	v := VideoConfig{
		IP:       "b.nps.kyunmao.com",
		Port:     82,
		UserName: "admin",
		Password: "mdzh12345",
	}
	dev, err := InitIpc(v)
	if err != nil {
		t.Error(err)
	}
	err = DelPoint(dev, "Profile_1", "12")
	if err != nil {
		t.Error(err)
	}
}

func TestSetPoint(t *testing.T) {
	v := VideoConfig{
		IP:       "b.nps.kyunmao.com",
		Port:     82,
		UserName: "admin",
		Password: "mdzh12345",
	}
	dev, err := InitIpc(v)
	if err != nil {
		t.Error(err)
	}
	_, err = SetPoint(dev, "Profile_1", "12", "12")
	if err != nil {
		t.Error(err)
	}
}

func TestGotoPoint(t *testing.T) {
	v := VideoConfig{
		IP:       "b.nps.kyunmao.com",
		Port:     82,
		UserName: "admin",
		Password: "mdzh12345",
	}
	dev, err := InitIpc(v)
	if err != nil {
		t.Error(err)
	}
	err = GoToPoint(dev, "Profile_1", "12")
	if err != nil {
		t.Error(err)
	}
}
