package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhis2-sre/im-user/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("NoErrorsAreANoop", func(t *testing.T) {
		r := gin.Default()
		r.Use(middleware.ErrorHandler())

		r.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "some body")
		})

		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "some body", w.Body.String())
	})

	t.Run("ErrorWithStatusWillOnlySetResponseBody", func(t *testing.T) {
		r := gin.Default()
		r.Use(middleware.ErrorHandler())

		r.GET("/", func(ctx *gin.Context) {
			_ = ctx.AbortWithError(http.StatusUnsupportedMediaType, errors.New("not supported"))
		})

		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)
		assert.Equal(t, "not supported", w.Body.String())
	})

	t.Run("ErrorWithoutStatusWillRespondWithInternalServerError", func(t *testing.T) {
		r := gin.New()
		r.Use(middleware.ErrorHandler())

		r.GET("/", func(ctx *gin.Context) {
			_ = ctx.Error(errors.New("something went wrong"))
		})

		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/", nil)
		require.NoError(t, err)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "something went wrong", w.Body.String())
	})
}
