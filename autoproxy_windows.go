//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
)

func main() {
	dll := syscall.MustLoadDLL("DLLTest.dll")
	p := dll.MustFindProc("Add")

	fmt.Println(p.Call(uintptr(1), uintptr(1)))
}
