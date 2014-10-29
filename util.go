package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

// 计算string的md5值，以32位字符串形式返回
func StringToMd5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}

// 计算[]byte的md5值，以32位字符串形式返回
func BytesToMd5(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

// 计算文件的md5值，以32位字符串形式返回
func FileToMd5(name string) (string, error) {
	h := md5.New()
	if err := readFile(name, h); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// 时间戳转换为string显示
func TimestampToString(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

// 辅助函数，扫描文件内容并编码到hash.Hash中
func readFile(name string, h hash.Hash) error {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return err
	}

	file, err := os.Open(name)
	if err != nil {
		return err
	}

	s := bufio.NewScanner(file)
	for s.Scan() {
		h.Write(s.Bytes())
	}

	return s.Err()
}
