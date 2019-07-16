package main
import "strings"
import "fmt"

func main() {
	dataPhrase := "Hello World I am Mayank Gupta"

	words := strings.Split(dataPhrase, " ")

	channelBufferData := make(chan string, len(words))

	for _, word := range words {
		channelBufferData <- word
	}

	close(channelBufferData)

	for channelMessage := range channelBufferData {
		fmt.Println(channelMessage)
	}
}