package substr_handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"tsarka/internal/api"
	substr_handlers "tsarka/internal/api/substr"
	"tsarka/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestFind(t *testing.T) {
	errorResponse := map[string]interface{}{"error": "input is not correct"}
	testCases := []struct {
		name             string
		substr           string
		expectedStatus   int
		expectedResponse map[string]interface{}
	}{
		{
			name:             "OK",
			substr:           "abcdefg",
			expectedStatus:   http.StatusOK,
			expectedResponse: responseGen("abcdefg", "abcdefg"),
		},
		{
			name:             "OK",
			substr:           "aaaaa",
			expectedStatus:   http.StatusOK,
			expectedResponse: responseGen("aaaaa", "a"),
		},
		{
			name:             "OK",
			substr:           "abcdeefghij",
			expectedStatus:   http.StatusOK,
			expectedResponse: responseGen("abcdeefghij", "efghij"),
		},
		{
			name:             "OK",
			substr:           "abcdefghijklmnopqrstuvwxyz",
			expectedStatus:   http.StatusOK,
			expectedResponse: responseGen("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"),
		},
		{
			name:             "OK",
			substr:           "abcdedefghijklmnop",
			expectedStatus:   http.StatusOK,
			expectedResponse: responseGen("abcdedefghijklmnop", "defghijklmnop"),
		},
		{
			name:             "BadRequest",
			substr:           "",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: errorResponse,
		},
		{
			name:             "BadRequest",
			substr:           "~ferf",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: errorResponse,
		},
		{
			name:             "BadRequest",
			substr:           "укйкйу",
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: errorResponse,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ss := service.NewSubstrService()
			sh := substr_handlers.NewSubstHandler(ss)
			gin.SetMode(gin.TestMode)
			routes := gin.Default()
			app := api.Application{Substr: sh}
			app.RegisterRoutes(routes)

			recorder := httptest.NewRecorder()

			req := map[string]string{"substr": tc.substr}

			request, err := http.NewRequest(http.MethodGet, "/rest/substr/find", mapToBody(t, req))

			require.NoError(t, err)
			routes.ServeHTTP(recorder, request)
			require.Equal(t, tc.expectedStatus, recorder.Code)
			requireBodyMatch(t, recorder.Body, tc.expectedResponse)
		})
	}
}

func responseGen(substr, res string) map[string]interface{} {
	keyStr := fmt.Sprintf("substr of %s", substr)
	return map[string]interface{}{keyStr: res}
}

func requireBodyMatch(t *testing.T, body *bytes.Buffer, want map[string]interface{}) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotRes map[string]interface{}
	err = json.Unmarshal(data, &gotRes)
	require.NoError(t, err)
	require.Equal(t, want, gotRes)
}

func mapToBody(t *testing.T, data map[string]string) io.Reader {
	jsonBytes, err := json.Marshal(data)
	require.NoError(t, err)
	bodyReader := bytes.NewReader(jsonBytes)
	return bodyReader
}
