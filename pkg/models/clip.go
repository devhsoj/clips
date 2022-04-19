package models

import "time"

type Clip struct {
	ClipID uint64 `pg:",pk" json:"clip_id"`
	UserID uint64 `pg:",notnull" json:"user_id"`
	Creator string `pg:"type:varchar(16),notnull" json:"creator"`
	Source string `pg:"type:varchar(4096),notnull" json:"-"`
	Type string `pg:"type:varchar(255),notnull" json:"type"`
	Title string `pg:"type:varchar(128),notnull" json:"title"`
	Description string `pg:"type:varchar(512),notnull" json:"description"`
	Views uint64 `pg:"default:0,notnull" json:"views"`
	UploadedAt time.Time `pg:"default:CURRENT_TIMESTAMP" json:"uploaded_at"`
}
