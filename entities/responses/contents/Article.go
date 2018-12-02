package contents

import "time"

// Article is the content of the article.
type Article struct {
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	URL         string    `json:"url"`
	ImageURL    string    `json:"img_url"`
	PublishedAt time.Time `json:"published_at"`
}
