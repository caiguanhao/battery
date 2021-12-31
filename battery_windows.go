package battery

import (
	"syscall"
	"unsafe"
)

type systemPowerStatus struct {
	AcLineStatus        byte
	BatteryFlag         byte
	BatteryLifePercent  byte
	SystemStatusFlag    byte
	BatteryLifeTime     uint32
	BatteryFullLifeTime uint32
}

var (
	kernel32, _       = syscall.LoadLibrary("kernel32.dll")
	getPowerStatus, _ = syscall.GetProcAddress(kernel32, "GetSystemPowerStatus")
)

func getStatus() (*Status, error) {
	battery := new(systemPowerStatus)
	syscall.Syscall(uintptr(getPowerStatus), 1, uintptr(unsafe.Pointer(battery)), 0, 0)
	return &Status{
		Percent:     int(battery.BatteryLifePercent),
		Discharging: battery.AcLineStatus == 0,
	}, nil
}
