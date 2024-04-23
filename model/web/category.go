package web

type RequestParam struct {
	ID   int    `uri:"id"`
}

type RequestBody struct {
	Name string `json:"name" param:"name" validate:"required"`
}
