package domains

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID uuid.UUID `gorm:"type:char(36);" json:"id"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return nil
}

func (b *Base) InsertID(uuidString string) ErrorCode {
	var err ErrorCode
	b.ID, err.Err = uuid.Parse(uuidString)
	if err.Err != nil {
		err.StatusCode = 400
		return err
	}
	return err
}

func (b *Base) GetIDString() string {
	uuid := strings.Replace(b.ID.String(), "-", "", -1)
	return uuid
}
