package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Mlstermass/LRUcache/pkg/cache/lru"
	"github.com/Mlstermass/LRUcache/pkg/env"
	"github.com/Mlstermass/LRUcache/storage"
)

const (
	NewsItemIDStr = "newsItemId"
)

type App struct {
	config  env.Config
	storage storage.DocumentActions
	cache   *lru.Cache
}

func NewApp(
	config env.Config,
	storage storage.DocumentActions,
	cache *lru.Cache,
) App {
	return App{
		config:  config,
		storage: storage,
		cache:   cache,
	}
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// @Summary Get all news
// @Description Fetch all news items from the database
// @Tags news
// @Produce json
// @Success 200 {array} httpentity.NewsItem
// @Router /news [get]
func (a *App) GetNews(w http.ResponseWriter, r *http.Request) {
	// Check if the news items are already in the cache
	newsItems, err := a.cache.Get("news")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newsItems != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newsItems)
		return
	}

	// Fetch news items from the storage
	newsItems, err = a.storage.GetNews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store news items in the cache
	a.cache.Put("news", newsItems)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newsItems)
}

// @Summary Get news by ID
// @Description Fetch a single news item from the database by ID
// @Tags news
// @Produce json
// @Param newsItemId path string true "News Item ID"
// @Success 200 {object} httpentity.NewsItem
// @Router /news/{newsItemId} [get]
func (a *App) GetNewsByID(w http.ResponseWriter, r *http.Request) {
	newsItemId := chi.URLParam(r, NewsItemIDStr)

	// Check if the news item is already in the cache
	newsItem, err := a.cache.Get(newsItemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newsItem != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newsItem)
		return
	}

	// Fetch the news item from the storage
	newsItem, err = a.storage.GetNewsByID(newsItemId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Store the news item in the cache
	a.cache.Put(newsItemId, newsItem)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newsItem)
}
