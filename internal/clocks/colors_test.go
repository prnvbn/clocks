package clocks_test

import (
	"testing"

	"github.com/prnvbn/clocks/internal/clocks"
	"gopkg.in/yaml.v3"
)

const (
	yamlStr = "test: red\n"
)

// test yaml unmarshalling for Color
func TestColor_UnmarshalYAML(t *testing.T) {
	var c struct {
		Test clocks.Color `yaml:"test"`
	}

	err := yaml.Unmarshal([]byte(yamlStr), &c)
	if err != nil {
		t.Fatal(err)
	}
	if c.Test != clocks.Red {
		t.Fatalf("expected %s, got %s", clocks.Red, c)
	}
}

// test yaml marshalling for Color
func TestColor_MarshalYAML(t *testing.T) {
	var c struct {
		Test clocks.Color `yaml:"test"`
	}
	c.Test = clocks.Red
	yamlBytes, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	if string(yamlBytes) != yamlStr {
		t.Fatalf("expected %s, got %s", yamlStr, string(yamlBytes))
	}
}
