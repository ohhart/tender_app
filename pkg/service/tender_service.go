package service

import (
	"log"
	"time"

	"github.com/ohhart/tender-restapi/models"
	"github.com/ohhart/tender-restapi/pkg/repository"
)

// TenderService handles business logic for tenders
type TenderService struct {
	repo *repository.TenderRepository
}

// NewTenderService creates a new TenderService
func NewTenderService(repo *repository.TenderRepository) *TenderService {
	return &TenderService{repo: repo}
}

// CreateTender creates a new tender and sets CreatedAt and UpdatedAt to the current time
func (s *TenderService) CreateTender(tender models.Tender) error {
	tender.CreatedAt = time.Now()
	tender.UpdatedAt = time.Now()

	err := s.repo.CreateTender(tender)
	if err != nil {
		log.Printf("Error creating tender: %v", err)
		return err
	}
	log.Printf("Tender created successfully: %+v", tender)
	return nil
}

// GetTender retrieves a tender by its ID
func (s *TenderService) GetTender(id uint) (*models.Tender, error) {
	return s.repo.GetTenderByID(id)
}

// UpdateTenderStatus updates the status of a tender
func (s *TenderService) UpdateTenderStatus(tenderID uint, status string) error {
	return s.repo.UpdateTenderStatus(tenderID, status)
}

// EditTender updates an existing tender and increments its version
func (s *TenderService) EditTender(tender models.Tender) error {
	tender.UpdatedAt = time.Now()
	tender.Version++
	return s.repo.EditTender(tender)
}

// DeleteTender deletes a tender by its ID
func (s *TenderService) DeleteTender(id uint) error {
	return s.repo.DeleteTender(id)
}

// ListTenders retrieves all tenders
func (s *TenderService) ListTenders() ([]models.Tender, error) {
	return s.repo.ListTenders()
}

// GetTendersByOrganization retrieves tenders by organization ID
func (s *TenderService) GetTendersByOrganization(orgID uint) ([]models.Tender, error) {
	return s.repo.GetTendersByOrganization(orgID)
}

// RollbackTenderVersion rolls back a tender to a previous version
func (s *TenderService) RollbackTenderVersion(id uint, version int) error {
	return s.repo.RollbackTenderVersion(id, version)
}

// GetTendersByUsername retrieves tenders created by a specific user
func (s *TenderService) GetTendersByUsername(username string) ([]models.Tender, error) {
	return s.repo.GetTendersByUsername(username)
}
