
type Todo struct {
	Id        string  `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	CreatedAt string  `json:"created_at" db:"created_at"`
	FinishAt  string  `json:"finish_at" db:"finish_at"`
	ParentId  *string `json:"parent_id" db:"parent_id"`
	IsDone    bool    `json:"is_done" db:"is_done"`
	Deleted   bool    `json:"deleted" db:"deleted"`
}
