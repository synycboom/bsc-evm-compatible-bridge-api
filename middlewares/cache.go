package middlewares

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/cache"
	logger "github.com/synycboom/bsc-evm-compatible-bridge-api/utils/log"
)

const (
	xBncCachedHeaderKey   = "x-bnc-ap-cache"
	xBncCacheTTLHeaderKey = "x-bnc-ap-cache-ttl"
)

var (
	log = logger.Logger
)

type MWCacher struct {
	cache    cache.Store
	cors     *cors.Cors
	duration time.Duration
}

type responderGetter func() middleware.Responder

func NewMWCacher(c *cors.Cors, store cache.Store, duration time.Duration) *MWCacher {
	return &MWCacher{
		cache:    store,
		cors:     c,
		duration: duration,
	}
}

func (c *MWCacher) Serve(req *http.Request, getter responderGetter, prod runtime.Producer) *mwResponder {
	return &mwResponder{c, req, getter, prod}
}

type mwResponder struct {
	*MWCacher
	req       *http.Request
	responder responderGetter
	producer  runtime.Producer
}

// WriteResponse returns a cached result if present, or performs the actual HTTP request if not
func (r *mwResponder) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	key := r.req.RequestURI

	content, exist := r.cache.Get(key)
	if exist {
		w.Header().Set(xBncCachedHeaderKey, "Hit")
		w.Header().Set(xBncCacheTTLHeaderKey,
			strconv.FormatInt(r.cache.TTL(key), 10))
		if err := utils.TryWrite(w, content); err != nil {
			log.Errorf("proxy middleware handler HIT write error: %s", err.Error())
			panic(err) // ends the current request
		}
		return
	}

	if didUpdate := r.cache.WaitOrUpdate(key, r.duration, func() ([]byte, bool) {
		// the updater will return the bytes for caching on success; failed responses will not be cached
		// run the handler, record its response
		var success bool

		rec := httptest.NewRecorder()
		producer := r.producer
		if producer == nil {
			producer = p
		}
		r.responder().WriteResponse(rec, producer)

		// only set cache on a 200-299 status code
		if 199 < rec.Code && rec.Code < 300 && r.duration > 0 {
			success = true
			log.Debugf("Caching page: %s for %s\n", key, r.duration)
		}

		// output response status code and headers
		for k, v := range rec.Result().Header {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)

		w.Header().Set(xBncCachedHeaderKey, "Miss")
		bz := rec.Body.Bytes()
		if bz != nil {
			if err := utils.TryWrite(w, bz); err != nil {
				log.Errorf("proxy middleware handler MISS write error: %s", err.Error())
				return nil, false
			}
		}
		if success {
			return bz, true
		} else {
			return nil, false
		}

	}); !didUpdate {
		// another goroutine ran the update; serve from cache (do this again)
		r.WriteResponse(w, nil)
	}
}

// WriteResponseHandlerFunc is used as a replacement for the standard method of using Cacher when a http.HandlerFunc is wanted
func (r *mwResponder) WriteResponseHandlerFunc(w http.ResponseWriter, req *http.Request) {
	r.cors.HandlerFunc(w, req)
	r.WriteResponse(w, nil)
}
