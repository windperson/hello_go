package main

import "fmt"

func printWow(times int){
    for i:=0;i<times;i++{
        fmt.Println("Wow")
    }
}

func max(i int, j int) int {
    if (i>j){
        return i
    }else {
        return j
    }
}

func maxWithPointer(i int, j int, k *int) int{
    if(i>j){
        *k = i
    }else {
        *k = j
    }
    fmt.Println("k inside func=",k)
    return *k
}

func main() {
    var i,j = 100,15
    var ret int
    fmt.Println("&ret=",&ret)
    // fmt.Println("max of ",i," and ",j," is ",max(i,j))
    fmt.Println("max of ",i," and ",j," is ",maxWithPointer(i,j, &ret))    
}