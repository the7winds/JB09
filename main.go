package main

import (
	"awesomeProject/object"
	"bufio"
	"fmt"
	"log"
	"os"
)

func handle(line string) {
	obj, err := object.Parse(line);
	if err != nil {
		log.Println(err.Error())
		return
	}

	switch obj.Action() {
	case object.Start:
		fmt.Printf("> %s\n", obj.Task())
	case object.Message:
		fmt.Printf("# %s\n", obj.Text())
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		handle(line)
	}
}
