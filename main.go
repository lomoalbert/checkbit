package main

import (
	"flag"
	"fmt"
	"os"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func read(finame string) string {
	fi, err := os.Open(finame)
	checkerr(err)
	defer fi.Close()
	b := make([]byte, 4)
	_, err = fi.ReadAt(b, 60)
	checkerr(err)
	offset := int(b[0]) + int(b[1])*256 + int(b[2])*256*256 + int(b[3])*256*256*256
	value := make([]byte, 2)
	_, err = fi.ReadAt(value, int64(offset)+4)
	checkerr(err)
	return string(value)

}

func main() {
	flag.Parse()
	if flag.NArg() <1{
		fmt.Println("please put finename in!")
		return
	}
	finame := flag.Arg(0)
	value := read(finame)
	switch value {
	case string([]byte{100, 134}):
		fmt.Println("amd64")
	case string([]byte{76, 1}):
		fmt.Println("win32")
	default:
		fmt.Println("not PE program")
		fmt.Println([]byte(value))
	}
}
