package main

import (
	"fmt"
)

func main() {
	a := make([]string, 7)
	a[0] = "ドラえもん"
	a[1] = "のび太"
	a[2] = "しずかちゃん"
	a[3] = "ジャイアン"
	a[4] = "スネ夫"
	a[5] = "ミニドラ"
	a[6] = "出木杉"

	b := append(a[:5], "ジャイ子", "のび助")
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	b = append(a[:5], "玉子", "スネ吉", "ジャイ子", "のび助")
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	b[5] = "サザエさん"
	b[6] = "しんのすけ"
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
}
