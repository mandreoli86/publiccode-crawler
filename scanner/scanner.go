package scanner

import (
	"errors"
	"net/url"

	"github.com/italia/developers-italia-backend/common"
)

var (
	ErrPubliccodeNotFound = errors.New("publiccode.yml not found")
)

type Scanner interface {
	ScanRepo(url url.URL, publisher common.Publisher, repositories chan common.Repository) error
	ScanGroupOfRepos(url url.URL, publisher common.Publisher, repositories chan common.Repository) error
}