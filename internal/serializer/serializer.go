// Package serializer provides functionality for serializing protocol buffer
// messages into different formats.
package serializer

import (
	"bytes"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/austinthao5/golang_proto_test/internal/parser"
)

// Serializer is an interface that wraps a method to serialize a proto.Message
// to a []byte.
type Serializer interface {
	Serialize(proto.Message) ([]byte, error)
}

type yaml struct{}

// Yaml is a Serializer that serializes a proto.Message in YAML format.
var Yaml Serializer = new(yaml)

func (yaml) Serialize(m proto.Message) ([]byte, error) {
	return parser.Marshal(m)
}

type settingsJs struct{}

// SettingsJs is a Serializer that serializes a proto.Message by outputting a
// JavaScript file to be consumed by Deck.
var SettingsJs Serializer = new(settingsJs)

func (settingsJs) Serialize(m proto.Message) ([]byte, error) {
	json, err := protojson.Marshal(m)
	if err != nil {
		return nil, err
	}

	const jsStart = "window.spinnakerSettings = JSON.parse('"
	const jsEnd = "');"
	js := append([]byte(jsStart), json...)
	js = append(js, []byte(jsEnd)...)

	return js, nil
}

type envFile struct{}

// EnvFile is a Serializer that serializes a proto.message as an environment
// variable file.
// For each top-level key in the message, a line key=value is appended to the
// output, where key is the field's JSON name and value is the string string
// representation of the field's value.
var EnvFile Serializer = new(envFile)

func (envFile) Serialize(m proto.Message) ([]byte, error) {
	var buf bytes.Buffer
	var err error
	r := m.ProtoReflect()
	r.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if _, err = fmt.Fprintf(&buf, "%s=%s\n", fd.JSONName(), v); err != nil {
			return false
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
