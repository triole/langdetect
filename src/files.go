package main

import "io/ioutil"

func readFile(fileName string) (content string) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	content = string(b)
	return
}
