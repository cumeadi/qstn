package app

type Error struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

var codes = map[int]string{
	400: "Bad Request",
	500: "Internal Server Error",
	404: "Not Found",
}
