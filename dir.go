package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var fileMap = make(map[string][]byte, 0)

func readDir(dirname string) error {
	if pkgName == "" {
		pkgName = dirname
	}
	if outFile == "" {
		outFile = dirname + ".f2b.go"
	} else {
		if !strings.HasSuffix(outFile, ".go") {
			outFile += ".go"
		}
	}
	if varName == "" {
		varName = strings.ToUpper(dirname[0:1]) + dirname[1:]
	}
	dir, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}
	sep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			readDir(dirname + sep + fi.Name())
			continue
		}
		filepath := dirname + sep + fi.Name()
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}
		if _, ok := fileMap[fi.Name()]; ok {
			fileMap[dirname+"-"+fi.Name()] = data
		} else {
			fileMap[fi.Name()] = data
		}
	}
	return writeMapToFile(fileMap)
}

func writeMapToFile(datas map[string][]byte) error {
	file, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	file.WriteString(fmt.Sprintf("var %s = map[string][]byte{\n", varName))
	for k, v := range datas {
		file.WriteString(fmt.Sprintf("\t\"%s\": %#v,\n", k, v))
	}
	file.WriteString("}\n")
	return nil
}
