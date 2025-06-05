package scanner

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"os"
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
	signatures []Signature
}

func NewScanner() *DefaultScanner {
	return &DefaultScanner{
		signatures: DefaultSignatures,
	}
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
			return &ScanResult{
				Filename:   filepath,
				IsInfected: true,
				ThreatName: sig.Name,
			}, nil
		}
	}

	return &ScanResult{
		Filename:   filepath,
		IsInfected: false,
	}, nil
}
