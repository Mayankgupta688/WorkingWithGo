package main

import(
	"fmt"
	"os"
	"io/ioutil"
)

const watchedPath = "./source";

func main() {

	watcher, _ := os.Open(watchedPath)
	files, _ := watcher.Readdir(-1)

	for _, file := range files {
		func(file) {
			filePath := watchedPath + "/" + file.Name()
			fileOpen, _ := os.Open(filePath)
			fileData, _ := ioutil.ReadAll(fileOpen)
			fileOpen.Close()
			os.Remove(filePath)
			fmt.Println(string(fileData))
		}(file)
	}
}