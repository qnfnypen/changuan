package opfile

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

// CreateFileWithTimeStamp 创建带有时间戳的文件
func CreateFileWithTimeStamp(fp string) (*os.File, error) {
	dir := filepath.Dir(fp)
	if err := os.MkdirAll(dir, 0666); err != nil {
		return nil, err
	}

	ext := filepath.Ext(fp)
	f := strings.TrimSuffix(fp, ext)
	f += "_" + time.Now().Format("20060102") + ext

	return os.OpenFile(f, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
}
