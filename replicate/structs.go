package replicate

type Client struct {
	Authorization string
	API           string
}

type Request struct {
	Version string `json:"version"`
	Input   any    `json:"input"`
}
