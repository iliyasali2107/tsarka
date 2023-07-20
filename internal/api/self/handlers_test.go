package self_handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"tsarka/internal/api"
	self_handlers "tsarka/internal/api/self"
	"tsarka/internal/mocks"
	"tsarka/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestFind(t *testing.T) {
	testCases := []struct {
		name           string
		substr         string
		expectedStatus int
	}{
		{
			name:           "OK",
			substr:         "re",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "OK",
			substr:         "Handler",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "NotFound",
			substr:         "",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "NotFound",
			substr:         "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
			expectedStatus: http.StatusNotFound,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ss := service.NewSelfService()

			sh := self_handlers.NewSelfHandlers(ss)
			gin.SetMode(gin.TestMode)
			routes := gin.Default()

			app := api.Application{Self: sh}
			app.RegisterRoutes(routes)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/rest/self/find/%s", tc.substr), nil)
			require.NoError(t, err)

			recorder := httptest.NewRecorder()

			routes.ServeHTTP(recorder, req)
			require.Equal(t, tc.expectedStatus, recorder.Code)
		})
	}
}

func TestFindInternal(t *testing.T) {
	tc := struct {
		name           string
		substr         string
		expectedStatus int
		buildStubs     func(svc *mocks.MockSelfService)
	}{
		name:           "Internal",
		substr:         "re",
		expectedStatus: http.StatusInternalServerError,
		buildStubs: func(svc *mocks.MockSelfService) {
			svc.EXPECT().Find(gomock.Any(), gomock.Any()).Times(1).Return(nil, fmt.Errorf("error"))
		},
	}
	t.Run("Internal", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ss := mocks.NewMockSelfService(ctrl)
		tc.buildStubs(ss)

		sh := self_handlers.NewSelfHandlers(ss)
		gin.SetMode(gin.TestMode)
		routes := gin.Default()

		app := api.Application{Self: sh}
		app.RegisterRoutes(routes)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/rest/self/find/%s", tc.substr), nil)
		require.NoError(t, err)

		recorder := httptest.NewRecorder()

		routes.ServeHTTP(recorder, req)
		require.Equal(t, tc.expectedStatus, recorder.Code)
	})
}

// func TestFind(t *testing.T) {
// 	f := fuzz.New()
// 	var substr string
// 	for i := 0; i < 100; i++ {
// 		f.Fuzz(&substr)
// 		ss := service.NewSelfService()

// 		sh := self_handlers.NewSelfHandlers(ss)
// 		gin.SetMode(gin.TestMode)
// 		routes := gin.Default()

// 		app := api.Application{Self: sh}
// 		app.RegisterRoutes(routes)

// 		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/rest/self/find/%s", substr), nil)
// 		require.NoError(t, err)

// 		recorder := httptest.NewRecorder()

// 		routes.ServeHTTP(recorder, req)

// 		var response map[string][]string

// 		body := recorder.Body.Bytes()

// 		err = json.Unmarshal(body, &response)
// 		require.NoError(t, err)
// 		require.Equal(t, http.StatusOK, recorder.Code)
// 	}
// }
