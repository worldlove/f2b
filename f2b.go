package main

import (
	"flag"
	"log"
)

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if dir != "" {
		if err := readDir(dir); err != nil {
			log.Fatal(err)
		}
	} else if file != "" {
		if err := readFile(file); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("-d and -f must choose one")
	}
}
