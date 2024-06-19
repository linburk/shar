package main

func compileCmdsList(cfg config) (compileCmds []string) {
	for i := 0; i < cfg.FileNum; i++ {
		if cfg.Languages[i] == "Python" {
			continue
		}
		compileCmd := cfg.Compilators[i] + " " + cfg.CompileFlags[i] + " " + cfg.FilenPaths[i] + " -o " + eraseExt(cfg.FilenPaths[i])
		compileCmds = append(compileCmds, compileCmd)
	}
	return
}
func execCmdsList(execFilenPaths, outFilenPaths []string) (execCmds []string) {
	for i, txtFile := range outFilenPaths[:len(outFilenPaths)-1] {
		execCmd := execFilenPaths[i] + " < " + outFilenPaths[len(outFilenPaths)-1] + " > " + txtFile
		execCmds = append(execCmds, execCmd)
	}
	execCmdGen := execFilenPaths[len(execFilenPaths)-1] + " > " + outFilenPaths[len(outFilenPaths)-1]
	execCmds = append(execCmds, execCmdGen)
	return
}
func rmCmdsList(filenPaths []string) (removeCmds []string) {
	for _, filenPath := range filenPaths {
		removeCmd := "rm " + filenPath
		removeCmds = append(removeCmds, removeCmd)
	}
	return
}
