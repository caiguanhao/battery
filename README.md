# battery

Cross-platform get battery status.

Tested on Arch Linux, Debian, Ubuntu, Windows, macOS.

```go
import "github.com/caiguanhao/battery"

battery.GetStatus()
// &battery.Status{Percent:80 Discharging:true}, nil
```
