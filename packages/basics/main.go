// packages/basics/main.go
package main

import (
	"fmt"
	"time"
)
func main() {

	fmt.Println("Hello Gopher!")
   
	fmt.Printf("Today is %s",time.Now().Weekday())
}
