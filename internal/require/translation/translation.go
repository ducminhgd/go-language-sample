package translation

// Service defines the interface for translation services
type Service interface {
	// Translate translates a key to the specified language
	Translate(key, lang string) (string, error)
}
