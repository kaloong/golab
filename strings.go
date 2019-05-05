package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	const placeOfInterest=`⌘`
	fmt.Printf("%+s\n", placeOfInterest)
	fmt.Printf("%+q\n", placeOfInterest)
	for i :=0; i<len(placeOfInterest); i++ {
		fmt.Printf("%x\n",placeOfInterest[i])
	}
}
