//go:build darwin
// +build darwin

package gokb

// syscall 包在 darwin 平台未导出 TCP_KEEPCNT（macOS 内核值为 0x102），此处用常量兜底。
const tcpKeepCnt = 0x102
