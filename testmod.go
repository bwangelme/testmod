package testmod

import "fmt"

// Hi returns a friendly greet
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}
