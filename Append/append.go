package main

import (
	"fmt"
)

func main() {
	a := []string{"ドラえもん", "のび太", "しずかちゃん", "ジャイアン", "スネ夫"}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, "ミニドラ", "出木杉")
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, "ジャイ子", "のび助", "玉子", "スネ吉")
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
