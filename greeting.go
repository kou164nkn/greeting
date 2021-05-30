package greeting

import (
	"fmt"
	"time"
)

// Do関数は挨拶文を返します.
// 以下のように指定した時刻によって返す文を変えます.
// 04:00-09:59: おはよう
// 10:00-16:59: こんにちは
// 17:00-03:59: こんばんは
func Do() {
	h := time.Now().Hour()
	switch {
	case h >= 4 && h <= 9:
		fmt.Println("おはよう")
	case h >= 10 && h <= 16:
		fmt.Println("こんにちは")
	default:
		fmt.Println("こんばんは")
	}
}
