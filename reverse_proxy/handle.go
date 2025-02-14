package reverseproxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"sync"

	"github.com/rafa-souza-dev/dog-balancer/utils"
)

func HandleRedirectRequest(
	proxy *httputil.ReverseProxy, 
	currentIndex *int,
	mu *sync.Mutex,
	slice []string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Forwarding request: %s %s", r.Method, r.URL.Path)

		mu.Lock()
		utils.IncrementSliceIndex(currentIndex, slice)
		mu.Unlock()
	
		proxy.ServeHTTP(w, r)
	}
}
