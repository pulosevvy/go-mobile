package methods

import (
	"fmt"
	"time"
)

func FloatToHours(f float64) string {
	duration := time.Duration(f * float64(time.Hour))
	return fmt.Sprintf("%02d:%02d:%02d", int(duration.Hours()), int(duration.Minutes())%60, int(duration.Seconds())%60)
}
