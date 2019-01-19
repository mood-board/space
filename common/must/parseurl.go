package must

import "net/url"

func ParseURL(s string) *url.URL {
	var result *url.URL

	Do(func() error {
		var err error
		result, err = url.Parse(s)
		return err
	}())

	return result
}
