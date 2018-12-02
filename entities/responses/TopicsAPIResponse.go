package responses

import (
	"reflect"

	"github.com/shirobrak/newsapi/entities/responses/contents"
)

// TopicsAPIResponse is a response of TopicsAPI.
type TopicsAPIResponse struct {
	ID         string `json:"id"`
	APIVersion int    `json:"version"`
	Lang       string `json:"lang"`
	Status     string `json:"status"`
	Rows       []Row  `json:"rows"`
}

// Row is a base object of View of Mobile Apps.
type Row struct {
	DataType string      `json:"type"`
	Layout   string      `json:"layout"`
	Content  interface{} `json:"content"`
}

// NewTopicsAPIResponse returns an instance of TopicsAPIResponse.
func NewTopicsAPIResponse() *TopicsAPIResponse {
	return &TopicsAPIResponse{}
}

// SetID sets ID of API.
func (t *TopicsAPIResponse) SetID(id string) {
	t.ID = id
	return
}

// SetStatus sets Status.
func (t *TopicsAPIResponse) SetStatus(status string) {
	t.Status = status
	return
}

// SetLang sets Language contained to Response.
func (t *TopicsAPIResponse) SetLang(lang string) {
	t.Lang = lang
	return
}

// SetAPIVersion sets version of API.
func (t *TopicsAPIResponse) SetAPIVersion(version int) {
	t.APIVersion = version
	return
}

// AddToRows adds Data into Rows
func (t *TopicsAPIResponse) AddToRows(data interface{}) error {
	switch reflect.TypeOf(data) {
	case reflect.TypeOf(contents.Article{}):
		row := Row{DataType: "article", Layout: "article_layout_001", Content: data}
		t.Rows = append(t.Rows, row)
		break
	default:
		// TODO : Error Message
		return nil
	}
	return nil
}
