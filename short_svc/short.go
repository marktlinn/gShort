package shortsvc

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const MaxKeyLen = 16

type Link struct {
	Key string
	URL string
}

// validateKey ensures the length of the passed Link's Key
// is less that the MaxKeyLen
func validateKey(k string) error {
	if strings.TrimSpace(k) == "" {
		return errors.New("empty key")
	}
	if len(k) > MaxKeyLen {
		return fmt.Errorf("key too long: max length = %d", MaxKeyLen)
	}
	return nil
}

// validate the passed link, ensuring it has a Host and
// the correct Scheme
func (ln Link) validateLink() error {
	if err := validateKey(ln.Key); err != nil {
		return err
	}
	u, err := url.ParseRequestURI(ln.URL)
	if err != nil {
		return err
	}
	if u.Host == "" {
		return errors.New("no host")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("scheme must be 'HTTP' or 'HTTPS'")
	}
	return nil
}
