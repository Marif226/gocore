package transformer

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"reflect"
)

type Format string

const (
	JSON Format = "json"
	XML  Format = "xml"
)

func Transform(data any, format Format) ([]byte, error) {
	switch format {
	case JSON:
		return json.Marshal(data)
	case XML:
		return xml.Marshal(data)
	default:
		return nil, errors.New("unsupported format: " + string(format))
	}
}

func Parse(data []byte, format Format, target any) error {
	if target == nil {
		return errors.New("target cannot be nil")
	}

	rv := reflect.ValueOf(target)
	if rv.Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
	}

	if rv.IsNil() {
		return errors.New("target pointer cannot be nil")
	}

	switch format {
	case JSON:
		return json.Unmarshal(data, target)
	case XML:
		return xml.Unmarshal(data, target)
	default:
		return errors.New("unsupported format: " + string(format))
	}
}
