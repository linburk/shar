package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const cfgPath string = "stress_cfg/"
const cfgFilename string = "config.json"

type config struct {
	FilePath     string   `json:"filepath"`
	Compare      bool     `json:"compare"`
	Removeout    bool     `json:"removeout"`
	FileNum      int      `json:"filenum"`
	Files        []string `json:"files"`
	Languages    []string `json:"languages"`
	Compilators  []string `json:"compilators"`
	CompileFlags []string `json:"compileflags"`

	FilenPaths     []string `json:"filenpaths"`
	FilenPathsExec []string `json:"filenpathsexec"`
	FilenPathsComp []string `json:"filenpathscomp"`
	OutFilenPath   []string `json:"outfilenpath"`
}

func writeFilenPaths(mode int, cfg config) (filenPaths []string) {
	for i, file := range cfg.Files {
		var filenPath string
		if mode == 1 { // exec compilated
			if cfg.Languages[i] == "Python" {
				filenPath = cfg.Compilators[i] + " " + cfg.FilePath + file
			} else {
				filenPath = cfg.FilePath + "./" + eraseExt(file)
			}
		}
		if mode == 2 { // compilated
			if cfg.Languages[i] == "Python" {
				continue
			} else {
				filenPath = cfg.FilePath + eraseExt(file)
			}
		}
		if mode == 3 { // outfiles
			filenPath = cfg.FilePath + eraseExt(file) + ".out"
		}
		if mode == 4 { // no changes
			filenPath = cfg.FilePath + file
		}
		filenPaths = append(filenPaths, filenPath)
	}
	return
}

func eraseExt(filename string) string {
	filename = filename[:len(filename)-len(filepath.Ext(filename))]
	return filename
}

func setupCfg(fileName string, rewrite bool) (cfg config, err error) {
	writed := false
	cfgFile, err := os.OpenFile(cfgPath+fileName, os.O_WRONLY, os.ModePerm)
	if os.IsNotExist(err) {
		err = os.MkdirAll(cfgPath, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprint(os.Stderr, "Mkdir error\n")
			return
		}
		_, err = createCfgFile(fileName)
		if err != nil {
			return
		}
		cfgFile, err = os.OpenFile(cfgPath+fileName, os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			fmt.Fprint(os.Stderr, "File open after create error\n")
			return
		}
		inputConfig(cfgFile)
		writed = true
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprint(os.Stderr, "File open error\n")
		return
	}
	defer cfgFile.Close()
	if !writed && rewrite {
		inputConfig(cfgFile)
	}
	cfg, err = readCfgFile(cfgFile)
	return
}

func readCfgFile(cfgFile *os.File) (cfg config, err error) {
	var cfgjson []byte
	fmt.Fscan(cfgFile, &cfgjson)
	err = json.Unmarshal(cfgjson, &cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprint(os.Stderr, "Unmarshall error\n")
		return
	}
	return
}
func createCfgFile(cfgName string) (cfgFile *os.File, err error) {
	cfgFile, err = os.OpenFile(cfgPath+cfgName, os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprint(os.Stderr, "File create error\n")
		return
	}
	return
}

func inputConfig(cfgFile *os.File) {
	var cfg config
	fmt.Print("Enter path to your code...\n")
	fmt.Scan(&cfg.FilePath)
	fmt.Print("Enter number of files, last file must be generator...\n")
	fmt.Scan(&cfg.FileNum)
	if cfg.FileNum == 3 {
		fmt.Print("Check comparison? Y/n \n")
		var comparisonFlag string
		fmt.Scan(&comparisonFlag)
		if comparisonFlag == "Y" || comparisonFlag == "y" {
			cfg.Compare = true
		}
	}
	fmt.Print("Remove .txt files? Y/n\n")
	var rmtxt string
	fmt.Scan(&rmtxt)
	if rmtxt == "Y" || rmtxt == "y" {
		cfg.Removeout = true
	}
	for i := 0; i < cfg.FileNum; i++ {
		fmt.Printf("Enter %d filename\n", i+1)
		var filename string
		fmt.Scan(&filename)
		cfg.Files = append(cfg.Files, filename)
	}
	for i := 0; i < cfg.FileNum; i++ {
		fmt.Printf("Enter %d file language (Python / C++ / C)\n", i+1)
		var lang string
		fmt.Scan(&lang)
		cfg.Languages = append(cfg.Languages, lang)
	}
	for i := 0; i < cfg.FileNum; i++ {
		fmt.Printf("Enter %d file compilator (g++, python3, etc.) \n", i+1)
		var compilator string
		fmt.Scan(&compilator)
		cfg.Compilators = append(cfg.Compilators, compilator)
	}
	for i := 0; i < cfg.FileNum; i++ {
		fmt.Printf("Enter %d file compile flags or 0 if not (without -) \n", i+1)
		var compileFlag string
		fmt.Scan(&compileFlag)
		if compileFlag == "0" {
			continue
		}
		cfg.CompileFlags = append(cfg.CompileFlags, compileFlag)
	}
	cfg.FilenPathsExec = writeFilenPaths(1, cfg)
	cfg.FilenPathsComp = writeFilenPaths(2, cfg)
	cfg.OutFilenPath = writeFilenPaths(3, cfg)
	cfg.FilenPaths = writeFilenPaths(4, cfg)
	cfgjson, err := json.Marshal(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprint(os.Stderr, "Marshal error\n")
		return
	}
	fmt.Fprint(cfgFile, cfgjson)
}
