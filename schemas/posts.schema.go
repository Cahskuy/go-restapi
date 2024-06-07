package schemas

type Post struct {
	Title  *string       `json:"title" validate:"required,min=3"`
	Body   *string       `json:"body" validate:"required,min=3"`
	Size   *int          `json:"size" validate:"required,min=10"`
	Nested *NestedObject `json:"nested"`
}

type NestedObject struct {
	Field1 string `json:"field1" validate:"required"`
	Field2 int    `json:"field2" validate:"required"`
}
