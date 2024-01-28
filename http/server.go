package http

import (
	"fmt"

	"github.com/yashwanth-reddy909/go-crud/http/handlers"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	articleHander *handlers.ArticleHandler
}

func NewServer(articleHandler *handlers.ArticleHandler) *Server {
	return &Server{
		articleHander: articleHandler,
	}
}

func (s *Server) ListenAndServe() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/articles", func(r chi.Router) {
		r.Get("/", s.articleHander.GetArticles)
		r.Post("/", s.articleHander.CreateArticle)
		r.Put("/{articleId}", s.articleHander.UpdateArticle)
		r.Delete("/{articleId}", s.articleHander.DeleteArticle)
	})
	fmt.Print("Starting server on port 9000")
	err := http.ListenAndServe(":9000", r)
	if err != nil {
		panic("Error while starting server" + err.Error())
	}
}
