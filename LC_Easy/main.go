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

	// fmt.Println(IsPalindrome(-121))
	fmt.Println("yeah")
	fmt.Println(FindJudge(3, [][]int{{1,3},{2,3}}))
}

func FindJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n ==1 {
		return 1
	}
	
	count := make([]int, int(n+1)) 
	fmt.Println(count)
	for _, val := range trust{
		count[val[0]]++
		count[val[1]]++
	}
	for i := range count {
		if count[i] == n-1{
			return i
		}
	}
	return -1
}

func maxSubArray(nums []int) int {
    return 1
}

//divide and conquer algorithm