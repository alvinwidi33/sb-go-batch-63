package structs

import "github.com/google/uuid"

type Users struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	IsActive bool      `json:"is_active"`
}

type Customer struct {
	CustomerID uuid.UUID `json:"customer_id"`
	UserID		uuid.UUID `json:"user_id"`
	User	   *Users	 `json:"user"`
	Status     string    `json:"status"`
}

type Admin struct {
	AdminID uuid.UUID `json:"admin_id"`
	UserID		uuid.UUID `json:"user_id"`
	User	   *Users	 `json:"user"`
}
