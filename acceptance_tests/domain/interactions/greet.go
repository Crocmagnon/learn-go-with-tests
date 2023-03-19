package interactions

import "fmt"

func Greet(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello, %s", name)
}

func Curse(name string) string {
	return fmt.Sprintf("Go to hell, %s!", name)
}
