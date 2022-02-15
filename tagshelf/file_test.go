package tagshelf

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tagshelf-api/go/tagshelf/typeutils"
)

func TestFileUploadPayload(t *testing.T) {
	expected := `{"url": "http://example.com", "metadata": {"hello": "world"}}`
	upload := NewFileUpload()
	upload.Add("http://example.com")
	upload.AddMeta(FileMetadata{
		"hello": "world",
	})

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

func TestFileUploadPayloadChannel(t *testing.T) {
	expected := `{"url": "http://example.com","channel": "email"}`
	upload := NewFileUpload()
	upload.Add("http://example.com")
	upload.Channel = "email"

	b, err := json.Marshal(upload)
	if err != nil {
		fmt.Println("error:", err)
	}

	require.JSONEq(t, expected, string(b))
}

func TestFileUploadPayloadPropagateMeta(t *testing.T) {
	expected := []string{
		`{"url": "http://example.com"}`,
		`{"url": "http://example.com","propagate_metadata": false}`,
		`{"url": "http://example.com","propagate_metadata": true}`,
		`{"url": "http://example.com","propagate_metadata": true}`,
	}
	cases := []File{
		{URL: "http://example.com"},
		{URL: "http://example.com", PropagateMeta: typeutils.PointerBool(false)},
		{URL: "http://example.com", PropagateMeta: typeutils.PointerBool(true)},
		{URL: "http://example.com"},
	}
	cases[3].PropagateMetadata(true)

	for i, v := range expected {
		c := cases[i]
		t.Run(fmt.Sprintf("case %d", i), func(tt *testing.T) {
			b, err := json.Marshal(c)
			if err != nil {
				t.Error("error:", err)
			}
			require.JSONEq(tt, v, string(b))
		})
	}
}
