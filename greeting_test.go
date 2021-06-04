package greeting_test

import (
	"bytes"
	"errors"
	"io"
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

type errorWriter struct {
	Err error
}

func (w *errorWriter) Write(p []byte) (n int, err error) {
	return 0, w.Err
}

func TestGreetin_Do(t *testing.T) {
	greeting.ExportSetLang(language.Japanese)()

	cases := map[string]struct {
		writer io.Writer
		clock  greeting.Clock

		msg       string
		expectErr bool
	}{
		"04時": {new(bytes.Buffer), mockClock(t, "2021-04-01T04:00:00+09:00"), "おはよう", false},
		"09時": {new(bytes.Buffer), mockClock(t, "2021-04-01T09:00:00+09:00"), "おはよう", false},
		"10時": {new(bytes.Buffer), mockClock(t, "2021-04-01T10:00:00+09:00"), "こんにちは", false},
		"16時": {new(bytes.Buffer), mockClock(t, "2021-04-01T16:00:00+09:00"), "こんにちは", false},
		"17時": {new(bytes.Buffer), mockClock(t, "2021-04-01T17:00:00+09:00"), "こんばんは", false},
		"03時": {new(bytes.Buffer), mockClock(t, "2021-04-01T03:00:00+09:00"), "こんばんは", false},
		"エラー": {&errorWriter{Err: errors.New("error")}, nil, "", true},
	}

	for name, tt := range cases {
		tt := tt

		t.Run(name, func(t *testing.T) {
			g := greeting.Greeting{
				Clock: tt.clock,
			}

			switch err := g.Do(tt.writer); true {
			case err == nil && tt.expectErr:
				t.Error("expected error did not occur")
			case err != nil && !tt.expectErr:
				t.Error("unexpected error:", err)
			}

			if buf, ok := tt.writer.(*bytes.Buffer); ok {
				if msg := buf.String(); tt.msg != msg {
					t.Errorf("greeting message want %s but got %s", tt.msg, msg)
				}
			}
		})
	}
}
