package expand

type config struct {
	replaceMapKey   bool
	quoteCollection bool
}

type Option func(*config) error

// ReplaceMapKey - Replace map key of YAML.
func ReplaceMapKey() Option {
	return func(c *config) error {
		c.replaceMapKey = true
		return nil
	}
}

// QuoteCollection - Quotes the replaced value as a string if it is a map or slice.
func QuoteCollection() Option {
	return func(c *config) error {
		c.quoteCollection = true
		return nil
	}
}
