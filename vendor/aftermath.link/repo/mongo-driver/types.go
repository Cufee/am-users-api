package mongodriver

import "aftermath.link/repo/mongo-driver/driver"

type SimpleDriver struct {
	Advanced *driver.Driver
}

type SimpleConfig struct {
	URI      string
	User     string
	Password string
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
	Documents []interface{} `json:"documents"`
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
	Query  map[string]interface{} `json:"query"`
	Update interface{}            `json:"update"`
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

type Count driver.QueryResult
