package service

import (
	"github.com/ohhart/tender-restapi/models"
	"github.com/ohhart/tender-restapi/pkg/repository"
)

// ReviewService handles business logic for reviews
type ReviewService struct {
	reviewRepo repository.ReviewRepository
	bidRepo    repository.BidRepository
}

// NewReviewService creates a new ReviewService
func NewReviewService(reviewRepo repository.ReviewRepository, bidRepo repository.BidRepository) *ReviewService {
	return &ReviewService{
		reviewRepo: reviewRepo,
		bidRepo:    bidRepo,
	}
}

// GetReviewsForTender retrieves reviews for a specific tender
func (s *ReviewService) GetReviewsForTender(tenderID uint, authorUsername string, organizationID uint) ([]models.Review, error) {
	return s.reviewRepo.FindReviewsByTenderAndAuthor(tenderID, authorUsername, organizationID)
}
