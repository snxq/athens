package http

import (
	"bytes"
	"context"
	"io"

	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/observ"
	"github.com/gomods/athens/pkg/storage"
)

func (s *storageImpl) Info(ctx context.Context, module,
	version string) ([]byte, error) {

	const op errors.Op = "http.Info"
	_, span := observ.StartSpan(ctx, op.String())
	defer span.End()
	// versionedPath := s.versionLocation(module, version)
	infoPath := s.baseUrl.JoinPath(module, version, version+".info").String()
	info, err := readFile(infoPath)
	if err != nil {
		return nil, errors.E(op, errors.M(module), errors.V(version), errors.KindNotFound)
	}
	return info, nil
}

func (s *storageImpl) GoMod(ctx context.Context, module,
	version string) ([]byte, error) {

	const op errors.Op = "http.GoMod"
	_, span := observ.StartSpan(ctx, op.String())
	defer span.End()
	modPath := s.baseUrl.JoinPath(module, version, "go.mod").String()
	mod, err := readFile(modPath)
	if err != nil {
		return nil, errors.E(op, errors.M(module), errors.V(version), errors.KindNotFound)
	}
	return mod, nil
}

func (s *storageImpl) Zip(ctx context.Context, module,
	version string) (storage.SizeReadCloser, error) {

	const op errors.Op = "http.Zip"
	_, span := observ.StartSpan(ctx, op.String())
	defer span.End()
	srcPath := s.baseUrl.JoinPath(module, version, "source.zip").String()
	src, err := readFile(srcPath)
	if err != nil {
		return nil, errors.E(op, errors.M(module), errors.V(version), errors.KindNotFound)
	}
	buf := bytes.NewBuffer(src)
	return storage.NewSizer(io.NopCloser(buf), int64(buf.Len())), nil
}
