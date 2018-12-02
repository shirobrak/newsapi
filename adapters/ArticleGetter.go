package adapters

import (
	"encoding/json"
	"time"

	"github.com/shirobrak/newsapi/entities/responses/contents"
)

// NewsAPIResponse is a data structure of response of Repositoy.
type NewsAPIResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// Article is a data structure contained in NewsAPIResponses.
type Article struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

// ArticleGetterInterface is the interface for the repository to connect.
type ArticleGetterInterface interface {
	CallTopHeadlinesAPI(category string) ([]byte, error)
}

// ArticleGetter is Adapter to convert Response of Repository into the Article Data.
type ArticleGetter struct {
	Repository ArticleGetterInterface
}

// NewArticleGetter returns an instance of ArticleGetter.
func NewArticleGetter(repository ArticleGetterInterface) *ArticleGetter {
	return &ArticleGetter{Repository: repository}
}

// SearchArticles returns Articles searched for Repository.
func (g *ArticleGetter) SearchArticles(genre string) ([]contents.Article, error) {
	var articles []contents.Article
	respTopHeadlinesAPI, err := g.Repository.CallTopHeadlinesAPI(genre)
	if err != nil {
		return nil, err
	}

	var newsAPIResponse NewsAPIResponse
	err = json.Unmarshal(respTopHeadlinesAPI, &newsAPIResponse)
	if err != nil {
		return nil, err
	}

	switch newsAPIResponse.Status {
	case "ok":
		articles = convertArticles(newsAPIResponse.Articles)
		return articles, nil
	case "error":
		// TODO : Throw Error
		return nil, nil
	default:
		// TODO : Throw Error
		return nil, nil
	}
}

func convertArticles(headLineArticles []Article) []contents.Article {
	var articles []contents.Article
	for _, headLineArticle := range headLineArticles {
		var article contents.Article
		article.Title = headLineArticle.Title
		article.Summary = headLineArticle.Description
		article.URL = headLineArticle.URL
		article.ImageURL = headLineArticle.URLToImage
		article.PublishedAt = headLineArticle.PublishedAt
		articles = append(articles, article)
	}
	return articles
}
