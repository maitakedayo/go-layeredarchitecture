package user


// --- 抽象型(リポジトリ) -s-
type IUserRepository interface {
	Save(user *User) error
	FindByFullName(lastName string, firstName string) (*User, error)
	FindAllUsers() ([]*User, error)
	FindFirstUser() (*User, error)
}
// --- e-