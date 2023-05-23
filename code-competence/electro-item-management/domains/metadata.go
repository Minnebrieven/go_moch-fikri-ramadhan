package domains

import (
	"time"

	"gorm.io/gorm"
)

type (
	DomainMetadata struct {
		DeletedAt gorm.DeletedAt `json:"DeletedAt"`
		CreatedAt time.Time      `json:"CreatedAt"`
		UpdatedAt time.Time      `json:"UpdatedAt"`
	}
)

