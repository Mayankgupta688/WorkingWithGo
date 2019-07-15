package main

import(
	"fmt"
	"os"
	"io/ioutil"
)

const watchedPathData = "./fileDropper";

func fileReaderApplicationWithGoRoutine() {

	watcher, _ := os.Open(watchedPath)
	files, _ := watcher.Readdir(-1)
	fileData := ""

	for _, file := range files {
		func(fileData) {
			filePath := watchedPathData + "/" + file.Name()
			fileOpen, _ := os.Open(filePath)
			fileData, _ := ioutil.ReadAll(fileOpen)
			fileOpen.Close()
			os.Remove(filePath)
			fmt.Println(string(fileData))
		}(file)
	}
}