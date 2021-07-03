package enumerator

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

type SearchObjectKind int

const (
	Interface SearchObjectKind = iota
	Struct
)

var regexpMap = map[SearchObjectKind]string{
	Interface: `type\s*(\w*)\s*interface\s*{`,
	Struct:    `type\s*(\w*)\s*struct\s*{`,
}

func enumerateFolder(path string, callback func(fileName string)) {
	filepath.WalkDir(path,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			fileInfo, err := os.Stat(path)
			if err != nil {
				return err
			}
			if fileInfo.IsDir() {
				return nil
			}
			callback(path)
			return nil
		},
	)
}

func Enumerate(objectKind SearchObjectKind, rootFolder string) {
	callback := func(fileName string) {
		regex, err := regexp.Compile(regexpMap[objectKind])
		if err != nil {
			fmt.Println(err)
		}

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var allMatches []string
		for scanner.Scan() {
			matches := regex.FindAllString(scanner.Text(), -1)
			allMatches = append(allMatches, matches...)
		}
		if len(allMatches) != 0 {
			fmt.Printf("File: %s\n", fileName)
			for _, val := range allMatches {
				fmt.Printf("\t%s\n", val)
			}
			fmt.Println()
		}
	}

	enumerateFolder(rootFolder, callback)
}
