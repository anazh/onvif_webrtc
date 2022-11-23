package onvif_device

import (
	"fmt"
	"testing"
	"time"
)

func TestRelativeMoveUp(t *testing.T) {
	dev := getTestConfig()
	err := RelativeMoveUp(dev.Device, "Profile_1", 1)
	if err != nil {
		t.Error(err)
	}
}
func TestRelativeMoveLow(t *testing.T) {
	dev := getTestConfig()
	err := RelativeMoveLow(dev.Device, "Profile_1", 1)
	if err != nil {
		t.Error(err)
	}
}

func TestRelativeMoveRight(t *testing.T) {
	dev := getTestConfig()
	err := RelativeMoveRight(dev.Device, "Profile_1", 1)
	if err != nil {
		t.Error(err)
	}
}
func TestRelativeMoveClose(t *testing.T) {
	dev := getTestConfig()
	err := RelativeMoveClose(dev.Device, "Profile_1", 1)
	if err != nil {
		t.Error(err)
	}
}
func TestRelativeMoveFar(t *testing.T) {
	dev := getTestConfig()
	err := RelativeMoveFar(dev.Device, "Profile_1", 1)
	if err != nil {
		t.Error(err)
	}
}

func TestContinuMove(t *testing.T) {
	dev := getTestConfig()
	ContinuousMoveRight(dev.Device, "Profile_1", 1)
	time.Sleep(3 * time.Second)
	ContinuousMoveLeft(dev.Device, "Profile_1", 1)
	time.Sleep(3 * time.Second)
	ContinuousMoveUp(dev.Device, "Profile_1", 1)
	time.Sleep(3 * time.Second)
	ContinuousMoveLow(dev.Device, "Profile_1", 1)
	time.Sleep(3 * time.Second)
	ContinuousMoveClose(dev.Device, "Profile_1", 1)
	time.Sleep(3 * time.Second)
	ContinuousMoveFar(dev.Device, "Profile_1", 1)
	time.Sleep(3 * time.Second)
	StopMove(dev.Device, "Profile_1")
}

func TestGetStatus(t *testing.T) {
	dev := getTestConfig()
	data, err := PztStatus(dev.Device, "Profile_1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
	t.Fail()
}
