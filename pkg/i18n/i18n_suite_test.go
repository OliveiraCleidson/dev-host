package i18n_test

import (
	"log"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	sut "github.com/oliveiracleidson/devhost/pkg/i18n"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/text/language"
)

func TestI18n(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "I18n Suite")
}

func getTranslation(message *i18n.Message, localizer *i18n.Localizer) string {
	config := &i18n.LocalizeConfig{
		MessageID:    message.ID,
		TemplateData: nil,
	}
	translation, err := localizer.Localize(config)
	if err != nil {
		log.Fatal(err)
	}
	return translation
}

var _ = Describe("I18n Unit Test", func() {
	var bundle *i18n.Bundle

	BeforeEach(func() {
		bundle = sut.CreateBundle(language.BrazilianPortuguese)
	})

	Context("when creating a bundle", func() {
		It("should not be nil", func() {
			Expect(bundle).ToNot(BeNil())
		})

		It("should load the correct translations", func() {
			localizer := i18n.NewLocalizer(bundle, "pt")
			simpleMessage, err := localizer.Localize(&i18n.LocalizeConfig{
				MessageID: "cli.name",
			})
			Expect(err).To(BeNil())
			Expect(simpleMessage).To(Equal("Dev Setup"))
		})
	})

	Context("when getting the bundle", func() {
		It("should return the bundle if it is created", func() {
			b, err := sut.GetBundle()
			Expect(err).To(BeNil())
			Expect(b).ToNot(BeNil())
		})

	})
})
