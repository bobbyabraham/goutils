/*
Utility package - provides commonly used utility functions
*/
package goutils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
PerformHttpPutRequest - self explanatory
sends an http put reuest to the supplied url (using the supplied user name and
password for basic http authentication - if needed) with the supplied data as
payload. The function returns the http return status code, the return payload
as a byte array
*/
func PerformHttpPutRequest(url, username, password string,
	dataToSend []byte) (int, []byte, error) {
	client := http.Client{}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(dataToSend))
	if err != nil {
		fmt.Println("Unable to put file to Artifactory :" + err.Error())
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error accessing url - " + url + ":" + err.Error())
		return resp.StatusCode, nil, errors.New(err.Error())
	}
	b, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, b, nil

}

/*
PerformHttpGetRequest- self explanatory
sends an http get request to the supplied url (using the supplied user name and
password for basic http authentication - if needed). The function returns
the http return status code, the return payload as a byte array
*/
func PerformHttpGetRequest(url, username, password string) (int, []byte, error) {
	client := http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error accessing url - " + url + ":" + err.Error())
		return resp.StatusCode, nil, errors.New(err.Error())
	}
	b, _ := ioutil.ReadAll(resp.Body)

	return resp.StatusCode, b, nil

}
