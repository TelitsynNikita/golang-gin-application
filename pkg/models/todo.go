package models

type TodoList struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UserList struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	ListID int `json:"list_id"`
}

type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	ID     int `json:"id"`
	ListID int `json:"list_id"`
	ItemID int `json:"item_id"`
}
