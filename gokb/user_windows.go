package gokb

import (
	"path/filepath"
	"syscall"
)

// GetUserNameEx相比于GetUserName有更广泛的可用名称
// 为了使输出与GetUserName相同，只返回最基础(或最后的)的组件
func userCurrent() (us string, rErr error) {
	pwName := make([]uint16, 128)
	pwNameSize := uint32(len(pwName)) - 1
	err := syscall.GetUserNameEx(syscall.NameSamCompatible, &pwName[0], &pwNameSize)
	if err != nil {
		us = ""
		rErr = ErrCouldNotDetectUsername
		return
	}
	s := syscall.UTF16ToString(pwName)
	u := filepath.Base(s)
	us = u
	err = nil
	return
}
