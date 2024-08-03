package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInput_FromFile(t *testing.T) {
	// create a temporary file with test content
	tmpFile, err := os.CreateTemp("", "testfile-*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	testContent := []byte("test: content")
	n, err := tmpFile.Write(testContent)
	assert.NoError(t, err)
	assert.Greater(t, n, 0)

	tmpFile.Close()

	// call with temp file path
	content, err := ReadInput(tmpFile.Name())
	assert.NoError(t, err)

	assert.Equal(t, string(testContent), string(content))
}

func TestReadInput_FromStdin(t *testing.T) {
	// prepare stdin with test content
	oldStdin := os.Stdin
	r, w, err := os.Pipe()
	assert.NoError(t, err)
	os.Stdin = r

	testContent := "test: content"

	done := make(chan struct{})
	// stdin writer go routine
	go func() {
		defer w.Close()
		_, err := w.Write([]byte(testContent))
		assert.NoError(t, err)
		close(done)
	}()

	// call with no file path
	content, err := ReadInput("")
	assert.NoError(t, err)

	assert.Equal(t, string(testContent), string(content))

	// wait for stdin writer goroutine to complete
	<-done

	// restore original stdin
	os.Stdin = oldStdin
}
