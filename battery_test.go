package battery

import (
	"testing"
)

func TestGetStatus(t *testing.T) {
	status, err := GetStatus()
	t.Logf("Current Battery Status: %+v; Error: %+v", status, err)
}

func Test_parsePmsetOutput(t *testing.T) {
	assert := makeAssert(t, parsePmsetOutput)

	assert(`
Now drawing from 'Battery Power'
 -InternalBattery-0 (id=6815843)        96%; discharging; 11:25 remaining present: true
`, Status{96, true})

	assert(`
Now drawing from 'AC Power'
 -InternalBattery-0 (id=6815843)        96%; AC attached; not charging present: true
`, Status{96, false})
}

func Test_parseIoregOutput(t *testing.T) {
	assert := makeAssert(t, parseIoregOutput)

	assert(`
+-o AppleSmartBattery  <class AppleSmartBattery, id 0x1000005d1, registered, matched, active, busy 0 (0 ms), retain 9>
    {
      "CurrentCapacity" = 94
      "ExternalConnected" = No
      "FullyCharged" = No
      "MaxCapacity" = 100
    }
`, Status{94, true})

	assert(`
+-o AppleSmartBattery  <class AppleSmartBattery, id 0x1000005d1, registered, matched, active, busy 0 (0 ms), retain 9>
    {
      "CurrentCapacity" = 94
      "ExternalConnected" = Yes
      "FullyCharged" = No
      "MaxCapacity" = 100
    }
`, Status{94, false})
}

func makeAssert(t *testing.T, f func(string) (*Status, error)) func(string, Status) {
	return func(str string, expected Status) {
		t.Helper()
		status, err := f(str)
		if err != nil {
			t.Error(err)
		} else if expected != *status {
			t.Errorf("status should equal to %+v instead of %+v", expected, *status)
		}
	}
}
