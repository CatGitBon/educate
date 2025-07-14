package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestAuth(t *testing.T) {
	client := NewMockClient()
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request.Header.Set("Authorization", "Bearer TOKEN")
	auth := NewAuthorization(client, func(*gin.Context) bool { return true }, zap.NewNop())

	// todo add auth header to
	auth.Authorize()(ctx)

}
