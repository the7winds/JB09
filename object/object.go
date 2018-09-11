package object

import (
	"encoding/json"
	"errors"
	"strings"
)

type Object map[string]string

// keys
const (
	action = "action"
	text = "text"
	task = "task"
)

type Action string
const (
	Start Action = "start"
	Message Action = "message"
)

func (obj Object) has(key string) bool {
	_, ok := obj[key]
	return ok
}

func (obj Object) isValid() bool {
	if !obj.has(action) {
		return false
	}

	action := obj.Action()

	return action == Start && obj.has(task) || action == Message && obj.has(text)
}

func Parse(str string) (Object, error) {
	lineReader := strings.NewReader(str)
	jsonDecoder := json.NewDecoder(lineReader)

	obj := Object{}

	if err := jsonDecoder.Decode(&obj); err != nil {
		return nil, err
	}

	if obj.isValid() {
		return nil, errors.New("invalid format")
	}

	return obj, nil
}

func (obj Object) Action() Action {
	return Action(obj[action])
}

func (obj Object) Task() string {
	return obj[task]
}

func (obj Object) Text() string {
	return obj[text]
}
