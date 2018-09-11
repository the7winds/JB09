package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type JsonDto map[string]string

func (dto JsonDto) has(key string) bool {
	_, ok := dto[key]
	return ok
}

type Action string
const (
	start Action = "start"
	message Action = "message"
)

func (dto JsonDto) isValid() bool {
	if !dto.has("action") {
		return false
	}

	action := Action(dto["action"])

	return action == start && dto.has("task") || action == message && dto.has("text")
}

func (dto JsonDto) action() Action {
	return Action(dto["action"])
}

func (dto JsonDto) task() string {
	return dto["task"]
}

func (dto JsonDto) text() string {
	return dto["text"]
}

func handle(line string) {
	lineReader := strings.NewReader(line)
	jsonDecoder := json.NewDecoder(lineReader)

	obj := JsonDto{}
	if err := jsonDecoder.Decode(&obj); err != nil {
		log.Println(err.Error())
		log.Println("can't parse objet")
		return
	}

	if !obj.isValid() {
		log.Println("invalid format")
		return
	}

	switch obj.action() {
	case start:
		fmt.Printf("> %s\n", obj.task())
	case message:
		fmt.Printf("# %s\n", obj.text())
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		handle(line)
	}
}
