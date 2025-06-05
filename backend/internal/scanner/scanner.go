package scanner

import (
	"context"
	"time"
)

type ScanResult struct {
	Filename     string
	IsInfected   bool
	ThreatName   string
	ScanTime     time.Time
	FileSize     int64
	ScanDuration time.Duration
}

type Scanner interface {
	ScanFile(ctx context.Context, filepath string) (*ScanResult, error)
	ScanDirectory(ctx context.Context, dirpath string) ([]ScanResult, error)
}

type DefaultScanner struct {
	// será implementado no próximo commit
}
