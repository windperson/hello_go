package main

import "fmt"

type flags byte

const (
    read flags = 1 << iota
    write
    exec
)

func main(){
    fileType := read | exec
    fmt.Printf("%b",fileType)
}