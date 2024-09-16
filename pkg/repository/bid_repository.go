package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ohhart/tender-restapi/models"
)

type BidRepository struct {
	db *sqlx.DB
}

func NewBidRepository(db *sqlx.DB) *BidRepository {
	return &BidRepository{db: db}
}

func (r *BidRepository) CreateBid(bid models.Bid) error {
	query := `
        INSERT INTO bids (name, description, status, tender_id, author_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, version
    `

	err := r.db.QueryRow(query,
		bid.Name, bid.Description, bid.Status, bid.TenderID,
		bid.AuthorID, bid.CreatedAt, bid.UpdatedAt).Scan(&bid.ID, &bid.Version)

	if err != nil {
		log.Printf("Error executing CreateBid query: %v", err)
		return err
	}

	return nil
}

func (r *BidRepository) GetBidByID(id uint) (*models.Bid, error) {
	var bid models.Bid
	query := `SELECT * FROM bids WHERE id = $1`
	err := r.db.Get(&bid, query, id)

	if err != nil {
		return nil, err
	}

	return &bid, nil
}

func (r *BidRepository) UpdateBidStatus(bid models.Bid) error {
	query := `
        UPDATE bids
        SET name = $1, description = $2, version = $3, updated_at = $4
        WHERE id = $5
    `

	_, err := r.db.Exec(query,
		bid.Name, bid.Description, bid.Version, time.Now(), bid.ID)

	return err
}

func (r *BidRepository) DeleteBid(id uint) error {
	query := `DELETE FROM bids WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *BidRepository) ListBidsForAuthor(authorID uint) ([]models.Bid, error) {
	var bids []models.Bid
	query := `SELECT * FROM bids WHERE author_id = $1`
	err := r.db.Select(&bids, query, authorID)

	if err != nil {
		return nil, err
	}

	return bids, nil
}

func (r *BidRepository) RollbackBidVersion(bidID uint, version int) error {
	query := `
        UPDATE bids
        SET version = version + 1, updated_at = $1
        WHERE id = $2 AND version = $3
    `

	result, err := r.db.Exec(query, time.Now(), bidID, version)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("bid version not found")
	}

	return nil
}

func (r *BidRepository) SubmitBidDecision(bidID uint, decision string) error {
	query := "UPDATE bids SET decision = $1 WHERE id = $2"
	_, err := r.db.Exec(query, decision, bidID)
	return err
}

func (r *BidRepository) GetBidReviews(bidID uint) ([]models.Review, error) {
	var reviews []models.Review
	query := `SELECT * FROM reviews WHERE bid_id = $1`
	err := r.db.Select(&reviews, query, bidID)

	if err != nil {
		return nil, err
	}

	return reviews, nil
}
