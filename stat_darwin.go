package goutils

import (
	"os"
	"syscall"
	"time"
)

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func getFileInfo(fi os.FileInfo) SMFileInfo {
	var smfi SMFileInfo

	smfi.AccessTime = timespecToTime(fi.Sys().(*syscall.Stat_t).Atimespec)
	smfi.CreateTime = timespecToTime(fi.Sys().(*syscall.Stat_t).Ctimespec)
	smfi.ModifyTime = timespecToTime(fi.Sys().(*syscall.Stat_t).Mtimespec)
	return smfi
}
