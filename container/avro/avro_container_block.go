// Code generated by github.com/ParspooyeshFanavar/gogen-avro/v5. DO NOT EDIT.
/*
 * SOURCES:
 *     block.avsc
 *     header.avsc
 */

package avro

import (
	"io"
)

type AvroContainerBlock struct {
	NumRecords  int64
	RecordBytes []byte
	Sync        Sync
}

func DeserializeAvroContainerBlock(r io.Reader) (*AvroContainerBlock, error) {
	return readAvroContainerBlock(r)
}

func (r *AvroContainerBlock) Schema() string {
	return "{\"fields\":[{\"name\":\"numRecords\",\"type\":\"long\"},{\"name\":\"recordBytes\",\"type\":\"bytes\"},{\"name\":\"sync\",\"type\":{\"name\":\"sync\",\"size\":16,\"type\":\"fixed\"}}],\"name\":\"AvroContainerBlock\",\"type\":\"record\"}"
}

func (r *AvroContainerBlock) Serialize(w io.Writer) error {
	return writeAvroContainerBlock(r, w)
}
