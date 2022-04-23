package models

import "time"

type User struct {
	UserID uint64 `pg:",pk" json:"user_id"`
	Type int8 `pg:"type:smallint,default:1,notnull" json:"type"`
	Username string `pg:"type:varchar(16),unique,notnull" json:"username"`
	Password string `pg:"type:varchar(60),notnull" json:"-"`
	APIKey string       `pg:"type:varchar(60),default:gen_random_uuid(),notnull" json:"api_key"`
	Clips []*Clip       `pg:"rel:has-many" json:"clips"`
	CreatedAt time.Time `pg:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
