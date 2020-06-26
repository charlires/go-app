package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unrolled/render"
)

func TestDemo(t *testing.T) {
	r := render.New()
	tests := []struct {
		name         string
		responseBody []byte
	}{
		{
			name:         "demo test",
			responseBody: []byte("This is fine! 🔥\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c := NewDemo(r)
			req := httptest.NewRequest("GET", "/", nil)
			c.Demo(w, req)
			assert.Equal(t, tt.responseBody, w.Body.Bytes())
		})
	}
}
