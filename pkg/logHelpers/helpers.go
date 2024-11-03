package logHelpers

import (
	"path/filepath"
	"runtime"
	"strconv"
)

func GetCodeLine() string {
	_, filename, line, _ := runtime.Caller(1)
	return filepath.Base(filename) + ":" + strconv.Itoa(line)
}
