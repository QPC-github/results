package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/viper"
)

func TestParseFileConfig(t *testing.T) {
	viper.SetConfigFile("./testdata/config.yaml")

	testConfig(t, &Config{
		Address: "a",
		Token:   "b",
		SSL: SSLConfig{
			RootsFilePath:      "c",
			ServerNameOverride: "d",
		},
		ServiceAccount: &ServiceAccount{
			Namespace: "e",
			Name:      "f",
		},
	})
}

func TestEnvVarConfig(t *testing.T) {
	viper.SetConfigFile("./testdata/empty.yaml")

	t.Setenv(EnvSSLRootFilePath, "a")
	t.Setenv(EnvSSLServerNameOverride, "b")

	testConfig(t, &Config{
		SSL: SSLConfig{
			RootsFilePath:      "a",
			ServerNameOverride: "b",
		},
	})
}
func TestFlagConfig(t *testing.T) {
	viper.SetConfigFile("./testdata/config.yaml")

	viper.Set("addr", "1")
	viper.Set("token", "2")

	testConfig(t, &Config{
		Address: "1",
		Token:   "2",
		SSL: SSLConfig{
			RootsFilePath:      "c",
			ServerNameOverride: "d",
		},
		ServiceAccount: &ServiceAccount{
			Namespace: "e",
			Name:      "f",
		},
	})
}

func testConfig(t *testing.T, want *Config) {
	cfg, err := GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(cfg, want); diff != "" {
		t.Error(diff)
	}
}
