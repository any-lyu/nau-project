package middleware

import (
	"runtime"
	"strconv"
	"strings"
)

func callers(skip int) []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip+1, pcs[:])
	return pcs[:n]
}

func stackString(stack []uintptr) string {
	if len(stack) == 0 {
		return ""
	}
	frames := runtime.CallersFrames(stack)

	var (
		frame    runtime.Frame
		more     bool
		funcName string
		fileName string
		buf      strings.Builder
	)
	for {
		frame, more = frames.Next()
		if frame.Function == "runtime.main" {
			break
		}
		if frame.Function == "runtime.goexit" {
			break
		}
		if frame.Function == "" {
			funcName = "unknown_function"
		} else {
			funcName = trimFuncName(frame.Function)
		}
		if frame.File == "" {
			fileName = "unknown_file"
		} else {
			fileName = trimFileName(frame.File)
		}
		if buf.Len() > 0 {
			buf.WriteByte('|')
		}
		buf.WriteString(funcName)
		buf.WriteString("(")
		buf.WriteString(fileName)
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(frame.Line))
		buf.WriteString(")")

		if !more {
			break
		}
	}
	return buf.String()
}

func trimFileName(name string) string {
	i := strings.Index(name, "/pkg/mod/") + len("/pkg/mod/")
	if i >= len("/pkg/mod/") && i < len(name) /* BCE */ {
		name = name[i:]
	}
	return name
}

func trimFuncName(name string) string {
	i := strings.LastIndexByte(name, '/')
	return name[i+1:]
}
