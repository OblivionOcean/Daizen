//go:build windows
// +build windows

package utils

import (
	"syscall"
	"unsafe"
)

func GetModTime(path string) (int64, error) {
	var fileInfo syscall.Win32FileAttributeData
	pt, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return 0, err
	}
	err = syscall.GetFileAttributesEx(
		pt,
		syscall.GetFileExInfoStandard,
		(*byte)(unsafe.Pointer(&fileInfo)),
	)
	if err != nil {
		return 0, err
	}

	// 转换为 100-nanosecond 间隔数 (从 1601-01-01 开始)
	ft := fileInfo.LastWriteTime.Nanoseconds()

	// 转换为 Unix 纪元毫秒时间戳
	//unixNs := ft - 116444736000000000 // 调整到 Unix 纪元
	return ft, nil // 转换为毫秒
}

func fileExists(path string) (bool, error) {
	pathPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return false, err
	}

	attrs, err := syscall.GetFileAttributes(pathPtr)
	if err != nil {
		if errno, ok := err.(syscall.Errno); ok {
			if errno == syscall.ERROR_FILE_NOT_FOUND ||
				errno == syscall.ERROR_PATH_NOT_FOUND {
				return false, nil
			}
		}
		return false, err
	}
	return attrs != syscall.INVALID_FILE_ATTRIBUTES, nil
}

func GetFdSize(fd uintptr) (int64, error) {
	var fileInfo syscall.ByHandleFileInformation
	err := syscall.GetFileInformationByHandle(
		syscall.Handle(fd),
		&fileInfo,
	)
	if err != nil {
		return 0, err
	}

	fs := int64(fileInfo.FileSizeHigh)*1e9 + int64(fileInfo.FileSizeLow)
	return fs, nil // 转换为毫秒
}
