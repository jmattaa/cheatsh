package main

import (
	"flag"
	"log"
	"os"

	"github.com/jmattaa/cheatsh/ai"
	"github.com/jmattaa/cheatsh/mdrenderer"
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

	var file *os.File
	var err error
	if outputfile != "" {
		file, err = os.Create(outputfile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	for chunk := range ai.GetCheatSheet(topic) {
		mdrenderer.Print(chunk) 

		if file != nil {
			_, err := file.WriteString(chunk)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

    // just add an newline at the end
    println("");
}
