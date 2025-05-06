package compression

import (
	"bytes"
	"io"

	"github.com/pierrec/lz4/v4"
)

// Compress compresses the input data using LZ4 and returns the compressed data.
func Compress(data []byte) ([]byte, error) {
	var compressedData bytes.Buffer
	writer := lz4.NewWriter(&compressedData)

	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return compressedData.Bytes(), nil
}

// Decompress decompresses the input LZ4 compressed data and returns the original data.
func Decompress(compressedData []byte) ([]byte, error) {
	reader := lz4.NewReader(bytes.NewReader(compressedData))
	var decompressedData bytes.Buffer

	_, err := io.Copy(&decompressedData, reader)
	if err != nil {
		return nil, err
	}

	return decompressedData.Bytes(), nil
}
