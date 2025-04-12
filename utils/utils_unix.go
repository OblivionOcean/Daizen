//go:build linux || darwin
// +build linux darwin

package utils

import (
	"syscall"
)

func GetModTime(path string) (int64, error) {
	var stat syscall.Stat_t
	if err := syscall.Stat(path, &stat); err != nil {
		return 0, err
	}
	return stat.Mtim.Nano(), nil // 直接返回纳秒级时间戳
}

func fileExists(path string) (bool, error) {
	var stat syscall.Stat_t
	err := syscall.Stat(path, &stat)
	if err == nil {
		return true, nil
	}
	if err == syscall.ENOENT { // 文件不存在
		return false, nil
	}
	return false, err // 其他错误
}

func GetFdSize(fd uintptr) (int64, error) {
	var stat syscall.Stat_t
	if err := syscall.Fstat(int(fd), &stat); err != nil {
		return 0, err
	}
	return stat.Size, nil // 直接返回纳秒级时间戳
}
