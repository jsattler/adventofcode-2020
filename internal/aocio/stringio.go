package aocio

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadFileSplit(fileName string, split string) []string {
	str := ReadFile(fileName)
	return strings.Split(str, split)
}

func ReadFile(fileName string) string {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(input)
}
