package http

import "context"

// Delete nothing to do
func (s *storageImpl) Delete(ctx context.Context, module, version string) error {
	return nil
}
