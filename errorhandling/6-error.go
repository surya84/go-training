package main

import (
	"errors"
	"fmt"
	"os"
)

type Query struct {
	path string
	err  error
}

//var ErrFileNotFound = errors.New("not in the root directory")

//var ErrnotFound = errors.New("not found")

func (q *Query) Error() string {
	return q.path + " : " + " input " + q.err.Error()
}

func main() {

	var q *Query
	_, err := openFile1("test.txt", "root")
	if err != nil {
		//log.Println(err)
		if errors.As(err, &q) {
			fmt.Println("custom error found in chain ", q.path)
		}
	}
}

func openFile1(fileName string, dirName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, &Query{
			path: "cd drive",
			err:  err,
		}

	}
	return f, nil
}
