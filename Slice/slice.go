package main

import (
    "fmt"
)

func main() {

    a := [5]string{"ドラえもん", "のび太", "しずかちゃん", "ジャイアン", "スネ夫"}
    b := a[:]
    c := a[2:4]

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    b[1] = "ドラミ"
    c[0] = "出木杉"

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}
