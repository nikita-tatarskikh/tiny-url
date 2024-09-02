package url_hash

import (
	"fmt"
	"hash/crc32"
)

type URLHashImpl struct {
}

func NewURLHash() URLHash {
	return &URLHashImpl{}
}

func (h URLHashImpl) HashURL(s string) string {
	return fmt.Sprintf("%x", crc32.ChecksumIEEE([]byte(s)))
}
