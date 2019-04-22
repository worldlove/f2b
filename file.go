package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func readFile(filename string) error {
	basename := filepath.Base(filename)
	if pkgName == "" {
		pkgName = basename
	}
	if outFile == "" {
		outFile = basename + ".f2b.go"
	} else {
		if !strings.HasSuffix(outFile, ".go") {
			outFile += ".go"
		}
	}

	if varName == "" {
		varName = strings.ToUpper(basename[0:1]) + basename[1:]
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return writeFile(data)
}

func writeFile(data []byte) error {
	file, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	file.WriteString(fmt.Sprintf("var %s = %#v", varName, data))
	return nil
}
