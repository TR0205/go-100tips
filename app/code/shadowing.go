package code

import (
	"log"
)

// varで変数宣言しているdayOfWeekはif文の中で再度:=を使用して宣言可能。これを変数のshadowingという
// この場合、最初に宣言されているdayOfWeekは使用されないため注意
func shadowingNG() {
	var isWeekDay bool
	// var dayOfWeek string

	if isWeekDay {
		dayOfWeek := "平日"
		log.Println(dayOfWeek)
	} else {
		dayOfWeek := "祝日"
		log.Println(dayOfWeek)
	}
}

// 以下のように変数を再度他の変数に入れてから、最初に宣言した変数に再度入れてあげるとエラーは発生しない
func ShadowingOK() {
	var isWeekDay bool
	var dayOfWeek string

	if isWeekDay {
		d := "平日"
		dayOfWeek = d
		log.Println(dayOfWeek)
	} else {
		d := "祝日"
		dayOfWeek = d
		log.Println(dayOfWeek)
	}
}
