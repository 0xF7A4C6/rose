package utils

import (
	"bufio"
	"fmt"
	"os"
)

func HandleError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func Log(Content string) {
	fmt.Println(fmt.Sprintf("%s.", Content))
}

func Debug(Content string) {
	fmt.Println(fmt.Sprintf("%s.", Content))
}

func FormatSocketString(Content string) string {
	return fmt.Sprintf(fmt.Sprintf("%s\n\r", Content))
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// append line if file
func AppendLine(path string, line string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s\n", line))
	return err
}

// check if string in list
func StringInList(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
