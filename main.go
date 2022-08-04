package main

import (
	"fmt"
	"math"
)

func main() {
	var totalMinute, afterMod, mod, total int
	var minute float64

	totalMinute = 130
	minute = 60

	mod = int(math.Mod(float64(totalMinute), minute))
	afterMod = totalMinute - mod
	total = afterMod / int(minute)

	if total == 0 {
		fmt.Println(mod, "นาที")
	} else {
		fmt.Println(total, "ชั่วโมง", mod, "นาที")
	}
  // 2 ชั่วโมง 10 นาที
}
