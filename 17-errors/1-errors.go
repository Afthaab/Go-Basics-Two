package main

import (
	"errors"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("not in the root directory")

type QueryError struct {
	Path string
	Err  error
}

func (q *QueryError) Error() string {
	return "main." + q.Path + ": " + "path " + q.Err.Error()
}

func main() {
	_, err := openFile("test.txt", "home")
	if err != nil {
		// ok := errors.Is(err, ErrFileNotFound)
		// if ok {
		// 	fmt.Println("Error found inside the chain ")
		// }
		log.Println(err)
		return
	}
}

func openFile(fileName string, dir string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, &QueryError{
			Path: dir,
			Err:  err,
		}
	}
	return f, nil
}
