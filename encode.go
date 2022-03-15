// Copyright 2022 smallmuou.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//公共基础包
package goutils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

/**
* generate string md5
*
* param: content    content to convert
* return: md5str    md5 hex string
 */
func Md5(content string) (md5str string) {
	h := md5.New()
	h.Write([]byte(content))

	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
* generate file md5
*
* param:    filename    file path
* return:   md5str      md5 hex string
            error       error
*/
func Md5File(filename string) (md5str string, err error) {
	//open file
	f, err := os.Open(filename)
	if nil != err {
		fmt.Println(err)
		return "", err
	}

	//delay close file
	defer f.Close()

	//generate md5
	md5Handle := md5.New()
	_, err = io.Copy(md5Handle, f)
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	md := md5Handle.Sum(nil)

	return fmt.Sprintf("%x", md), nil
}

/**
* base64 encode
*
* param:    data        the data to encode
* return:   base64str   string
 */
func Base64Encode(data []byte) (base64str string) {
	return base64.StdEncoding.EncodeToString(data)
}

/**
* base64 decode
*
* param:    content     the content to decode
* return:   data        decoded data
            err         error
*/
func Base64Decode(content string) (data []byte, err error) {
	return base64.StdEncoding.DecodeString(content)
}
