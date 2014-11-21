package goutils

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
)

// Computes and returns the checksum of of the given file (name)
func computeChecksum(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file - " + fileName)
		return "", err
	}
	checksum := sha1.Sum(data)
	sd := fmt.Sprintf("%x", checksum)

	return sd, nil
}
