package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// yamlConfig representa a estrutura do arquivo config.yml
type yamlConfig struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"log_level"`
	} `yaml:"log"`
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"app"`
}

// Config armazena todas as configurações da aplicação
type Config struct {
	ServerPort string
	LogLevel   string
	AppName    string
	AppVersion string
}

// Load carrega as configurações a partir do arquivo config.yml.
// Caso o arquivo não exista ou algum campo esteja vazio, utiliza
// variáveis de ambiente como fallback e, por último, valores padrão.
func Load() (*Config, error) {
	path := getEnv("CONFIG_PATH", "config.yml")
	yc, err := loadYAML(path)
	if err != nil {
		// Se não conseguir abrir o arquivo, continua com struct vazia
		yc = &yamlConfig{}
	}

	return &Config{
		ServerPort: firstNonEmpty(yc.Server.Port, getEnv("SERVER_PORT", "8000")),
		LogLevel:   firstNonEmpty(yc.Log.Level, getEnv("LOG_LEVEL", "info")),
		AppName:    firstNonEmpty(yc.App.Name, getEnv("APP_NAME", "contatos-api")),
		AppVersion: firstNonEmpty(yc.App.Version, getEnv("APP_VERSION", "1.0.0")),
	}, nil
}

// loadYAML abre e decodifica o arquivo YAML indicado pelo caminho.
func loadYAML(path string) (*yamlConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var yc yamlConfig
	if err := yaml.NewDecoder(f).Decode(&yc); err != nil {
		return nil, err
	}
	return &yc, nil
}

// firstNonEmpty retorna o primeiro valor não-vazio da lista.
func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// getEnv retorna o valor de uma variável de ambiente ou um valor padrão.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
