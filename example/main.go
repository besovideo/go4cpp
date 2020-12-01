package main

import (
	"fmt"
	"github.com/besovideo/go4cpp"
	"log"
	"time"
)

func main() {
	go4cpp.InitLibrary([]byte("go data"), func(err error, data []byte) {
		log.Printf("=== %v %v", err, string(data))
	})

	var data = fmt.Sprintf("%v", time.Now().Format(time.RFC3339))
	go4cpp.Command([]byte(data), func(err error, data []byte) {
		log.Printf("+++ %v %v", err, string(data))
	})

	go4cpp.ReleaseLibrary()
}
