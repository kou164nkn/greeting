package greeting_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/kou164nkn/greeting"
	"golang.org/x/text/language"
)

func mockClock(t *testing.T, v string) greeting.Clock {
	t.Helper()
	now, err := time.Parse(time.RFC3339, v)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	return greeting.ClockFunc(func() time.Time {
		return now
	})
}

func TestGreetin_Do(t *testing.T) {
	greeting.ExportSetLang(language.Japanese)()

	g := greeting.Greeting{
		Clock: mockClock(t, "2021-04-01T07:00:00+09:00"),
	}

	var buf bytes.Buffer
	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error:", err)
	}

	if expected, actual := "おはよう", buf.String(); expected != actual {
		t.Errorf("greeting message want %s but got %s", expected, actual)
	}
}
