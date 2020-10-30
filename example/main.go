package main

import (
	"fmt"
	"github.com/besovideo/go4cpp"
	"log"
	"time"
)

func main() {
	go4cpp.InitLibrary(func(data []byte) {
		log.Printf("=== %v", string(data))
	})

	var data = fmt.Sprintf("%v", time.Now().Format(time.RFC3339))
	go4cpp.Command([]byte(data), func(data []byte) {
		log.Println("+++" + string(data))
	})

	go4cpp.ReleaseLibrary()
}
