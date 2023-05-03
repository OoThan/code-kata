package repository

import (
	"loan-back-services/pkg/ds"
)

type Repository struct {
	DS    *ds.DataSource
	Admin *adminRepository
}

type RepoConfig struct {
	DS *ds.DataSource
}

func NewRepository(c *RepoConfig) *Repository {
	return &Repository{
		DS:    c.DS,
		Admin: newAdminRepository(c),
	}
}
