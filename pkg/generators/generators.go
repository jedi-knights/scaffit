package generators

// ComponentGenerator is the interface for all generators
type ComponentGenerator interface {
	Location() string
	Generate() error
}
