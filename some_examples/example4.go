package main

import (
	"fmt"
)

func main() {
	s := []int{5,6,7,8,9}
	for i, j:=0, len(s)-1; i<j; i,j=i+1, j-1 {
		fmt.Printf("i=[%d],%d j=[%d],%d\n",i,s[i],j,s[j])
		s[i], s[j]=s[j], s[i]
	}
	fmt.Println(s)
}
