package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fn := os.Args[1]
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	md5 := h.Sum(nil)

	h = sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	sha1 := h.Sum(nil)

	h = sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	sha256 := h.Sum(nil)

	fmt.Printf("%s ", fn)
	fmt.Printf("%x ", md5)
	fmt.Printf("%x ", sha1)
	fmt.Printf("%x ", sha256)
}
