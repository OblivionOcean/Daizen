//go:build linux
// +build linux

package utils

import "syscall"

func GetModTime(path string) (int64, error) {
	var stat syscall.Stat_t
	if err := syscall.Stat(path, &stat); err != nil {
		return 0, err
	}
	return stat.Mtim.Nano(), nil // 直接返回纳秒级时间戳
}
