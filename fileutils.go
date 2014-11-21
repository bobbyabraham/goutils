package goutils

import (
	"io/ioutil"
	"os"
	"strings"
)

type FileInfo struct {
	Name     string
	Basepath string
}

/*
createListOfAllFiles takes in a basepath for a directory and returns the
list of all files present within that directory and all its sub directories
recursively
*/
func CreateListOfAllFiles(path string) ([]FileInfo, error) {
	acc := make([]FileInfo, 0)
	if filestat, err := os.Stat(path); err == nil {
		inpList := []os.FileInfo{filestat}
		acc = createFileListHelper(&acc, inpList,
			path[:strings.LastIndex(path, "/")],
			func(item os.FileInfo) bool {
				return true
			})
	} else {
		return nil, err
	}
	return acc, nil

}

/*
createFileListHelper is a helper for createFileListHelper function. It would be
ideal to have this helper function defined as an inner function in
createListOfAllFiles function. However, go does not allow inner function
definitions - go only allows functions to be assigned to a variable - but
this does not allow a recursive call to the same function
acc - is a pointer to a slice - which is an accumulator. The initial call to
this function should pass in a pointer to an empty slice. On return acc will
have all the
*/
func createFileListHelper(acc *[]FileInfo, inpList []os.FileInfo,
	basepath string, predicate func(os.FileInfo) bool) []FileInfo {
	for _, item := range inpList {
		if item.IsDir() {
			subdirFilesInfo, _ := ioutil.ReadDir(basepath +
				"/" + item.Name())
			createFileListHelper(acc, subdirFilesInfo,
				basepath+"/"+item.Name(), predicate)
		} else {
			*acc = append(*acc, FileInfo{item.Name(), basepath})
		}
	}
	return *acc
}
