package greeting

import "golang.org/x/text/language"

func ExportSetLang(l language.Tag) {
	orgLang := lang
	lang = l

	defer func() {
		lang = orgLang
	}()
}
