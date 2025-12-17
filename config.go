package hr32

type Config struct {
	alphabet   string
	generators []int
	separator  rune

	version  int
	versions map[int]struct{}

	prefix string
	maxLen int
	minLen int

	excludePrefix bool

	// internal
	initialized bool
}

// Option
/*
	- WithAlphabet
	- WithGenerators
	- WithSeparator
	- WithVersion
	- WithMinLen
	- WithMaxLen
	- WithPrefix
	- WithExcludePrefix
*/
type Option func(*Config)

func WithAlphabet(value string) Option {
	return func(cfg *Config) { cfg.alphabet = value }
}
func WithGenerators(value []int) Option {
	return func(cfg *Config) { cfg.generators = value }
}
func WithSeparator(value rune) Option {
	return func(cfg *Config) { cfg.separator = value }
}
func WithVersion(value int) Option {
	return func(cfg *Config) { cfg.version = value }
}
func WithMaxLen(value int) Option {
	return func(cfg *Config) { cfg.maxLen = value }
}
func WithMinLen(value int) Option {
	return func(cfg *Config) { cfg.minLen = value }
}
func WithPrefix(value string) Option {
	return func(cfg *Config) { cfg.prefix = value }
}
func WithExcludePrefix(value bool) Option {
	return func(cfg *Config) { cfg.excludePrefix = value }
}

func (m *Config) Clone(options ...Option) (*Config, error) {
	cfg := *m

	for _, option := range options {
		option(&cfg)
	}

	if len(cfg.alphabet) != 32 {
		return nil, ErrConfig{Field: "alphabet", Description: "must be 32 characters long"}
	}

	charMap := make(map[rune]int, len(cfg.alphabet))
	for i, chr := range cfg.alphabet {
		charMap[chr] = i

		if cfg.separator != 0 && cfg.separator == chr {
			return nil, ErrConfig{Field: "separator", Description: "can't be in alphabet"}
		}
	}
	if len(charMap) != len(cfg.alphabet) {
		return nil, ErrConfig{Field: "alphabet", Description: "must contain unique characters"}
	}

	if len(cfg.generators) != 5 {
		return nil, ErrConfig{Field: "generators", Description: "must be 5 values"}
	}

	var gens = make(map[int]struct{}, len(cfg.generators))
	for _, chr := range cfg.generators {
		gens[chr] = struct{}{}
	}
	if len(gens) != len(cfg.generators) {
		return nil, ErrConfig{Field: "generators", Description: "must contain unique values"}
	}

	if cfg.separator == 0 && cfg.prefix != "" {
		return nil, ErrConfig{Field: "separator", Description: "can't be empty if prefix is set"}
	}

	if cfg.versions == nil {
		cfg.versions = make(map[int]struct{}, 1)
	}
	cfg.versions[cfg.version] = struct{}{}

	cfg.initialized = true

	return &cfg, nil
}

func NewConfig(options ...Option) (*Config, error) {
	return new(Config).Clone(options...)
}
