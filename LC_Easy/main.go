package main

import (
	"fmt"
	"strconv"
	// shortener "github.com/Richd0tcom/gophercises/LC_Easy/LC_Apple"))
)
func IsPalindrome(x int) bool {
//121
var isTrue bool = true
str:= strconv.Itoa(x)
length:= len(str)/2

for i:=0; i<=length;i++{
	if str[i]!=str[len(str)-(i+1)]{
		isTrue= false
		if !isTrue{
			break
		}
	}
}

return isTrue
}

func main(){

	fmt.Println(IsPalindrome(-121))
}