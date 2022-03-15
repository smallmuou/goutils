package goutils

import (
	"os"
	"syscall"
	"time"
)

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func atime(fi os.FileInfo) time.Time {
	return timespecToTime(fi.Sys().(*syscall.Stat_t).Atim)
}

func getFileInfo(fi os.FileInfo) SMFileInfo {
	var smfi SMFileInfo

	smfi.AccessTime = timespecToTime(fi.Sys().(*syscall.Stat_t).Atim)
	smfi.CreateTime = timespecToTime(fi.Sys().(*syscall.Stat_t).Ctim)
	smfi.ModifyTime = timespecToTime(fi.Sys().(*syscall.Stat_t).Mtim)
	return smfi
}
