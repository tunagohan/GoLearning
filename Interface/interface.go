package main

import (
	"fmt"
	"strconv"
)

type Human interface {
	run(int) string
	stop()
}

type User struct {
	name  string
	sex   int
	speed int
}

func (u *User) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "Kmで走るお"
}

func (u *User) stop() {
	fmt.Println("キキーッ(止まった")
	u.speed = 0
}
func main() {
	kazutomo := &User{name: "一友", speed: 0}

	var i Human = kazutomo
	fmt.Println(i.run(50))
	i.stop()
}
