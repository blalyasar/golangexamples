package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// message := fmt.Sprintf("Hi %v Welcome",name)
	// create  meesage fromrandom format
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat()) // for test ...

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// a map associate names with messages
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// ın the map associate the retrieved
		// the name
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi %v Welcome!",
		"Great to see you %v!",
		"Hail %v well met",
	}
	return formats[rand.Intn(len(formats))]
}
