package quarantine

import (
	"crypto/rand"
	"os"
	"time"
)

type QuarantineInfo struct {
	OriginalPath   string    `json:"originalPath"`
	QuarantineDate time.Time `json:"quarantineDate"`
	ThreatName     string    `json:"threatName"`
	FileSize       int64     `json:"fileSize"`
}

type Quarantine struct {
	basePath string
	key      []byte
}

func NewQuarantine(path string) (*Quarantine, error) {
	if err := os.MkdirAll(path, 0750); err != nil {
		return nil, err
	}

	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return &Quarantine{
		basePath: path,
		key:      key,
	}, nil
}

func (q *Quarantine) QuarantineFile(filepath string, threatName string) error {
	info, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	quarantineID := time.Now().Format("20060102150405") + "_" + filepath

	quarantineInfo := QuarantineInfo{
		OriginalPath:   filepath,
		QuarantineDate: time.Now(),
		ThreatName:     threatName,
		FileSize:       info.Size(),
	}

	return q.moveToQuarantine(filepath, quarantineID, quarantineInfo)
}
