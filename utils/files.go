package utils

import "io/ioutil"

func FilePathWalkDir(root string) []string {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files
	}
	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files
}
