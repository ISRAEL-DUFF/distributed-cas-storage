package main

import (
	"fmt"
	"log"

	"github.com/israel-duff/distributed-cas-storage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello Duff!!!")

	select {}
}
