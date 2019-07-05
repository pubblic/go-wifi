package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var (
	DLLUser32 = syscall.NewLazyDLL("User32.dll")
)

func WarnBox(caption, text string) error {
	const NULL = 0
	const FuncName = "MessageBoxW"
	const (
		MB_OK        = 0x0
		MB_ICONERROR = 0x10
	)
	err := DLLUser32.Load()
	if err != nil {
		return err
	}
	_, _, err = DLLUser32.NewProc(FuncName).Call(
		NULL,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		MB_OK|MB_ICONERROR,
	)
	return err
}

func perror(code int, text string) {
	err := WarnBox("Error", text)
	if err == nil {
		os.Exit(code)
		return
	}
	fmt.Fprintln(os.Stderr, text)
	fmt.Println("엔터를 누르시면 프로그램을 종료합니다.")
	getchar()
	os.Exit(code)
}

func getchar() (byte, error) {
	var buf [1]byte
	_, err := os.Stdin.Read(buf[:])
	return buf[0], err
}
