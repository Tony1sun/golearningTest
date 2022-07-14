package util

import "crypto/md5"

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
}
