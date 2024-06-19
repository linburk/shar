package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func getCmds(cfg config) (compileCmds, rmCmds, execCmds []string) {
	compileCmds = compileCmdsList(cfg)
	rmCmds = append(rmCmds, rmCmdsList(cfg.FilenPathsComp)...)
	if cfg.Removeout {
		rmCmds = append(rmCmds, rmCmdsList(cfg.OutFilenPath)...)
	}
	execCmds = execCmdsList(cfg.FilenPathsExec, cfg.OutFilenPath)
	return compileCmds, rmCmds, execCmds
}

var stderr bytes.Buffer

const shell string = "/bin/zsh"

func launchCmds(cmdsList ...string) (err error) {
	for _, strcmd := range cmdsList {
		cmd := exec.Command(shell, "-c", strcmd)
		cmd.Stderr = &stderr
		err = cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err, stderr.String())
			fmt.Fprint(os.Stderr, "Command error\n")
			return
		}
	}
	return
}

func compareSolves(outFilenPathA, outFilenPathB string) (result bool, err error) {
	var outA, outB string
	outFileA, err := os.OpenFile(outFilenPathA, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprint(os.Stderr, "Out file open error\n")
	}
	outFileB, err := os.OpenFile(outFilenPathB, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprint(os.Stderr, "Out file open error\n")
	}
	resA := bufio.NewScanner(outFileA)
	resB := bufio.NewScanner(outFileB)
	for resA.Scan() && resB.Scan() {
		outA = resA.Text()
		outB = resB.Text()
		if outA != outB {
			return false, err
		}
	}
	return true, err
}

func main() {
	var cfgMain config
	var execCmds []string
	var rmCmds []string
	var compileCmds []string
	iters, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprint(os.Stderr, "Wrong arguments!\n")
		return
	}
	cfgMain, err = setupCfg(cfgFilename, len(os.Args) == 3)
	if err != nil {
		return
	}
	compileCmds, rmCmds, execCmds = getCmds(cfgMain)
	err = launchCmds(compileCmds...)
	if err != nil {
		return
	}
	for i := iters; i < iters; i++ {
		err = launchCmds(execCmds[len(execCmds)-1])
		if err != nil {
			return
		}
		fmt.Printf("------------------\nTest %d generated\n------------------\n", i+1)
		for j, execSolveCmd := range cfgMain.FilenPathsExec[:cfgMain.FileNum-1] {
			start := time.Now()
			err = launchCmds(execSolveCmd)
			if err != nil {
				return
			}
			fmt.Printf("Time %s : %s\n", cfgMain.Files[j], time.Since(start))
		}
		if cfgMain.Compare {
			res, err := compareSolves(cfgMain.OutFilenPath[0], cfgMain.OutFilenPath[1])
			if err != nil {
				return
			}
			if !res {
				fmt.Printf("Wrong answer!\n")
				return
			}
		}
	}
	err = launchCmds(rmCmds...)
	if err != nil {
		return
	}
}
