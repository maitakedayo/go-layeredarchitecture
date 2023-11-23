package user

import (
	"errors"
	"fmt"
	"unicode/utf8"
	"net/mail"
	"time"

	"github.com/oklog/ulid/v2"
)

// --- app/domain/
// --- 処理型(エンティティ domain集約ルート) -s-
type User struct {
	id        string
	lastName  string
	firstName string
	email     string
	posts     []string   // Change: Add a field for posts
	idlimit   time.Time
}
// private
func newUser(id string, lastName string, firstName string, email string, posts []string, idlimit time.Time) (*User, error) {
	// ガード節
	if utf8.RuneCountInString(lastName) < nameLengthMin || utf8.RuneCountInString(lastName) > nameLengthMax {
		return nil, errors.New("名前（姓）の値が不正です。")
	}
	if utf8.RuneCountInString(firstName) < nameLengthMin || utf8.RuneCountInString(firstName) > nameLengthMax {
		return nil, errors.New("名前（名）の値が不正です。")
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errors.New("メールアドレスの値が不正です。")
	}
	if !isDateFormatValid(idlimit.Format("2006-01-02")) {
		return nil, errors.New("無効な日付形式です")
	}
	
	return &User{id: id, lastName: lastName, firstName: firstName, email: email, posts: posts, idlimit: idlimit}, nil
}
// 日付の形式が有効かどうかを検証する関数
func isDateFormatValid(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
//
const (
	nameLengthMax = 255
	nameLengthMin = 1
)
//
func NewUser(lastName string, firstName string, email string, posts []string, idlimit time.Time) (*User, error) {
	return newUser(NewULID(), lastName, firstName, email, posts, idlimit)
}
//
func Reconstruct(id string, lastName string, firstName string, email string, posts []string, idlimit time.Time) (*User, error) {
	return newUser(id, lastName, firstName, email, posts, idlimit)
}
//
func (u *User) ID() string {
	return u.id
}
//
func (u *User) LastName() string {
	return u.lastName
}
//
func (u *User) FirstName() string {
	return u.firstName
}
//
func (u *User) Email() string {
	return u.email
}
//
func (u *User) Posts() []string {
	return u.posts
}
//
func (u *User) Idlimit() time.Time {
	return u.idlimit
}
//
func (u *User) PrintUserDetails() {
	fmt.Printf("ID: %s\nLastName: %s\nFirstName: %s\nEmail: %s\nPost: %s\nIdlimit: %s\n\n",
	    u.ID(), u.LastName(), u.FirstName(), u.Email(), u.Posts(), u.Idlimit())
}
//---e-

//---pkg/ulid
func NewULID() string { //同じidを生成してしまう対策
	return ulid.Make().String()
}
//
func IsValid(s string) bool {
	_, err := ulid.Parse(s)
	return err == nil
}
//---e-