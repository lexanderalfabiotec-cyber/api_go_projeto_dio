package config

import (
	"os"
	"path/filepath"
	"testing"
)

// ---------- helpers ----------

// writeYAML grava um arquivo YAML temporário e retorna o caminho e uma função de limpeza.
func writeYAML(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("não foi possível criar config.yml temporário: %v", err)
	}
	return path
}

// ---------- loadYAML ----------

func TestLoadYAML_ValidFile(t *testing.T) {
	path := writeYAML(t, `
server:
  port: "9090"
log:
  level: "debug"
app:
  name: "minha-api"
  version: "2.0.0"
`)

	yc, err := loadYAML(path)
	if err != nil {
		t.Fatalf("esperava nil, obteve erro: %v", err)
	}
	if yc.Server.Port != "9090" {
		t.Errorf("Server.Port: esperava '9090', obteve '%s'", yc.Server.Port)
	}
	if yc.App.Name != "minha-api" {
		t.Errorf("App.Name: esperava 'minha-api', obteve '%s'", yc.App.Name)
	}
	if yc.App.Version != "2.0.0" {
		t.Errorf("App.Version: esperava '2.0.0', obteve '%s'", yc.App.Version)
	}
}

func TestLoadYAML_FileNotFound(t *testing.T) {
	_, err := loadYAML("/caminho/inexistente/config.yml")
	if err == nil {
		t.Error("esperava erro para arquivo inexistente, obteve nil")
	}
}

func TestLoadYAML_InvalidYAML(t *testing.T) {
	path := writeYAML(t, ":: yaml inválido ::")
	_, err := loadYAML(path)
	if err == nil {
		t.Error("esperava erro para YAML inválido, obteve nil")
	}
}

// ---------- firstNonEmpty ----------

func TestFirstNonEmpty(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected string
	}{
		{"primeiro não-vazio", []string{"", "b", "c"}, "b"},
		{"todos vazios", []string{"", "", ""}, ""},
		{"primeiro preenchido", []string{"a", "b"}, "a"},
		{"lista vazia", []string{}, ""},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := firstNonEmpty(tc.values...)
			if got != tc.expected {
				t.Errorf("firstNonEmpty(%v) = '%s'; esperava '%s'", tc.values, got, tc.expected)
			}
		})
	}
}

// ---------- getEnv ----------

func TestGetEnv_WithEnvSet(t *testing.T) {
	t.Setenv("TEST_PORT", "1234")
	got := getEnv("TEST_PORT", "9999")
	if got != "1234" {
		t.Errorf("getEnv: esperava '1234', obteve '%s'", got)
	}
}

func TestGetEnv_WithDefault(t *testing.T) {
	os.Unsetenv("TEST_PORT_MISSING")
	got := getEnv("TEST_PORT_MISSING", "5678")
	if got != "5678" {
		t.Errorf("getEnv: esperava '5678', obteve '%s'", got)
	}
}

// ---------- Load (integração) ----------

// TestLoad_FromYAML verifica que Load usa os valores do YAML.
// Ajustamos a função Load para aceitar um path opcional via variável de ambiente
// CONFIG_PATH para permitir testes sem alterar o diretório de trabalho.
func TestLoad_DefaultValues(t *testing.T) {
	// Remove env vars que possam interferir nos valores padrão
	for _, key := range []string{"SERVER_PORT", "LOG_LEVEL", "APP_NAME", "APP_VERSION", "CONFIG_PATH"} {
		old, hadOld := os.LookupEnv(key)
		os.Unsetenv(key)
		if hadOld {
			t.Cleanup(func() { os.Setenv(key, old) })
		}
	}
	// Aponta para arquivo inexistente → deve usar valores padrão
	os.Setenv("CONFIG_PATH", "/inexistente/config.yml")
	t.Cleanup(func() { os.Unsetenv("CONFIG_PATH") })

	cfg, err := Load()
	// Erro é esperado (arquivo não existe), mas cfg não deve ser nil
	if cfg == nil {
		t.Fatal("Load() não deve retornar cfg nil mesmo com erro")
	}
	_ = err // erro tolerado — fallback é o comportamento correto

	if cfg.ServerPort != "8000" {
		t.Errorf("ServerPort padrão: esperava '8000', obteve '%s'", cfg.ServerPort)
	}
	if cfg.LogLevel != "info" {
		t.Errorf("LogLevel padrão: esperava 'info', obteve '%s'", cfg.LogLevel)
	}
	if cfg.AppName != "contatos-api" {
		t.Errorf("AppName padrão: esperava 'contatos-api', obteve '%s'", cfg.AppName)
	}
	if cfg.AppVersion != "1.0.0" {
		t.Errorf("AppVersion padrão: esperava '1.0.0', obteve '%s'", cfg.AppVersion)
	}
}

func TestLoad_EnvVarsOverride(t *testing.T) {
	t.Setenv("SERVER_PORT", "7777")
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("APP_NAME", "env-api")
	t.Setenv("APP_VERSION", "3.0.0")
	t.Setenv("CONFIG_PATH", "/inexistente/config.yml")

	cfg, _ := Load()
	if cfg.ServerPort != "7777" {
		t.Errorf("ServerPort: esperava '7777', obteve '%s'", cfg.ServerPort)
	}
	if cfg.LogLevel != "debug" {
		t.Errorf("LogLevel: esperava 'debug', obteve '%s'", cfg.LogLevel)
	}
	if cfg.AppName != "env-api" {
		t.Errorf("AppName: esperava 'env-api', obteve '%s'", cfg.AppName)
	}
	if cfg.AppVersion != "3.0.0" {
		t.Errorf("AppVersion: esperava '3.0.0', obteve '%s'", cfg.AppVersion)
	}
}
