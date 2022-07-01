package categories

type CategoryRequestParams struct {
	ID int `json:"id"`
}

type CategoryRequestBody struct {
	Category string `json:"category"`
}
