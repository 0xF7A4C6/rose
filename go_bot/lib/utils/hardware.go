package utils

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetMemory() int {
	out, err := exec.Command("free", "-m").Output()
	if HandleError(err) {
		return 0
	}

	lines := strings.Split(string(out), "\n")
	if HandleError(err) {
		return 0
	}

	fields := strings.Fields(lines[1])
	if len(fields) < 2 {
		return 0
	}

	i, err := strconv.Atoi(fields[1])
	if HandleError(err) {
		return 0
	}

	return i
}

func GetDiskSpace() int {
	return 0
	/*wd, err := os.Getwd()
	if HandleError(err) {
		return "nan"
	}

	var stat unix.Statfs_t
	unix.Statfs(wd, &stat)
	return fmt.Sprint(stat.Bavail * uint64(stat.Bsize) / 1024 / 1024 / 1024) // GB*/
}

// get CPU arch
func GetCPUArch() string {
	out, err := exec.Command("uname", "-m").Output()
	if HandleError(err) {
		return ""
	}

	return strings.TrimSpace(string(out))
}
