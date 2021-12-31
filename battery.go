package battery

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Status struct {
	Percent     int
	Discharging bool
}

var (
	ErrParsingData  = errors.New("error parsing data")
	ErrNotSupported = errors.New("not supported")
)

func GetStatus() (*Status, error) {
	return getStatus()
}

func parsePmsetOutput(out string) (*Status, error) {
	acPower := strings.Contains(out, "AC Power")
	percentMatches := regexp.MustCompile("([0-9]+)%").FindStringSubmatch(out)
	if len(percentMatches) < 2 {
		return nil, ErrParsingData
	}
	percent, _ := strconv.Atoi(percentMatches[1])
	return &Status{
		Percent:     percent,
		Discharging: !acPower,
	}, nil
}

func parseIoregOutput(out string) (*Status, error) {
	mcMatches := regexp.MustCompile(`MaxCapacity.+?([0-9]+)`).FindStringSubmatch(out)
	ccMatches := regexp.MustCompile(`CurrentCapacity.+?([0-9]+)`).FindStringSubmatch(out)
	ecMatches := regexp.MustCompile(`ExternalConnected.+?(\w+)`).FindStringSubmatch(out)
	if len(mcMatches) == 2 && len(ccMatches) == 2 && len(ecMatches) == 2 {
		mc, merr := strconv.Atoi(mcMatches[1])
		cc, cerr := strconv.Atoi(ccMatches[1])
		if merr == nil && cerr == nil {
			percent := int(float64(cc) / float64(mc) * 100)
			return &Status{
				Percent:     percent,
				Discharging: ecMatches[1] == "No",
			}, nil
		}
	}
	return nil, ErrParsingData
}
