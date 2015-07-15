package looper

import (
	"fmt"
	"strconv"
)

func Looper(n int) string {
	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Printf("POWERLOOP CYCLE %d\n", i)
		}
		return strconv.Itoa(n)
	}
	return "WHAT IS THIS YA JABRONEY"
}
