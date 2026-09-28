//go:build !darwin && !windows
// +build !darwin,!windows

package gokb

import "syscall"

const tcpKeepCnt = syscall.TCP_KEEPCNT
