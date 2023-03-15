package http

import (
	errs "errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gomods/athens/pkg/errors"
	"github.com/gomods/athens/pkg/storage"
)

type storageImpl struct {
	baseUrl *url.URL
}

func (s *storageImpl) moduleLocation(module string) string {
	return s.baseUrl.JoinPath(module).String()
}

// NewStorage returns a new ListerSaver implemention that stores
// everything under http server.
func NewStorage(baseUrl string) (storage.Backend, error) {
	const op errors.Op = "http.NewStorage"
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, errors.E(op, fmt.Errorf("could not parse url string"))
	}

	_, err = readFile(baseUrl)
	if err != nil {
		return nil, errors.E(op, fmt.Errorf("url unreachable"))
	}
	return &storageImpl{baseUrl: url}, nil
}

func readFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errs.New("request http site failed")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return body, err
}
