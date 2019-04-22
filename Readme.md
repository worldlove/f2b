# f2b

f2b is a simple cli tool to build a static file or directory into a go language file. When the input is a file, it will set the file datas to a []byte type variable; when the input is a dir, it will set all the files in the dir to a map[key<string>]value<[]byte> variable that with the key of filename and value of file datas.

## Installation
    go get github.com/worldlove/f2b

## Usage
```
f2b version: 0.1
Usage: f2b [-h] [-d dirname] [-f filename] [-p package] [-v var]

Options:
  -d string
        Set the static dir, build a map[key<string>]value<[]byte> that with the key of filename and the value of filedata
  -f string
        Set the static file, build a var named of filename and value of the filedata
  -h    Help, print the usage
  -o string
        Set the output file name, default is ./{filename|dirname}.f2b.go
  -p string
        Set the package name, default the dir name or the file's basename
  -v string
        Set the var name, default the dir name or file name
```

## Version 0.1


## Licence
This project is licensed under the terms of the
[MIT license](https://github.com/callemall/material-ui/blob/master/LICENSE)



