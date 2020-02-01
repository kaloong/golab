package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{ 
		"Dalice":31, 
		"Charlie":34, 
	}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings( names )
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
		fmt.Printf("%d\n",&names)
}
