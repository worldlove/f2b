package main

import (
	"flag"
	"fmt"
	"os"
)

// 实际中应该用更好的变量名
var (
	help bool

	dir     string
	file    string
	varName string
	pkgName string
	outFile string
)

func init() {
	flag.BoolVar(&help, "h", false, "Help, print the usage ")

	flag.StringVar(&dir, "d", "", "Set the static dir, build a map[key<string>]value<[]byte> that with the key of filename and the value of filedata")
	flag.StringVar(&file, "f", "", "Set the static file, build a var named of filename and value of the filedata")
	flag.StringVar(&varName, "v", "", "Set the var name, default the dir name or file name")
	flag.StringVar(&pkgName, "p", "", "Set the package name, default the dir name or the file's basename")
	flag.StringVar(&outFile, "o", "", "Set the output file name, default is ./{filename|dirname}.f2b.go")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `f2b version: 0.1
Usage: f2b [-h] [-d dirname] [-f filename] [-p package] [-v var]

Options:
`)
	flag.PrintDefaults()
}
