package poc

type Poc struct {
	URL   string `json:"url"`
	Count int    `json:"count"`
}

type HttpResponse struct {
	StatusCode int
	Body       string
}
