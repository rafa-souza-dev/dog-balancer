package reverseproxy

import (
	"log"
	"net/http"
	"sync"

	"github.com/rafa-souza-dev/dog-balancer/utils"
)

func HandleRedirectRequest(
	currentIndex *int,
	mu *sync.Mutex,
	slice []string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Forwarding request: %s %s", r.Method, r.URL.Path)

		mu.Lock()
		proxy, err := newReverseProxy(slice[*currentIndex])
		if err != nil {
			log.Fatalf("Error when create proxy: %v", err)
		}

		utils.IncrementSliceIndex(currentIndex, slice)
		mu.Unlock()
	
		proxy.ServeHTTP(w, r)
	}
}
