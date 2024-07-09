package i18n

import (
	"embed"
	"fmt"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var (
	bundle *i18n.Bundle
	//go:embed locales/*.yaml
	localeFS       embed.FS
	languagesFiles = map[language.Tag]string{
		language.English:             "locales/en.yaml",
		language.BrazilianPortuguese: "locales/pt-br.yaml",
	}
)

// CreateBundle creates a new i18n.Bundle with the given language
// Line for PR test purpuse 02
func CreateBundle(lang language.Tag) *i18n.Bundle {
	bundle = i18n.NewBundle(lang)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	for _, file := range languagesFiles {
		_, err := localeFS.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
	}
	selectedLanguage := languagesFiles[lang]
	bundle.LoadMessageFileFS(localeFS, selectedLanguage)

	return bundle
}

// GetBundle returns the i18n.Bundle Singleton
// CreateBundle must be called before calling this function
func GetBundle() (*i18n.Bundle, error) {
	if bundle == nil {
		return nil, fmt.Errorf("bundle is not created")
	}

	return bundle, nil
}
