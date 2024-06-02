package schemas

type Post struct {
	Title string `json:"title" validate:"required,min=3"`
	Body  string `json:"body" validate:"required,min=3"`
}
