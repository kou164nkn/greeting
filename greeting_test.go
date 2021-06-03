package greeting_test

import (
	"bytes"
	_ "fmt"
	"testing"
	"time"

	"github.com/kou164nkn/greeting"
	"golang.org/x/text/language"
)

func TestGreetin_Do(t *testing.T) {
	greeting.ExportSetLang(language.Japanese)

	g := greeting.Greeting{
		Clock: greeting.ClockFunc(func() time.Time {
			return time.Date(2021, 6, 1, 06, 0, 0, 0, time.Local)
		}),
	}

	var buf bytes.Buffer
	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error:", err)
	}

	if expected, actual := "おはよう", buf.String(); expected != actual {
		t.Errorf("greeting message want %s but got %s", expected, actual)
	}
}
