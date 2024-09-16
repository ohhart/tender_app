package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/ohhart/tender-restapi/models"
)

type TenderRepository struct {
	db *sqlx.DB
}

func NewTenderRepository(db *sqlx.DB) *TenderRepository {
	return &TenderRepository{db: db}
}

func (r *TenderRepository) CreateTender(tender models.Tender) error {
	query := `
        INSERT INTO tenders (name, description, service_type, status, version, organization_id, creator_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, version
    `
	err := r.db.QueryRow(query,
		tender.Name, tender.Description, tender.ServiceType,
		tender.Status, tender.Version, tender.OrganizationID,
		tender.CreatorID, tender.CreatedAt, tender.UpdatedAt,
	).Scan(&tender.ID, &tender.Version)

	if err != nil {
		log.Printf("Error executing CreateTender query: %v", err)
		return err
	}
	return nil
}

func (r *TenderRepository) UpdateTenderStatus(tenderID uint, status string) error {
	query := `
        UPDATE tenders 
        SET status = $1, updated_at = NOW()
        WHERE id = $2
    `
	_, err := r.db.Exec(query, status, tenderID)
	return err
}

func (r *TenderRepository) EditTender(tender models.Tender) error {
	query := `
        UPDATE tenders
        SET name = $1, description = $2, version = $3, updated_at = $4
        WHERE id = $5
    `
	_, err := r.db.Exec(query,
		tender.Name, tender.Description, tender.Version, tender.UpdatedAt, tender.ID,
	)

	return err
}

func (r *TenderRepository) DeleteTender(id uint) error {
	query := `DELETE FROM tenders WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *TenderRepository) ListTenders() ([]models.Tender, error) {
	var tenders []models.Tender
	query := `SELECT * FROM tenders`

	if err := r.db.Select(&tenders, query); err != nil {
		return nil, err
	}
	return tenders, nil
}

func (r *TenderRepository) GetTendersByOrganization(orgID uint) ([]models.Tender, error) {
	var tenders []models.Tender
	query := `SELECT * FROM tenders WHERE organization_id = $1`

	if err := r.db.Select(&tenders, query, orgID); err != nil {
		return nil, err
	}
	return tenders, nil
}

func (r *TenderRepository) RollbackTenderVersion(id uint, version int) error {
	query := `
        UPDATE tenders
        SET version = version + 1
        WHERE id = $1 AND version = $2
    `
	result, err := r.db.Exec(query, id, version)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("tender version not found")
	}
	return nil
}

func (r *TenderRepository) GetTenderByID(id uint) (*models.Tender, error) {
	var tender models.Tender
	query := `SELECT * FROM tenders WHERE id = $1`

	if err := r.db.Get(&tender, query, id); err != nil {
		return nil, err
	}

	reviewsQuery := `SELECT * FROM reviews WHERE tender_id = $1`
	var reviews []models.Review
	if err := r.db.Select(&reviews, reviewsQuery, id); err == nil {
		tender.Reviews = reviews
	}

	return &tender, nil
}

func (r *TenderRepository) GetUserIDByUsername(username string) (uint, error) {
	var userID uint
	query := `SELECT id FROM employee WHERE username = $1`

	if err := r.db.Get(&userID, query, username); err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *TenderRepository) GetTendersByUsername(username string) ([]models.Tender, error) {
	userID, err := r.GetUserIDByUsername(username)
	if err != nil {
		return nil, err
	}

	var tenders []models.Tender
	query := `SELECT * FROM tenders WHERE creator_id = $1`

	if err = r.db.Select(&tenders, query, userID); err != nil {
		return nil, err
	}
	return tenders, nil
}
