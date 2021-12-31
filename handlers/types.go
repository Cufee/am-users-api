package handlers

type ResponseJSON struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}
