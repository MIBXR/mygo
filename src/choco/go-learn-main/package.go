package main

import "fmt"

import "choco/greetings"

import choco_utils "choco/utils"

func main() {
    fmt.Println(greetings.Hello("Choco Xie"))
    fmt.Println(choco_utils.RandomBetween(1, 100))
    fmt.Println(choco_utils.CopyFile("/home/xr/go/src/choco-go-learn-main/", "/home/xr/go/src/choco-go-learn-copy/", false))
}