package gin

type JsonFieldType string

const (
	Nil        JsonFieldType = "Nil"
	Map        JsonFieldType = "Map"
	StringList JsonFieldType = "StringList"
	IntList    JsonFieldType = "IntList"
	Obj        JsonFieldType = "Object"
)

type CreateEntityInput struct {
	UUIDField     string   `json:"uuid_field" binding:"uuid"`
	IntField      *int     `json:"int_field" binding:"required"`
	FloatField    *float64 `json:"float_field" binding:"required"`
	DatetimeField string   `json:"datetime_field" binding:"datetime=2006-01-02T15:04:05Z07:00"`
	StringField   string   `json:"string_field" binding:"required"`
	BoolField     *bool    `json:"bool_field" binding:"required"`
	// JsonFieldType JsonFieldType `json:"json_field_type"`
	// JsonField     any           `json:"json_field"`
}

type CreateEntityOutput struct {
	ID int `json:"id"`
}
