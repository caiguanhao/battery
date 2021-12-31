package battery

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func getStatus() (*Status, error) {
	const prefix = "/sys/class/power_supply/BAT0"
	statusBytes, err := ioutil.ReadFile(prefix + "/status")
	if err != nil {
		return nil, ErrParsingData
	}

	fBytes, err := ioutil.ReadFile(prefix + "/energy_full")
	if err != nil {
		fBytes, err = ioutil.ReadFile(prefix + "/charge_full")
	}
	if err != nil {
		return nil, ErrParsingData
	}

	nBytes, err := ioutil.ReadFile(prefix + "/energy_now")
	if err != nil {
		nBytes, err = ioutil.ReadFile(prefix + "/charge_now")
	}
	if err != nil {
		return nil, ErrParsingData
	}

	status := strings.TrimSpace(string(statusBytes))
	n, nerr := strconv.Atoi(strings.TrimSpace(string(nBytes)))
	f, ferr := strconv.Atoi(strings.TrimSpace(string(fBytes)))
	if nerr == nil && ferr == nil {
		percent := int(float64(n) / float64(f) * 100)
		return &Status{
			Percent:     percent,
			Discharging: status == "Discharging",
		}, nil
	}
	return nil, ErrParsingData
}
