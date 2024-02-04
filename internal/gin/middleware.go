package gin

import (
	"encoding/json"
	"errors"
	"go-api/internal/apps"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func Logging(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		logger.LogAttrs(
			ctx,
			slog.LevelInfo,
			"Incoming request",
			slog.Group(
				"request",
				slog.String("method", ctx.Request.Method),
				slog.String("URL", ctx.Request.URL.String()),
				slog.String("x_request_id", ctx.GetString("request_id")),
			),
			slog.Group(
				"response",
				slog.Int("status", ctx.Writer.Status()),
				slog.Int("size", ctx.Writer.Size()),
			),
			slog.Duration("latency", time.Since(start)),
			slog.String("remote_addr", ctx.Request.RemoteAddr),
		)
	}
}

func HandleErrors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if err := ctx.Errors.Last(); err != nil {
			var (
				jsonError        *json.UnmarshalTypeError
				validationErrors validator.ValidationErrors
			)
			switch {
			case errors.Is(err, apps.ErrUnique):
				ctx.JSON(http.StatusBadRequest, gin.H{"code": CodeUnique})
			case errors.As(err, &jsonError):
			case errors.As(err, &validationErrors):
				ctx.JSON(http.StatusBadRequest, gin.H{"code": CodeInvalidInput})
			default:
				ctx.Status(http.StatusInternalServerError)
			}
		}
	}
}

func HandleRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		ctx.Set("request_id", requestID)

		ctx.Header("X-Request-ID", requestID)

		ctx.Next()
	}
}
