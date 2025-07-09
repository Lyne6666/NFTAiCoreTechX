// cmd/nftaicoretechx/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"nftaicoretechx/internal/nftaicoretechx"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := nftaicoretechx.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
