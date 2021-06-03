package greeting

import (
	"fmt"
	"io"
	"time"

	"github.com/kou164nkn/greeting/pkg/text"
)

var lang = text.DefaultLang()

// テスタビリティを考慮し,
// Now()メソッドを定義したClockインターフェースを用意
type Clock interface {
	Now() time.Time
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}

type Greeting struct {
	Clock Clock
}

func (g *Greeting) now() time.Time {
	if g.Clock == nil {
		return time.Now()
	}
	return g.Clock.Now()
}

// Do関数は挨拶文を返します.
// 以下のように指定した時刻によって返す文を変えます.
// 04:00-09:59: おはよう
// 10:00-16:59: こんにちは
// 17:00-03:59: こんばんは

// 引数にio.Writerを取ることで出力先を自由に変えることができる
func (g *Greeting) Do(w io.Writer) error {
	h := g.now().Hour()

	var msg string
	switch {
	case h >= 4 && h <= 9:
		msg = text.GoodMorning(lang)
	case h >= 10 && h <= 16:
		msg = text.Hello(lang)
	default:
		msg = text.GoodEvening(lang)
	}

	_, err := fmt.Fprint(w, msg)
	if err != nil {
		return err
	}

	return nil
}
