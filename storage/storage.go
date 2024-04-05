package storage

import (
	"github.com/Mlstermass/LRUcache/api/controller/httpentity"
)

type DocumentActions interface {
	NewsExists(newsArticleID string) (bool, error)
	AddNews(newsItem httpentity.NewsItem) error
	GetNews() ([]httpentity.NewsItem, error)
	GetNewsByID(newsItemID string) (httpentity.NewsItem, error)
}
