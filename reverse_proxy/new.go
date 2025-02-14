package reverseproxy

import (
	"net/http/httputil"
	"net/url"
)

func newReverseProxy(target string) (*httputil.ReverseProxy, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return proxy, nil
}
