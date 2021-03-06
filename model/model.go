package model

import (
	"time"
)

type ResetMode struct {
	Hold     bool          `json:"hold"`
	Mode     string        `json:"mode"`
	Duration time.Duration `json:"duration"`
}

type DisplayUpdateProgress struct {
	Progress float64 `json:"progress"`
}

type IconRequest struct {
	Icon        string `json:"icon"`
	DisplayTime int    `json:"displayTime"`
}
