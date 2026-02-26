package encoding

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

// LoadAndUnmarshalYAML loads data from the specified path and decodes it into
// the specified structure.
func LoadAndUnmarshalYAML(path string, value any) error {
	return LoadAndUnmarshal(path, func(data []byte) error {
		decoder := yaml.NewDecoder(bytes.NewReader(data))
		decoder.KnownFields(true)
		return decoder.Decode(value)
	})
}
