package weatherservice

type ProviderType string

const (
	PROVIDER_YAHOO       = "YAHOO"
	PROVIDER_OPENWEATHER = "OPEN"
)

//Provider a weather service object
type Provider struct {
	Type ProviderType `json:"name"`
	URL  string       `json:"url"`
}

//NewProvider returns a new weahter service object
func NewProvider(pType ProviderType, url string) *Provider {
	return &Provider{
		Type: pType,
		URL:  url,
	}
}
