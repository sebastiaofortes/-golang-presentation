package main

import "fmt"



func main() {
 go d("primeiro")
 go d("segundo")
}

func d(name string){
 fmt.Println(name)
}