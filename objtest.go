package main

import ( 
	"fmt"
	"myobj/obj"
)

func main() {
	var o obj.Obj
	o.Set("kaloong")
	fmt.Printf("%s\n",o.Get())
	fmt.Printf("%s\n",o.Name)
}
