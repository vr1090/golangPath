package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	pos := rand.Intn(len(formats))

	if name == "" {
		return "", errors.New("should have name")
	}

	return fmt.Sprintf(formats[pos], name), nil
}
