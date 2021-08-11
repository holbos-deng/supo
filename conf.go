package main

import (
    "fmt"
    "github.com/holbos-deng/readconf"
)

func main() {
    conf := rdc.New("C:\\go\\gopath\\easonsu\\rdc\\test.yaml")

    a := conf.Get("steve.age").Value()     // 35
    b := conf.Get("steve.hobbies").Value() // [skateboarding snowboarding go ]
    steve := conf.Get("steve")
    c := steve.Get("clothing").Get("jacket").Value() // leather
    d := steve.Get("clothing").Get("jacket").Key()   // steve.clothing.jacket
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d
}
