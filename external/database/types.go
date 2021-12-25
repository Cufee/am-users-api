package database

type ResponseJSON struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

// Find
type FindRequest struct {
	Query map[string]interface{} `json:"query"`
	Sort  map[string]interface{} `json:"sort"`
	Limit int                    `json:"limit"`
}
type FindResponse struct {
	Documents []map[string]interface{} `json:"documents"`
	Count     `json:"count"`
}

// Insert
type InsertRequest struct {
	Documents []map[string]interface{} `json:"documents"`
}
type InsertResponse struct {
	Count `json:"count"`
}

// Delete
type DeleteRequest struct {
	Query map[string]interface{} `json:"query"`
}
type DeleteResponse struct {
	Count `json:"count"`
}

// Update
type UpdateRequest struct {
	Updates []UpdateItem `json:"updates"`
}
type UpdateItem struct {
	Query  map[string]interface{} `json:"query"`
	Update map[string]interface{} `json:"update"`
	Upsert bool                   `json:"upsert"`
}
type UpdateResponse struct {
	ErrorCount int                   `json:"errorCount"`
	Errors     []UpdateResponseError `json:"errors"`
	Count      `json:"count"`
}
type UpdateResponseError struct {
	Index int    `json:"index"`
	Error string `json:"error"`
}

type Count struct {
	Inserted int `json:"inserted"`
	Updated  int `json:"updated"`
	Deleted  int `json:"deleted"`
	Matched  int `json:"matched"`
}
