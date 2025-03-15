package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo a un nombre
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	// devuelve un saludo que incluye el hnombre en un mensaje
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos devuelve un mapa que asocia nombres con mensajes de saludo
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// randomFormat devuelve uno de varios formatos de mensaje de saludo
// se selecciona de forma aleatoria
func randomFormat() string {
	formats := []string{
		"Hi, %v! Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
