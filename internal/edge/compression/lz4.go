package compression

import (
	"bytes"

	"github.com/pierrec/lz4/v4"
)

func Compress(
	payload []byte,
) ([]byte, error) {
	var buf bytes.Buffer

	writer := lz4.NewWriter(&buf)

	if _, err := writer.Write(
		payload,
	); err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
