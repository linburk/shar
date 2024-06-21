package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var cfg config

var stderr bytes.Buffer

func shellcmd(command string) (err error) {
	var cmd exec.Cmd
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Удали виндоус")
	default:
		cmd = *exec.Command("/bin/zsh", "-c", command)
	}
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err, stderr.String())
		return
	}
	return
}

func compile() (err error) {
	err = shellcmd(cfg.Compiler + " -std=c++20 -O2 " + cfg.Solve1 + " -o " + homedir + "/shar/solve1")
	if err != nil {
		return
	}
	err = shellcmd(cfg.Compiler + " -std=c++20 -O2 " + cfg.Solve2 + " -o " + homedir + "/shar/solve2")
	if err != nil {
		return
	}
	err = shellcmd(cfg.Compiler + " -std=c++20 -O2 " + cfg.Generator + " -o " + homedir + "/shar/gen")
	if err != nil {
		return
	}
	return
}

func tempout() (err error) {
	_, err = os.Create(homedir + "/shar/out1.o")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = os.Create(homedir + "/shar/out2.o")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = os.Create(homedir + "/shar/gen.o")
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func test(num int) (wa bool, err error) {
	result := true
	err = shellcmd(homedir + "/shar/./gen > " + homedir + "/shar/gen.o")
	if err != nil {
		fmt.Println(err)
		return
	}
	start1 := time.Now()
	err = shellcmd(homedir + "/shar/./solve1 < " + homedir + "/shar/gen.o > " + homedir + "/shar/out1.o")
	if err != nil {
		fmt.Println(err)
		return
	}
	time1 := time.Since(start1)
	start2 := time.Now()
	err = shellcmd(homedir + "/shar/./solve2 < " + homedir + "/shar/gen.o > " + homedir + "/shar/out2.o")
	if err != nil {
		fmt.Println(err)
		return
	}
	time2 := time.Since(start2)
	file1, err := os.Open(homedir + "/shar/out1.o")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file1.Close()
	file2, err := os.Open(homedir + "/shar/out2.o")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()
	out1 := bufio.NewScanner(file1)
	out2 := bufio.NewScanner(file2)
	for out1.Scan() && out2.Scan() {
		if out1.Text() != out2.Text() {
			result = false
		}
	}
	if out1.Scan() || out2.Scan() {
		result = false
	}
	fmt.Printf("\tTest %d\n", num)
	switch result {
	case true:
		fmt.Println("----------------------")
		fmt.Println("\t  OK\t")
		fmt.Println("----------------------")
		fmt.Println("Time 1st:", time1)
		fmt.Println("Time 2nd:", time2)
		fmt.Println("----------------------")
	case false:
		wa = true
		fmt.Println("----------------------")
		fmt.Println("\t  WA\t")
		fmt.Println("----------------------")
		fmt.Println("Time 1st:", time1)
		fmt.Println("Time 2nd:", time2)
		fmt.Println("----------------------\nTest case:")
		testfile, err := os.Open(homedir + "/shar/gen.o")
		if err != nil {
			fmt.Println(err)
			return wa, err
		}
		defer testfile.Close()
		test := bufio.NewScanner(testfile)
		for test.Scan() {
			fmt.Println(test.Text())
		}
		if out1.Scan() || out2.Scan() {
			fmt.Println("Extra output")
		} else {
			file1, err = os.Open(homedir + "/shar/out1.o")
			if err != nil {
				fmt.Println(err)
				return wa, err
			}
			defer file1.Close()
			file2, err = os.Open(homedir + "/shar/out2.o")
			if err != nil {
				fmt.Println(err)
				return wa, err
			}
			out1 = bufio.NewScanner(file1)
			out2 = bufio.NewScanner(file2)
			fmt.Println("1st solve's output:")
			for out1.Scan() {
				fmt.Println(out1.Text())
			}
			fmt.Println("2nd solve's output:")
			for out2.Scan() {
				fmt.Println(out2.Text())
			}

		}
		fmt.Println("----------------------")
	}
	return
}

func runmain() (err error) {
	if len(os.Args) < 3 {
		fmt.Println("Not enough args")
		return
	}
	iters, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	err = readcurcfg()
	if err != nil {
		fmt.Println(err)
		return
	}
	cfgfile, err := os.OpenFile(string(homedir+"/shar/"+curcfg), os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cfgfile.Close()
	cfg, err = cfgreadfile(cfgfile)
	if err != nil {
		return
	}
	err = compile()
	if err != nil {
		return
	}
	fmt.Println("Compiled")
	err = tempout()
	if err != nil {
		return
	}
	for i := range iters {
		wa, err := test(i + 1)
		if err != nil {
			return
		}
		if wa {
			break
		}
	}
	return
}
