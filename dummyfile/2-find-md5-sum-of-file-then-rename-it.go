package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

func main() {
	log.Println("start")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "Seconds")
}

func proceed() {
	counterTotal := 0
	counterRenamed := 0
	err := filepath.Walk(tempPath, func(path string, info os.FileInfo, err error) error {
		// if there is an error, return immediately

		if err != nil {
			return err
		}

		// if it is a sub directory, return immediately
		if info.IsDir() {
			return nil
		}

		counterTotal++

		// Read file
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// sum it
		sum := fmt.Sprintf("%x", md5.Sum(buf))


		// rename file
		destinationPath := filepath.Join(tempPath, fmt.Sprintf("file-%s.txt", sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})

	if err != nil {
		log.Println("ERROR", err.Error())
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
}