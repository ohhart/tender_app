package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/ohhart/tender-restapi/models"
)

type ReviewRepository struct {
	db *sqlx.DB
}

func NewReviewRepository(db *sqlx.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) FindReviewsByTenderAndAuthor(tenderID uint, authorUsername string, organizationID uint) ([]models.Review, error) {
	var reviews []models.Review
	query := `
        SELECT * FROM reviews 
        WHERE tender_id = $1 AND author_username = $2 AND organization_id = $3
    `
	err := r.db.Select(&reviews, query, tenderID, authorUsername, organizationID)
	return reviews, err
}

func (r *ReviewRepository) CreateReview(review models.Review) error {
	query := `
        INSERT INTO reviews (bid_id, author_username, organization_id, comment, rating, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err := r.db.Exec(query,
		review.BidID, review.Reviewer, review.OrganizationID,
		review.Comment, review.Rating, review.CreatedAt, review.UpdatedAt,
	)
	return err
}
