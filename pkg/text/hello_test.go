package text_test

import (
	"os"
	"testing"

	"github.com/kou164nkn/greeting/pkg/text"
	"golang.org/x/text/language"
)

func setEnv(t *testing.T, kvs ...keyval) func() {
	original := make([]keyval, len(kvs))
	for i := range kvs {
		original[i].set(kvs[i].key(), os.Getenv(kvs[i].key()))
	}

	for i := range kvs {
		if err := os.Setenv(kvs[i].key(), kvs[i].val()); err != nil {
			t.Fatal("unexpected error:", err)
		}
	}

	return func() {
		for i := range original {
			if err := os.Setenv(original[i].key(), original[i].val()); err != nil {
				t.Fatal("unexpected error:", err)
			}
		}
	}
}

type keyval [2]string

func (kv *keyval) set(k, v string) {
	kv[0] = k
	kv[1] = v
}

func (kv keyval) key() string {
	return kv[0]
}

func (kv keyval) val() string {
	return kv[1]
}

func TestDefaultLang(t *testing.T) {
	cases := map[string]struct {
		keyval []keyval
		expect language.Tag
	}{
		"empty":          {[]keyval{{"LC_ALL", ""}, {"LC_MESSAGES", ""}, {"LANG", ""}}, language.English},
		"LC_ALL=ja":      {[]keyval{{"LC_ALL", "ja"}, {"LC_MESSAGES", ""}, {"LANG", ""}}, language.Japanese},
		"LC_MESSAGES=ja": {[]keyval{{"LC_ALL", ""}, {"LC_MESSAGES", "ja"}, {"LANG", ""}}, language.Japanese},
		"LANG=ja":        {[]keyval{{"LC_ALL", ""}, {"LC_MESSAGES", ""}, {"LANG", "ja"}}, language.Japanese},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if len(tt.keyval) != 0 {
				defer setEnv(t, tt.keyval...)()
			}

			if lang := text.DefaultLang(); lang != tt.expect {
				t.Errorf("want %v but got %v", tt.expect, lang)
			}
		})
	}
}

func TestGoodMorning(t *testing.T) {
	cases := map[string]struct {
		lang   language.Tag
		expect string
	}{
		"en": {language.English, "Good Morning"},
		"ja": {language.Japanese, "おはよう"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if actual := text.GoodMorning(tt.lang); actual != tt.expect {
				t.Errorf("want %v but got %v", tt.expect, actual)
			}
		})
	}
}

func TestHello(t *testing.T) {
	cases := map[string]struct {
		lang   language.Tag
		expect string
	}{
		"en": {language.English, "Hello"},
		"ja": {language.Japanese, "こんにちは"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if actual := text.Hello(tt.lang); actual != tt.expect {
				t.Errorf("want %v but got %v", tt.expect, actual)
			}
		})
	}
}

func TestGoodEvening(t *testing.T) {
	cases := map[string]struct {
		lang   language.Tag
		expect string
	}{
		"en": {language.English, "Good Evening"},
		"ja": {language.Japanese, "こんばんは"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if actual := text.GoodEvening(tt.lang); actual != tt.expect {
				t.Errorf("want %v but got %v", tt.expect, actual)
			}
		})
	}
}
