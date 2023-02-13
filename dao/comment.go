package dao

type Comment struct {
	Id         int64    `json:"id,omitempty"`
	User       UserInfo `json:"user"`
	Content    string   `json:"content,omitempty"`
	CreateDate string   `json:"create_date,omitempty"`
}
