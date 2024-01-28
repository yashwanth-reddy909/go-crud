package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yashwanth-reddy909/go-crud/models"

	"github.com/go-chi/chi/v5"
)

type Service interface {
	GetAll() ([]*models.Article, error)
	Create(article *models.Article) (bool, error)
	Update(id string, article *models.ArticleModelWithoutID) (int, error)
	Delete(id string) (int, error)
}

type ArticleHandler struct {
	Service Service
}

func NewArticleHandler(service Service) *ArticleHandler {
	return &ArticleHandler{Service: service}
}

func validateArticleModel(article *models.ArticleModelWithoutID) bool {
	return article.Name != "" && len(article.Authors) != 0
}

func (s *ArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := s.Service.GetAll()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(articles)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func (s *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var articleReqModel models.ArticleModelWithoutID
	err := json.NewDecoder(r.Body).Decode(&articleReqModel)
	if err != nil {
		log.Print("Invalid Article Model")
		w.Write([]byte("Invalid article body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newArticle := models.GetArticleDBModel(&articleReqModel)
	if !validateArticleModel(&articleReqModel) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("required field are missing"))
		return
	} else {
		response, err := s.Service.Create(newArticle)
		w.Header().Set("Content-Type", "application/json")

		jsonResp, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}

func (s *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "articleId")
	var articleReqModel models.ArticleModelWithoutID
	err := json.NewDecoder(r.Body).Decode(&articleReqModel)
	if err != nil {
		log.Print("Invalid Article Model")
		w.Write([]byte("Invalid article body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validateArticleModel(&articleReqModel) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("required field are empty"))
		return
	}
	updatedCount, err := s.Service.Update(articleId, &articleReqModel)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if updatedCount == 0 {
		w.Write([]byte("no articles found with id " + articleId))
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(true)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}

func (s *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	articleId := chi.URLParam(r, "articleId")
	deletedCount, err := s.Service.Delete(articleId)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if deletedCount == 0 {
		w.Write([]byte("no articles found with id " + articleId))
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(true)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}
