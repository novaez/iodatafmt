package iodatafmt

import (
	// Base packages.
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	// Third party packages.
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

// DataFmt represents which data serialization is used YAML, JSON or TOML.
type DataFmt int

// Constants for data format.
const (
	YAML DataFmt = iota
	TOML
	JSON
)

// Unmarshal YAML/JSON/TOML serialized data.
func Unmarshal(b []byte, f DataFmt) (map[string]interface{}, error) {
	d := make(map[string]interface{})

	switch f {
	case YAML:
		if err := yaml.Unmarshal(b, &d); err != nil {
			return nil, err
		}
	case TOML:
		if err := toml.Unmarshal(b, &d); err != nil {
			return nil, err
		}
	case JSON:
		if err := json.Unmarshal(b, &d); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported data format")
	}

	return d, nil
}

// Marshal YAML/JSON/TOML serialized data.
func Marshal(d map[string]interface{}, f DataFmt) ([]byte, error) {
	switch f {
	case "YAML":
		if b, err := yaml.Marshal(&d); err != nil {
			return nil, err
		}
		return b, nil
	case "TOML":
		b := new(bytes.Buffer)
		if err := toml.NewEncoder(s).Encode(&b); err != nil {
			return nil, err
		}
		return b.Bytes(), nil
	case "JSON":
		if b, err := json.MarshalIndent(&b, "", "    "); err != nil {
			return nil, err
		}
		return b, nil
	default:
		return nil, errors.New("unsupported data format")
	}
}

// Format returns DataFmt constant based on a string.
func Format(s string) (DataFmt, error) {
	switch strings.ToUpper(s) {
	case "YAML":
		return YAML, nil
	case "TOML":
		return TOML, nil
	case "JSON":
		return JSON, nil
	default:
		return nil, errors.New("unsupported data format")
	}
}

// FileFormat returns DataFmt constant based on file extension.
func FileFormat(fn string) (DataFmt, error) {
	switch filepath.Ext(fn) {
	case ".yaml":
		return YAML, nil
	case ".yml":
		return YAML, nil
	case ".json":
		return JSON, nil
	case ".toml":
		return TOML, nil
	case ".tml":
		return TOML, nil
	default:
		return nil, errors.New("unsupported data format")
	}
}

// Load a file with serialized data.
func Load(fn string, f DataFmt) (map[string]interface{}, error) {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return nil, errors.New("file doesn't exist")
	}

	if b, err := ioutil.ReadFile(fn); err != nil {
		return nil, err
	}

	if d, err := Unmarshal(b, f); err != nil {
		return nil, err
	}

	return d, nil
}

// Write a file with serialized data.
func Write(fn string, f DataFmt, d []byte) error {
	if b, err := Marshal(d, f); err != nil {
		return nil, err
	}

	if w, err := os.Create(fn); err != nil {
		return err
	}

	if _, err = w.Write(b); err != nil {
		return err
	}

	w.Close()
	return nil
}