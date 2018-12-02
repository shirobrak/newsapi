package services

import (
	"encoding/json"

	"github.com/shirobrak/newsapi/entities/responses"
	"github.com/shirobrak/newsapi/entities/responses/contents"
)

// TopicsAPIServiceInterface is the interface for the adapter to connect.
type TopicsAPIServiceInterface interface {
	SearchArticles(genre string) ([]contents.Article, error)
}

// TopicsAPIService is a Service to provide TopicsAPI.
type TopicsAPIService struct {
	ArticleGetter TopicsAPIServiceInterface
}

// NewTopicsAPIService returns the instance of TopicsAPIService.
func NewTopicsAPIService(articleGetter TopicsAPIServiceInterface) *TopicsAPIService {
	return &TopicsAPIService{ArticleGetter: articleGetter}
}

// Run is an use case of creating the TopicsApiResponse
func (ts *TopicsAPIService) Run(genre string) ([]byte, error) {
	// search articles related to "genre".
	respArticles, err := ts.ArticleGetter.SearchArticles(genre)
	if err != nil {
		return nil, err
	}

	// create response.
	apiResponse := responses.NewTopicsAPIResponse()
	apiResponse.SetID("TOPICS")
	apiResponse.SetAPIVersion(1)
	apiResponse.SetLang("ja")
	apiResponse.SetStatus("ok")
	for _, article := range respArticles {
		apiResponse.AddToRows(article)
	}
	response, err := json.Marshal(apiResponse)
	if err != nil {
		return nil, err
	}
	return response, nil
}
