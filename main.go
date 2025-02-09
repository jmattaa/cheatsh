package main

import (
	"flag"
	"log"

	"github.com/jmattaa/cheatsh/ai"
	"github.com/jmattaa/cheatsh/formatter"
)

func main() {
	var topic string
	flag.StringVar(&topic, "topic", "", "get a cheat sheet on this topic")
	flag.Parse()

	if topic == "" {
		println("usage: cheatsh -topic <topic>")
		log.Fatalln("man give me a topic")
	}

	println("getting a cheat sheet on:", topic)

    res := ai.GetCheatSheet(topic)

    formatter.Print(res)
}

