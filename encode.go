package goutils

import (
    "os"
    "fmt"
    "io"
    "crypto/md5"
)

func Md5(content string) (string, error) {
    h := md5.New()
    h.Write([]byte(content))

    //将 []byte 转为 string
    md5str := fmt.Sprintf("%x", h.Sum(nil))

    return md5str, nil
}

func Md5File(filename string) (string, error) {
    //打开文件
    f, err := os.Open(filename)
    if nil != err {
        fmt.Println(err)
        return "", err
    }
    defer f.Close()

    //计算 MD5 值，返回 []byte
    md5Handle := md5.New()
    _, err = io.Copy(md5Handle, f)
    if nil != err {
        fmt.Println(err)
        return "", err
    }
    md := md5Handle.Sum(nil)

    //将 []byte 转为 string
    md5str := fmt.Sprintf("%x", md)
    return md5str, nil
}
