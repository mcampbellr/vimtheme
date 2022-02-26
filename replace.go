package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func SearchAndReplace(search string, replaceStr string, fileName string, debug bool) error {
	home := os.Getenv("HOME")
	path := filepath.Join(home, fileName)

	file, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	out := string(file)
	re := regexp.MustCompile(search)

	newContent := []byte(re.ReplaceAllString(out, replaceStr))

	if debug {
		fmt.Println(string(newContent))
	}

	err = ioutil.WriteFile(path, newContent, 0644)
	if err != nil {
		panic(err)
	}

	return err
}
