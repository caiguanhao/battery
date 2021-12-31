package battery

import (
	"os/exec"
)

func getStatus() (*Status, error) {
	status, err := getStatusWithPmset()
	if err != nil {
		status, err = getStatusWithIoreg()
	}
	return status, err
}

func getStatusWithPmset() (*Status, error) {
	cmd := exec.Command("pmset", "-g", "batt")
	outBytes, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parsePmsetOutput(string(outBytes))
}

func getStatusWithIoreg() (*Status, error) {
	cmd := exec.Command("ioreg", "-n", "AppleSmartBattery", "-r")
	outBytes, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseIoregOutput(string(outBytes))
}
