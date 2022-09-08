package persistence

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AcademicFieldRepositoryImpl struct {
	db    *sqlx.DB
	cache []string
}

func NewAcademicFieldRepositoryImpl(db *sqlx.DB) *AcademicFieldRepositoryImpl {
	return &AcademicFieldRepositoryImpl{db: db}
}

func (afR *AcademicFieldRepositoryImpl) Get() ([]string, error) {
	if len(afR.cache) > 0 {
		return afR.cache, nil
	}

	sql := `
		SELECT DISTINCT academic_field
		FROM subjects
		WHERE academic_field IS NOT NULL
		ORDER BY academic_field
	`
	if err := afR.db.Select(&afR.cache, sql); err != nil {
		return nil, fmt.Errorf("failed on select to `subejects` table: %w", err)
	}
	return afR.cache, nil
}
