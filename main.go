package main

import (
	"flag"
	"log"
	"os"

	"github.com/jmattaa/cheatsh/ai"
	"github.com/jmattaa/cheatsh/formatter"
)

func main() {
	var topic, outputfile string
	flag.StringVar(&topic, "topic", "", "get a cheat sheet on this topic")
	flag.StringVar(&outputfile, "o", "", "write to this file")
	flag.Parse()

	if topic == "" {
		println("usage: cheatsh -topic <topic>")
		log.Fatalln("man give me a topic")
	}

	println("getting a cheat sheet on:", topic)

	res := ai.GetCheatSheet(topic)
	formatter.Print(res)

	// write to file
    if outputfile == "" {
        return
    }

    err := os.WriteFile(outputfile, []byte(res), 0644)
    if err != nil {
        log.Fatal(err)
    }
}
