package tagshelf

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileUploadPayload(t *testing.T) {
	expected := `{"url": "http://example.com", "metadata": {"hello": "world"}}`
	upload := NewFileUpload()
	upload.Add("http://example.com")
	upload.MetaData = FileMetadata{
		"hello": "world",
	}

	b, err := json.Marshal(upload)
	if err != nil {
		fmt.Println("error:", err)
	}

	require.JSONEq(t, expected, string(b))
}

func TestFileUploadPayloadNoMeta(t *testing.T) {
	expected := `{"url": "http://example.com"}`
	upload := NewFileUpload()
	upload.Add("http://example.com")

	b, err := json.Marshal(upload)
	if err != nil {
		fmt.Println("error:", err)
	}

	require.JSONEq(t, expected, string(b))
}
