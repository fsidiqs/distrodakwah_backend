package modelhelper

import "time"

type UID32Timestamp struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp;not null" json:"updated_at"`
}

type UID64TimeStamp struct {
	ID        structuint64 `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time    `gorm:"default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time    `gorm:"default:current_timestamp;not null" json:"updated_at"`
}

type ID64 struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `gorm:"default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp;not null" json:"updated_at"`
}

type IDTimestamp struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
