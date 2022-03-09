package goutils

import (
    "os"
    "fmt"
    "io"
    "crypto/md5"
    "encoding/hex"
)

func Md5(content string) (string, error) {
    h := md5.New()
    h.Write([]byte(content))
    return hex.EncodeToString(h.Sum(nil)), nil
}

func Md5File(filename string) (string, error) {
    f, err := os.Open(filename) //打开文件
    if nil != err {
        fmt.Println(err)
        return "", err
    }
    defer f.Close()

    md5Handle := md5.New()      //创建 md5 句柄
    _, err = io.Copy(md5Handle, f)  //将文件内容拷贝到 md5 句柄中
    if nil != err {
        fmt.Println(err)
        return "", err
    }
    md := md5Handle.Sum(nil)    //计算 MD5 值，返回 []byte

    md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
    return md5str, nil
}
