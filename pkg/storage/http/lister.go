package http

import (
	"bytes"
	"context"
	"strings"

	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/observ"
	"golang.org/x/net/html"
)

func (s *storageImpl) List(ctx context.Context, module string) ([]string, error) {
	const op errors.Op = "http.List"
	_, span := observ.StartSpan(ctx, op.String())
	defer span.End()
	loc := s.moduleLocation(module)
	fileInfos, err := readFile(loc)
	if err != nil {
		return nil, errors.E(op, errors.M(module), err, errors.KindUnexpected)
	}

	ret, err := fileListFromHTML(fileInfos)
	if err != nil {
		return nil, errors.E(op, errors.M(module), err, errors.KindUnexpected)
	}
	return ret, nil
}

func fileListFromHTML(content []byte) (ret []string, err error) {
	tkn := html.NewTokenizer(bytes.NewReader(content))

	var isA bool
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tkn.Token()
			isA = t.Data == "a"
		case tt == html.TextToken:
			t := tkn.Token()
			if isA && !strings.HasSuffix(t.Data, "/") {
				ret = append(ret, t.Data)
			}
			isA = false
		}
	}
}
