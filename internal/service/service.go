package service

import (
	"fmt"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysUntilCounter() string {
	endTime := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(endTime)
	days := fmt.Sprintf("Days until 25 JAN 2025:%d", hoursToDays(dur))
	return days
}

func hoursToDays(d time.Duration) int64 {
	return int64(d.Hours() / 24)
}
