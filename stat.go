package goutils

import (
	"os"
	"time"
)

//文件信息
type SMFileInfo struct {
	CreateTime time.Time
	AccessTime time.Time
	ModifyTime time.Time
}

func Stat(name string) (SMFileInfo, error) {
	fi, err := os.Stat(name)

	if err != nil {
		return SMFileInfo{}, err
	}
	return getFileInfo(fi), nil
}
