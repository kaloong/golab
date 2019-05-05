package main
import (
	"fmt"
	"time"
)
func main() {
	t:= time.Now()
	switch {
	case t.Hour() < 12: 
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}
	fmt.Printf("%s\n",t)
}
