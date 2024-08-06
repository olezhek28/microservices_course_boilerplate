package app

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// CORSMux обертка над runtime.ServeMux для отключения CORS
type CORSMux struct {
	mux *runtime.ServeMux
}

// NewCORSMux новый экземпляр
func NewCORSMux(mux *runtime.ServeMux) *CORSMux {
	return &CORSMux{
		mux: mux,
	}
}

// ServeHTTP отключает CORS
func (m *CORSMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	m.mux.ServeHTTP(w, r)
}
