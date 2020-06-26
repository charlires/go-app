package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/charlires/go-app/router"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

//go:generate mockgen -destination demo_controller_mock_test.go -package router_test github.com/charlires/go-app/router DemoController
func TestDemo(t *testing.T) {
	testCases := []struct {
		desc        string
		handlerName string
		status      int
	}{
		{
			desc:        "/",
			handlerName: "Demo",
			status:      http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// cfg := &config.Configuration{}

			ctrl := gomock.NewController(t)

			defer ctrl.Finish()

			DemoController := NewMockDemoController(ctrl)

			r := router.Setup(DemoController)

			req, err := http.NewRequest("GET", tc.desc, nil)
			if err != nil {
				t.Fatal(err)
			}

			DemoController.EXPECT().
				Demo(gomock.Any(), gomock.Any()).
				Do(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				})

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			assert.Equal(t, tc.status, rr.Code)
		})
	}
}
