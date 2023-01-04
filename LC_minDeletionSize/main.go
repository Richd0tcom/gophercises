package main

import (
	"fmt"
	"sort"
	// "time"
	"sync"
)

func minDeletionSize(strs []string) int {
	var wg sync.WaitGroup
	//create a minDeletionSize counter
	var counter int = 0
	//get the length of the strings from the first element in the array
	n := len(strs[0])
	

	//write a method that gets all the individual characters and puts them in a slice
	for i := 0; i < n; i++ {
		wg.Add(1)
		fmt.Println("current loop index: ",i)
        go func(index int){
			defer wg.Done()
			fmt.Println("current routine index: ",index)
			c:= make([]string,len(strs))
			for _, val :=range strs {
				c = append(c, string(val[index]))
			}
			fmt.Println(sort.StringsAreSorted(c))
			if sort.StringsAreSorted(c) {
				return
			}else{
				counter++
				return
			}
			
		}(i)
    }
	wg.Wait()
	//check if the resulting slice is sorted. if not sorted increment the minDeletionsize counter
    return counter
}

func main() {
	
	fmt.Println("starting...")
	fmt.Println("I returned here",minDeletionSize([]string{"cba","daf","ghi"}))
	// time.Sleep(5 * time.Second)
	fmt.Println(twoSum([]int{3,2,4}, 6))
	
}

func twoSum(nums []int, target int) []int {

    m:=make(map[int]int)

	for i, val:= range nums{
		index, ok:= m[target-val]
		if ok {
			return []int{index, i}
		}
		m[val]=i
	}
	return nil
}