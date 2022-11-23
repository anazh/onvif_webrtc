package onvif_device

import "testing"

func TestDeviceAllPoints(t *testing.T) {
	dev := getTestConfig()
	points, err := DeviceAllPoints(dev, "Profile_1")
	if err != nil {
		t.Error(err)
	}
	if len(points) != 300 && points[0].Token != "1" {
		t.Fail()
	}
}

func TestDelPoint(t *testing.T) {
	dev := getTestConfig()
	err := DelPoint(dev, "Profile_1", "12")
	if err != nil {
		t.Error(err)
	}
}

func TestSetPoint(t *testing.T) {
	dev := getTestConfig()
	_, err := SetPoint(dev, "Profile_1", "12", "12")
	if err != nil {
		t.Error(err)
	}
}

func TestGotoPoint(t *testing.T) {
	dev := getTestConfig()
	err := GoToPoint(dev, "Profile_1", "12")
	if err != nil {
		t.Error(err)
	}
}
