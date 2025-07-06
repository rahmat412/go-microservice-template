package dto

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date" validate:"required"`
	IsActive  bool   `json:"is_active" validate:"required"`
}

type CreateUserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	CreatedAt string `json:"created_at"`
	IsActive  bool   `json:"is_active"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date" validate:"required"`
}

type UpdateUserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
}

type GetUserByIDResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
	CreatedAt string `json:"created_at"`
	IsActive  bool   `json:"is_active"`
}
