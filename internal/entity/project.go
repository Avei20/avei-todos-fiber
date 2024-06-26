package entity

type Project struct {
	Id          string `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Deleted     bool   `json:"deleted" db:"deleted"`
}
