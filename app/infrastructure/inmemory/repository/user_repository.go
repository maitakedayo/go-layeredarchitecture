package repository

import (
	"errors"
	"time"

	userDomain "github.com/maitakedayo/go-layeredarchitecture/app/domain/user"
)

// DTO domain領域(User)からAppli領域にDTO変換
type FindUseCaseRepoDto struct {
	ID        string
	LastName  string
	FirstName string
	Email     string
	Posts     []string   // Change: Add a field for posts
	Idlimit   time.Time  // Change: Change the type to time.Time
}
//---e-

// --- 具象実装(リポジトリ) -s-
type InMemoryUserRepository struct {
	users map[string]*userDomain.User
}
//
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*userDomain.User),
	}
}
//
func (r *InMemoryUserRepository) Save(user *userDomain.User) error {
	r.users[user.ID()] = user
	return nil
}
//
func (r *InMemoryUserRepository) FindByFullName(lastName string, firstName string) (*userDomain.User, error) {
	for _, user := range r.users {
		if user.LastName() == lastName && user.FirstName() == firstName {
			return user, nil
		}
	}
	return nil, errors.New("User not found")
}
//
func (r *InMemoryUserRepository) FindFirstUser() (*userDomain.User, error) {
	//面倒だからこれでいいや
	for _, user := range r.users {
		return user, nil
	}
	return nil, errors.New("No users found")
}
//
func (r *InMemoryUserRepository) FindAllUsers() ([]*userDomain.User, error) {
	var users []*userDomain.User
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
//---e-