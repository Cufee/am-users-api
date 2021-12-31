package customization

const (
	BackgroundOptionKey = "background.image.url"
)

func NewBackgroundOption(backgroundImageURL string) InternalCustomizationOption {
	return InternalCustomizationOption{
		Key:             BackgroundOptionKey,
		Value:           backgroundImageURL,
		PlusRequired:    true,
		PremiumRequired: false,
	}
}
