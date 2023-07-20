package counter_handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"tsarka/internal/api"
	counter_handlers "tsarka/internal/api/counter"
	"tsarka/internal/mocks"
	"tsarka/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestIncrement(t *testing.T) {
	validIncr := int64(1)
	validExpected := int64(1)
	invalidIncr := "invalid"

	testCases := []struct {
		name          string
		url           string
		buildStubs    func(repo *mocks.MockCounterRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			url:  fmt.Sprintf("/rest/counter/add/%d", validIncr),
			buildStubs: func(repo *mocks.MockCounterRepository) {
				repo.EXPECT().Increment(gomock.Any(), gomock.Eq(validIncr)).Times(1).Return(int64(1), nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				want := map[string]interface{}{"success": fmt.Sprintf("counter incremented by %d and now it is %d", validIncr, validExpected)}
				requireBodyMatch(t, recorder.Body, want)
			},
		},
		{
			name: "BadRequest",
			url:  fmt.Sprintf("/rest/counter/add/%s", invalidIncr),
			buildStubs: func(repo *mocks.MockCounterRepository) {
				repo.EXPECT().Increment(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				want := map[string]interface{}{"error": "increment value is incorrect"}
				requireBodyMatch(t, recorder.Body, want)
			},
		},
		{
			name: "Internal",
			url:  fmt.Sprintf("/rest/counter/add/%d", validIncr),
			buildStubs: func(repo *mocks.MockCounterRepository) {
				repo.EXPECT().Increment(gomock.Any(), gomock.Any()).Times(1).Return(int64(0), redis.ErrClosed)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockCounterRepository(ctrl)
			tc.buildStubs(repo)
			cs := service.NewCounterService(repo)
			gin.SetMode(gin.TestMode)
			routes := gin.Default()
			ch := counter_handlers.NewCounterHandlers(cs)
			app := api.Application{Counter: ch}
			app.RegisterRoutes(routes)

			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodPut, tc.url, nil)
			require.NoError(t, err)

			routes.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestDecrement(t *testing.T) {
	validDecr := int64(1)
	validExpected := int64(-1)
	invalidDecr := "invalid"

	testCases := []struct {
		name          string
		url           string
		buildStubs    func(repo *mocks.MockCounterRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			url:  fmt.Sprintf("/rest/counter/sub/%d", validDecr),
			buildStubs: func(repo *mocks.MockCounterRepository) {
				repo.EXPECT().Decrement(gomock.Any(), gomock.Eq(validDecr)).Times(1).Return(int64(-1), nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				want := map[string]interface{}{"success": fmt.Sprintf("counter decremented by %d and now it is %d", validDecr, validExpected)}
				requireBodyMatch(t, recorder.Body, want)
			},
		},
		{
			name: "BadRequest",
			url:  fmt.Sprintf("/rest/counter/sub/%s", invalidDecr),
			buildStubs: func(repo *mocks.MockCounterRepository) {
				repo.EXPECT().Decrement(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				want := map[string]interface{}{"error": "decrement value is incorrect"}
				requireBodyMatch(t, recorder.Body, want)
			},
		},
		{
			name: "Internal",
			url:  fmt.Sprintf("/rest/counter/sub/%d", validDecr),
			buildStubs: func(repo *mocks.MockCounterRepository) {
				repo.EXPECT().Decrement(gomock.Any(), gomock.Any()).Times(1).Return(int64(0), redis.ErrClosed)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockCounterRepository(ctrl)
			tc.buildStubs(repo)
			cs := service.NewCounterService(repo)
			gin.SetMode(gin.TestMode)
			routes := gin.Default()
			ch := counter_handlers.NewCounterHandlers(cs)
			app := api.Application{Counter: ch}
			app.RegisterRoutes(routes)

			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodPut, tc.url, nil)
			require.NoError(t, err)

			routes.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestGetValue(t *testing.T) {
	retValue := int64(1)
	testCases := []struct {
		name          string
		url           string
		buildStubs    func(repo *mocks.MockCounterRepository)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			url:  "/rest/counter/val",
			buildStubs: func(repo *mocks.MockCounterRepository) {
				repo.EXPECT().GetValue(gomock.Any()).Times(1).Return(retValue, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				want := map[string]interface{}{"value": float64(retValue)}
				requireBodyMatch(t, recorder.Body, want)
			},
		},
		{
			name: "Internal",
			url:  "/rest/counter/val",
			buildStubs: func(repo *mocks.MockCounterRepository) {
				repo.EXPECT().GetValue(gomock.Any()).Times(1).Return(int64(0), redis.ErrClosed)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockCounterRepository(ctrl)
			cs := service.NewCounterService(repo)
			ch := counter_handlers.NewCounterHandlers(cs)
			app := api.Application{Counter: ch}

			tc.buildStubs(repo)

			gin.SetMode(gin.TestMode)
			routes := gin.Default()
			app.RegisterRoutes(routes)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, tc.url, nil)
			require.NoError(t, err)

			routes.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

func requireBodyMatch(t *testing.T, body *bytes.Buffer, want map[string]interface{}) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotRes map[string]interface{}
	err = json.Unmarshal(data, &gotRes)
	require.NoError(t, err)
	require.Equal(t, want, gotRes)
}
