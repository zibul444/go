package main

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	pass := "/home/e460/go/src/awesomeProject/booktasks/13-core-packages.go"
	//readeFile(pass)

	//fastread(pass)

	//readdir(pass)

	//walk(pass)

	//erro(pass)

	//hashFun(pass)
	hashFunsha1(pass)
}

func readeFile(pas string) {
	f, err := os.Open(pas)
	if err != nil {
		return
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	countReadByte, err := f.Read(bs)
	if err != nil {
		return
	}

	fmt.Println("countReadByte: ", countReadByte)
	fmt.Println("bs:\n", string(bs))
}

func fastread(pas string) {
	bs, err := ioutil.ReadFile(pas)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

func readdir(pas string) {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func walk(pas string) {
	filepath.Walk(
		".",
		func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			return nil
		})
}

func erro(pas string) error {
	err := errors.New("error message-1 " + pas)
	//panic(err)
	return err
}

func hashFun(s string) {
	h := crc32.NewIEEE()
	h.Write([]byte(s))
	v := h.Sum32()
	fmt.Println(v)
}

func hashFunsha1(s string) {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
