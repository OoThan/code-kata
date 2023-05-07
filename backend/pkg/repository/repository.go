package repository

import (
	"loan-back-services/pkg/ds"
)

type Repository struct {
	DS          *ds.DataSource
	Admin       *adminRepository
	User        *userRepository
	LoanPackage *loanPackageRepository
	UserLoan    *userLoanRepository
}

type RepoConfig struct {
	DS *ds.DataSource
}

func NewRepository(c *RepoConfig) *Repository {
	return &Repository{
		DS:          c.DS,
		Admin:       newAdminRepository(c),
		User:        newUserRepository(c),
		LoanPackage: newLoanPackageRepository(c),
		UserLoan:    newUserLoanRepository(c),
	}
}
