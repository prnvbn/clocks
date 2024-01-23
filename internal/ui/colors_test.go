package ui_test

import (
	"testing"

	"github.com/prnvbn/clocks/internal/ui"
	"gopkg.in/yaml.v3"
)

const (
	yamlStr = "test: red\n"
)

// test yaml unmarshalling for Color
func TestColor_UnmarshalYAML(t *testing.T) {
	var c struct {
		Test ui.Color `yaml:"test"`
	}

	err := yaml.Unmarshal([]byte(yamlStr), &c)
	if err != nil {
		t.Fatal(err)
	}
	if c.Test != ui.Red {
		t.Fatalf("expected %s, got %s", ui.Red, c)
	}
}

// test yaml marshalling for Color
func TestColor_MarshalYAML(t *testing.T) {
	var c struct {
		Test ui.Color `yaml:"test"`
	}
	c.Test = ui.Red
	yamlBytes, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}
	if string(yamlBytes) != yamlStr {
		t.Fatalf("expected %s, got %s", yamlStr, string(yamlBytes))
	}
}
