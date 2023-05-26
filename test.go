package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

func (is IntSlice) String() string{
  var strs []string
	for _,v:=range is{
		strs=append(strs,  strconv.Itoa(v))
	}
	return "["+strings.Join(strs, ":")+"]"
}

func main() {
  var v IntSlice=[]int{1,2,3,4}
  fmt.Println(v)
	fmt.Println("Hello, 世界")
}
