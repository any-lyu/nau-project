package model

import "time"

const (
	MaxID = 2 << 32
)

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginReply struct {
	Token          string      `json:"token"`
	User           interface{} `json:"user"`
	Permission     uint8       `json:"permission"`
	PermissionName string      `json:"permission_name"`
}

func GetPermissionName(permission uint8) string {
	if permission == 2 {
		return "teacher"
	}
	return "student"
}

type NoticeListRequest struct {
	Cursor string `json:"cursor"`
}

type NoticeListReply struct {
	List   interface{} `json:"list"`
	Tips   string      `json:"tips"`
	Cursor string      `json:"cursor"`
}

type NoticePublishRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Level   uint   `json:"level"`
}

type NoticePublishReply struct {
	ID int64 `json:"id"`
}

type ClassListRequest struct {
}

type ClassListReply struct {
	List interface{} `json:"list"`
}

type HomeWorkListRequest struct {
	Cursor string `json:"cursor"`
}

type HomeWorkListReply struct {
	List   []*Homework `json:"list"`
	Cursor string      `json:"cursor"`
}

type HomeWorkPublishRequest struct {
	ClassID uint8  `json:"class_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Pics    []Pic  `json:"pics"`
}
type HomeWorkPublishReply struct {
	ID int64 `json:"id"`
}

type HomeWorkSubmitRequest struct {
	UserID     uint64 `json:"-"`
	HomeworkID int64  `json:"homework_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Pics       []Pic  `json:"pics"`
}

type HomeWorkSubmitReply struct {
	ID int64 `json:"id"`
}

type UserHomeWorkDetailRequest struct {
	UserID     uint64 `json:"-"`
	HomeworkID int64  `json:"homework_id"`
}

type Homework struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Pics      []Pic     `json:"pics"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type UserHomeWorkDetailReply struct {
	Status       int       `json:"status"` // 0-未提交 1-已提交
	Homework     *Homework `json:"homework"`
	UserHomework *Homework `json:"user_homework"`
}

type UserHomeWorkRequest struct {
	Cursor string `json:"cursor"`
}

type UserHomeWorkListReply struct {
	List   []*Homework `json:"list"`
	Cursor string      `json:"cursor"`
}
