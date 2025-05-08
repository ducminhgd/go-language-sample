package translation

import (
	"errors"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/ducminhgd/go-language-sample/internal/require/translation"
	"github.com/leonelquinteros/gotext"
)

// DefaultLanguage is the fallback language
const DefaultLanguage = "en"

// Service implements the translation.Service interface
type Service struct {
	localesDir string
	domains    map[string]*gotext.Domain
	mu         sync.RWMutex
}

// NewService creates a new translation service
func NewService(localesDir string) translation.Service {
	return &Service{
		localesDir: localesDir,
		domains:    make(map[string]*gotext.Domain),
	}
}

// Translate translates a key to the specified language
func (s *Service) Translate(key, lang string) (string, error) {
	if lang == "" {
		lang = DefaultLanguage
	}

	domain, err := s.getDomain(lang)
	if err != nil {
		// Fallback to default language
		domain, err = s.getDomain(DefaultLanguage)
		if err != nil {
			return "", fmt.Errorf("failed to get domain for default language: %w", err)
		}
	}

	translation := domain.Get(key)
	if translation == key {
		// If the translation is the same as the key, it might not be found
		// Try with default language
		if lang != DefaultLanguage {
			defaultDomain, err := s.getDomain(DefaultLanguage)
			if err == nil {
				translation = defaultDomain.Get(key)
			}
		}
	}

	return translation, nil
}

// getDomain gets or creates a domain for the specified language
func (s *Service) getDomain(lang string) (*gotext.Domain, error) {
	s.mu.RLock()
	domain, ok := s.domains[lang]
	s.mu.RUnlock()

	if ok {
		return domain, nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Check again in case another goroutine created it while we were waiting
	domain, ok = s.domains[lang]
	if ok {
		return domain, nil
	}

	// Create a Po instance and parse the file
	po := gotext.NewPo()
	poPath := filepath.Join(s.localesDir, lang, "message.po")

	// ParseFile doesn't return an error, so we need to check if the file exists
	po.ParseFile(poPath)

	// Get the domain from the Po object
	domain = po.GetDomain()

	// Check if any translations were loaded
	if len(domain.GetTranslations()) == 0 {
		return nil, errors.New("language not supported")
	}

	s.domains[lang] = domain
	return domain, nil
}
