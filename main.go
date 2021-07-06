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
	fn := os.Args[1:]

	if len(fn) < 1 {
		fmt.Println("Input file")
	} else {
		for _, v := range fn {
			md5, err := hash_file_md5(v)
			if err != nil {
				log.Fatal(err)
			}

			sha1, err := hash_file_sha1(v)
			if err != nil {
				log.Fatal(err)
			}

			sha256, err := hash_file_sha256(v)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s ", v)
			fmt.Printf("%x ", md5)
			fmt.Printf("%x ", sha1)
			fmt.Printf("%x ", sha256)
			fmt.Printf("\n")
		}
	}
}

func hash_file_md5(v string) ([]byte, error) {
	var returnMD5String []byte

	f, err := os.Open(v)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	a := md5.New()
	if _, err := io.Copy(a, f); err != nil {
		return returnMD5String, err
	}
	returnMD5String = a.Sum(nil)

	return returnMD5String, nil
}

func hash_file_sha1(v string) ([]byte, error) {
	var returnSHA1String []byte

	f, err := os.Open(v)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	a := sha1.New()
	if _, err := io.Copy(a, f); err != nil {
		return returnSHA1String, err
	}
	returnSHA1String = a.Sum(nil)

	return returnSHA1String, nil
}

func hash_file_sha256(v string) ([]byte, error) {
	var returnSHA256String []byte

	f, err := os.Open(v)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	a := sha256.New()
	if _, err := io.Copy(a, f); err != nil {
		return returnSHA256String, err
	}
	returnSHA256String = a.Sum(nil)

	return returnSHA256String, nil
}
