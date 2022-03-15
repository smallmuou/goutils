import (
	"os"
	"syscall"
	"time"
)

func getFileInfo(fi os.FileInfo) SMFileInfo {
	var smfi SMFileInfo

	smfi.AccessTime = time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).LastAccessTime.Nanoseconds())
	smfi.CreateTime = time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).CreationTime.Nanoseconds())
	smfi.ModifyTime = time.Unix(0, fi.Sys().(*syscall.Win32FileAttributeData).LastWriteTime.Nanoseconds())
	return smfi
}
