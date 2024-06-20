package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var homedir string
var curcfg string

type config struct {
	Generator string `json:"generator"`
	Solve1    string `json:"working"`
	Solve2    string `json:"optimal"`
	Removeout bool   `json:"removeout"`
	Compiler  string `json:"gnuorclang"`
}

func cfgreadfile(cfgfile *os.File) (cfg config, err error) {
	var cfgjson []byte
	fmt.Fscanf(cfgfile, "%s", &cfgjson)
	err = json.Unmarshal(cfgjson, &cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	return cfg, err
}

func cfgwritefile(cfgfile *os.File, cfg config) (err error) {
	var cfgjson []byte
	cfgjson, err = json.Marshal(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(cfgfile, "%s", cfgjson)
	return
}

func dircheckcurcfg() (err error) {
	_, err = os.Stat(string(homedir + "/shar/"))
	if os.IsNotExist(err) {
		err = os.Mkdir(string(homedir+"/shar/"), os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = os.Stat(string(homedir + "/shar/cur.cfg"))
	if os.IsNotExist(err) {
		_, err = os.Create(string(homedir + "/shar/cur.cfg"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func writecurcfg(cfgname string) (err error) {
	filecurcfg, err := os.OpenFile(string(homedir+"/shar/cur.cfg"), os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filecurcfg.Close()
	fmt.Fprint(filecurcfg, cfgname)
	return
}

func readcurcfg() (err error) {
	filecurcfg, err := os.OpenFile(string(homedir+"/shar/cur.cfg"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filecurcfg.Close()
	fmt.Fscan(filecurcfg, &curcfg)
	return
}
