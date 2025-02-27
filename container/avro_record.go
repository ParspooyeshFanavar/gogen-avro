package container

import (
	"io"
)

/*
AvroRecord is an interface fulfilled by the structs generated by gogen-avro for Avro records.
*/
type AvroRecord interface {
	Serialize(io.Writer) error
	Schema() string
}
