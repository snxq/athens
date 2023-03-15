package http

import (
	"context"
	"io"
)

// Save nothing to do.
func (s *storageImpl) Save(ctx context.Context, module,
	version string, mod []byte, zip io.Reader, info []byte) error {

	return nil
}
