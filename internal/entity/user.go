package entity

type User struct {
	Id         string   `json:"id" db:"id"`
	Username   string   `json:"username" db:"username"`
	Password   string   `json:"password" db:"password"`
	Email      string   `json:"email" db:"email"`
	ProjectIds []string `json:"project_ids" db:"project_ids"`
}
