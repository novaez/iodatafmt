# Go package for input/output of YAML/JSON/TOML files

# GoDoc

```
PACKAGE DOCUMENTATION

package iodatafmt
    import "."


FUNCTIONS

func Format(fn string) error
    Format deduces file format based on extension.

func Load(fn string) (map[string]interface{}, error)
    Load a file with serialized data.

func Marshal(d map[string]interface{}, f DataFmt) ([]byte, error)
    Marshal YAML/JSON/TOML serialized data.

func Unmarshal(b []byte, f DataFmt) (map[string]interface{}, error)
    Unmarshal YAML/JSON/TOML serialized data.

func Write(fn string, d []byte) error
    Write a file with serialized data.

TYPES

type DataFmt int
    DataFmt represents which data serialization is used YAML, JSON or TOML.

const (
    YAML DataFmt = iota
    TOML
    JSON
)
    Constants for data format.
```
