package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ohhart/tender-restapi/models"
	"github.com/ohhart/tender-restapi/pkg/repository"
)

// BidService handles business logic for bids
type BidService struct {
	repo       *repository.BidRepository
	reviewRepo *repository.ReviewRepository
}

// NewBidService creates a new BidService
func NewBidService(repo *repository.BidRepository, reviewRepo *repository.ReviewRepository) *BidService {
	return &BidService{repo: repo, reviewRepo: reviewRepo}
}

// CreateBid creates a new bid
func (s *BidService) CreateBid(bid models.Bid) error {
	err := s.repo.CreateBid(bid)
	if err != nil {
		log.Printf("Error creating bid: %v", err)
		return err
	}
	log.Printf("Bid created successfully: %+v", bid)
	return nil
}

// GetBid retrieves a bid by its ID
func (s *BidService) GetBid(bidID uint) (*models.Bid, error) {
	return s.repo.GetBidByID(bidID)
}

// ListBids retrieves all bids for a specific author
func (s *BidService) ListBids(authorID uint) ([]models.Bid, error) {
	return s.repo.ListBidsForAuthor(authorID)
}

// EditBid increments the version and updates bid status
func (s *BidService) EditBid(bid models.Bid) error {
	bid.Version++
	return s.repo.UpdateBidStatus(bid)
}

// UpdateBidStatus updates the status of a bid
func (s *BidService) UpdateBidStatus(bid models.Bid) error {
	return s.repo.UpdateBidStatus(bid)
}

// DeleteBid deletes a bid by its ID
func (s *BidService) DeleteBid(bidID uint) error {
	return s.repo.DeleteBid(bidID)
}

// SubmitBidDecision updates the decision for a bid
func (s *BidService) SubmitBidDecision(bidID uint, decision string) error {
	log.Printf("Updating decision for bidID: %d with decision: %s", bidID, decision)
	err := s.repo.SubmitBidDecision(bidID, decision)
	if err != nil {
		log.Printf("Error updating bid decision: %v", err)
	}
	return err
}

// SubmitBidFeedback submits feedback for a bid
func (s *BidService) SubmitBidFeedback(bidID uint, feedback models.Review) error {
	// Retrieve the bid from the database
	bid, err := s.repo.GetBidByID(bidID)
	if err != nil {
		return fmt.Errorf("failed to retrieve bid: %v", err)
	}

	// Check that the reviewer is not the author of the bid
	if fmt.Sprintf("%d", bid.AuthorID) == feedback.Reviewer {
		return errors.New("authors cannot submit feedback for their own bids")
	}

	// Set feedback fields
	feedback.BidID = bidID
	feedback.CreatedAt = time.Now()
	feedback.UpdatedAt = time.Now()

	// Save the feedback to the database
	err = s.reviewRepo.CreateReview(feedback)
	if err != nil {
		return fmt.Errorf("failed to submit feedback: %v", err)
	}

	return nil
}

// RollbackBidVersion rolls back a bid to a previous version
func (s *BidService) RollbackBidVersion(bidID uint, version int) error {
	return s.repo.RollbackBidVersion(bidID, version)
}

// GetBidReviews retrieves all reviews for a specific bid
func (s *BidService) GetBidReviews(bidID uint) ([]models.Review, error) {
	return s.repo.GetBidReviews(bidID)
}
