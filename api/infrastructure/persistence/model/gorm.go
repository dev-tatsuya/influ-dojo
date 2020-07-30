package model

import "time"

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type GormModel interface {
	IsNew() bool
	AttachID() error
}
