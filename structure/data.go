package structure

type Configuration struct {
	Version  string   `yaml:"version"`
	Schemas  []Schema `yaml:"schemas"`
	Provider Provider `yaml:"provider"`
}

type Schema struct {
	Name    string   `yaml:"name"`
	Columns []Column `yaml:"columns"`
}

type Column struct {
	Name       string `yaml:"name"`
	Type       string `yaml:"type"`
	MinLength  int    `yaml:"min_length,omitempty"`
	MaxLength  int    `yaml:"max_length,omitempty"`
	Faker      string `yaml:"faker,omitempty"`
	ForeignKey string `yaml:"foreign_key,omitempty"`
	Optional   bool   `yaml:"optional"`
}

type Provider interface {
	isProvider()
}

type RestConfig struct {
	Name           string   `yaml:"name"`
	URL            string   `yaml:"url"`
	EntityEndpoint Endpoint `yaml:"entity_endpoint"`
	Method         string   `yaml:"method"`
	Headers        []Header `yaml:"headers"`
}

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbType   string `yaml:"db_type"`
}

type SqLiteConfig struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type Endpoint struct {
	SchemaName string `yaml:"schema_name"`
	Path       string `yaml:"path"`
}

type Header struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// Méthode pour satisfaire l'interface Provider
func (RestConfig) isProvider()     {}
func (DatabaseConfig) isProvider() {}
func (SqLiteConfig) isProvider()   {}
