package utils

import (
	"bytes"
	"context"
	"testing"

	"github.com/lyft/flytestdlib/storage"
	"github.com/stretchr/testify/assert"
)

func TestFailingRawStore(t *testing.T) {
	ctx := context.TODO()
	f := FailingRawStore{}
	_, err := f.Head(ctx, "")
	assert.Error(t, err)

	c := f.GetBaseContainerFQN(ctx)
	assert.Equal(t, storage.DataReference(""), c)

	_, err = f.ReadRaw(ctx, "")
	assert.Error(t, err)

	err = f.WriteRaw(ctx, "", 0, storage.Options{}, bytes.NewReader(nil))
	assert.Error(t, err)
}
