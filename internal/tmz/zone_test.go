package tmz_test

import (
	"testing"

	"github.com/prnvbn/clocks/internal/tmz"
	"gopkg.in/yaml.v3"
)

const (
	yamlStr = "test: Asia/Kolkata\n"
)

// test yaml unmarshalling for Color
func Test_ZoneUnmarshalYAML(t *testing.T) {
	var c struct {
		Test tmz.Zone `yaml:"test"`
	}

	err := yaml.Unmarshal([]byte(yamlStr), &c)
	if err != nil {
		t.Fatal(err)
	}

	expected := tmz.Zone("Asia/Kolkata")
	if c.Test != expected {
		t.Fatalf("expected %s, got %s", expected, c)
	}
}

// test yaml marshalling for Color
func Test_ZoneMarshalYAML(t *testing.T) {
	var c struct {
		Test tmz.Zone `yaml:"test"`
	}
	c.Test = tmz.Zone("Asia/Kolkata")
	yamlBytes, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	if string(yamlBytes) != yamlStr {
		t.Fatalf("expected %s, got %s", yamlStr, string(yamlBytes))
	}
}
