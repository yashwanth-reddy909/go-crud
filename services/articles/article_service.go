package articles

import (
	"fmt"

	"github.com/yashwanth-reddy909/go-crud/models"
)

type ArticleRepository interface {
	FindAll() ([]*models.Article, error)
	Create(article *models.Article) (bool, error)
	Update(id string, article *models.ArticleModelWithoutID) (int, error)
	Delete(id string) (int, error)
}

// any business logics goes here
type Service struct {
	repo ArticleRepository
}

func NewArticleService(repo ArticleRepository) *Service{
	return &Service{repo: repo}	
}

func (s *Service) GetAll() ([]*models.Article, error){
	articles, err := s.repo.FindAll()
	if err != nil {
		return []*models.Article{}, fmt.Errorf("failed to get apps: %w", err)
	}
	return articles, nil
}

func (s *Service) Create(article *models.Article) (bool, error){
	isCreated, err := s.repo.Create(article)
	if err != nil || isCreated == false{
		return isCreated, fmt.Errorf("failed to get apps: %w", err)
	}
	return true, nil
}

func (s *Service) Update(id string, article *models.ArticleModelWithoutID) (int, error){
	updatedCount, err := s.repo.Update(id, article)
	if err != nil {
		return updatedCount, fmt.Errorf("failed update: %w", err)
	}
	return updatedCount, nil
}

func (s *Service) Delete(id string) (int, error){
	deletedCount, err := s.repo.Delete(id)
	if err != nil {
		return deletedCount, fmt.Errorf("failed to delete article with %s : %w", id, err)
	}
	return deletedCount, nil
}
