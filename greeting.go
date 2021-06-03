package greeting

import (
	"fmt"
	"time"

	"github.com/kou164nkn/greeting/pkg/text"
)

var lang = text.DefaultLang()

// Do関数は挨拶文を返します.
// 以下のように指定した時刻によって返す文を変えます.
// 04:00-09:59: おはよう
// 10:00-16:59: こんにちは
// 17:00-03:59: こんばんは
func Do() {
	h := time.Now().Hour()
	switch {
	case h >= 4 && h <= 9:
		fmt.Println(text.GoodMorning(lang))
	case h >= 10 && h <= 16:
		fmt.Println(text.Hello(lang))
	default:
		fmt.Println(text.GoodEvening(lang))
	}
}
