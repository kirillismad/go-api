package gin

import (
	"go-api/internal/apps/entities"
	"go-api/pkg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateEntityHandler(useCase entities.CreateEntityUseCase) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// parse data
		var input CreateEntityInput
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.Error(err)
			return
		}

		/* Business Logic */
		data := entities.CreateEntityData{
			UUIDField:     uuid.MustParse(input.UUIDField),
			IntField:      *input.IntField,
			FloatField:    *input.FloatField,
			DatetimeField: pkg.Must(time.Parse(time.RFC3339, input.DatetimeField)),
			StringField:   input.StringField,
			BoolField:     *input.BoolField,
		}
		result, err := useCase.CreateEntity(ctx, data)
		/* Business Logic ends */

		// handle error
		if err != nil {
			ctx.Error(err)
			return
		}

		// set output
		output := CreateEntityOutput{ID: result.ID}
		ctx.JSON(http.StatusCreated, output)
	}
}
