package structs

import "time"

type Categories struct {
	ID        int    `json:"id"`
	Name       string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string `json:"modified_by"`
}