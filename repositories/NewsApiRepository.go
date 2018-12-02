package repositories

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// NewsAPIRepository is a repository to dealing with NewsAPI.
type NewsAPIRepository struct {
	Client *http.Client
}

// NewNewsAPIRepository returns an instance of NewsAPIRepository.
func NewNewsAPIRepository() *NewsAPIRepository {
	client := &http.Client{}
	return &NewsAPIRepository{Client: client}
}

// CallTopHeadlinesAPI calls the TopHeadlinesAPI provided to NewsAPI.
func (r *NewsAPIRepository) CallTopHeadlinesAPI(category string) ([]byte, error) {
	baseURL := "https://newsapi.org/v2/top-headlines"
	values := url.Values{}
	values.Set("country", "us")
	values.Set("category", category)
	query := values.Encode()
	req, err := http.NewRequest("GET", baseURL+"?"+query, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("NEWS_API_TOKEN"))
	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	// TODO Error Message
	return nil, nil
}
