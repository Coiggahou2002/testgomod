package testmod

import (
	"fmt"
	"errors"
)


func Hi(name string, lang string) (string, error) {
	switch lang {
	case "cn":
		return fmt.Sprintf("你好，%s！", name), nil
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s!", name), nil
	default:
		return "", errors.New("unknown language")

	}
}