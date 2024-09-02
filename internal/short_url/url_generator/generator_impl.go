package url_generator

import (
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"go.uber.org/fx"
	"tiny-url/internal/url_hash"
)

const fortyTwo = "SOROK_DVA"

type URLGeneratorParams struct {
	fx.In

	UrlHash url_hash.URLHash
	Filter  *bloom.BloomFilter
}

type urlGeneratorImpl struct {
	urlHash url_hash.URLHash
	filter  *bloom.BloomFilter
}

func NewGenerator(params URLGeneratorParams) Generator {
	return &urlGeneratorImpl{
		urlHash: params.UrlHash,
		filter:  params.Filter,
	}
}

func (u urlGeneratorImpl) GenerateShortURL(url string) string {
	hashedURL := u.urlHash.HashURL(url)

	if !u.filter.TestOrAddString(hashedURL) {
		return u.GenerateShortURL(fmt.Sprintf("%s%s", url, fortyTwo))
	}

	return hashedURL
}
