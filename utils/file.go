package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func GetImagName(name, ext string) string {
	m := md5.New()
	m.Write([]byte(name))
	hexName := hex.EncodeToString(m.Sum(nil))

	return "uploads/" + time.Now().Format("2006/01/02") + "/" + hexName + ext
}
