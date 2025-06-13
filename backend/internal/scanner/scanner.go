package scanner

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os"
	"time"

	"github.com/GabrielNat1/ViruScanGo/internal/quarantine"
)

type ScanResult struct {
	Filename     string
	IsInfected   bool
	ThreatName   string
	ScanTime     time.Time
	FileSize     int64
	ScanDuration time.Duration
	Quarantined  bool
}

type Scanner interface {
	ScanFile(ctx context.Context, filepath string) (*ScanResult, error)
	ScanDirectory(ctx context.Context, dirpath string) ([]ScanResult, error)
}

type DefaultScanner struct {
	signatures []Signature
	quarantine *quarantine.Quarantine
}

func NewScanner(quarantinePath string) (*DefaultScanner, error) {
	q, err := quarantine.NewQuarantine(quarantinePath)
	if err != nil {
		return nil, err
	}

	return &DefaultScanner{
		signatures: DefaultSignatures,
		quarantine: q,
	}, nil
}

func (s *DefaultScanner) ScanFile(ctx context.Context, filepath string) (*ScanResult, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 4096)

	for _, sig := range s.signatures {
		file.Seek(sig.Offset, 0)
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if bytes.Contains(buffer[:n], sig.Pattern) {
			err := s.quarantine.QuarantineFile(filepath, sig.Name)
			if err != nil {
				return nil, err
			}

			return &ScanResult{
				Filename:    filepath,
				IsInfected:  true,
				ThreatName:  sig.Name,
				Quarantined: true,
			}, nil
		}
	}

	return &ScanResult{
		Filename:   filepath,
		IsInfected: false,
	}, nil
}
