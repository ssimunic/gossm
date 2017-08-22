package gossm

import (
	"time"
)

type StatusData struct {
	StatusAtTime map[time.Time]bool `json:"statusAtTime"`
}
