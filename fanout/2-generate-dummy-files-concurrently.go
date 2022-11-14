package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const totalFile = 3000
const contentLength = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"),  "chapter-A.60-worker-pool")

type FileInfo struct {
	Index		int
	FileName	string
	WorkerIndex	int
	Err 		error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)


	// pipeline 1: job distribution
	chanFileIndex := generateFilesIndexes()

	// pipeline 2: the main logic (creating files)
	createFileWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFileWorker)

	// track and print output
	counterTotal := 0
	counterSuccess := 0
	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
		}else{
			counterSuccess++
		}
		counterTotal++
	}

	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

func generateFilesIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			chanOut <- FileInfo {
				Index:		i,
				FileName:	fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)


	// wait group to control the workers
	wg := new(sync.WaitGroup)


	// allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {
		// dispatch workers
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(WorkerIndex int) {

				// listen to chanIn channel for incoming jobs
				for job := range chanIn {
					// do the jobs
					filePath := filepath.Join(tempPath, job.FileName)
					content := randomString(contentLength)
					err := ioutil.WriteFile(filePath, []byte(content), os.ModePerm)


					log.Println("worker", workerIndex, "working on", job.FileName, "file generation")


					// construct the job's result , and send it to charOut
					chanOut <- FileInfo{
						FileName:	job.FileName,
						WorkerIndex:	workerIndex,
						Err:	err,
					}
				}


				// if chanIn is Closed, and the remainining jobs are finished
				// only then we mark the worker as complete
				wg.Done()
			}(workerIndex)
		}
	}()

	// wait until `chanIn` closed and then all workers are done,
	// because right after that - we need to close the `chanOut` channel.

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}