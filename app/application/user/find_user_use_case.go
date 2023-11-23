package user

import (
	"time"

	userDomain "github.com/maitakedayo/go-layeredarchitecture/app/domain/user"
)

// DTO domain領域(User)からAppli領域にDTO変換
type FindUseCaseDto struct {
	ID        string
	LastName  string
	FirstName string
	Email     string
	Posts     []string   // Change: Add a field for posts
	Idlimit   time.Time
}
//---e-

// ---appli領域-s-
type FindUserUseCase struct {
	userServ *userDomain.UserService // domain領域を依存注入でレンタル
}
//
func NewFindUserUseCase(userServ *userDomain.UserService) *FindUserUseCase {
	return &FindUserUseCase{userServ: userServ}
}
//
func (uc *FindUserUseCase) Run(dto FindUseCaseDto) (*FindUseCaseDto, error) {
	user, err := uc.userServ.Repository.FindByFullName(dto.LastName, dto.FirstName)
	if err != nil {
		return nil, err
	}

	return &FindUseCaseDto{
		ID:        user.ID(),
		LastName:  user.LastName(),
		FirstName: user.FirstName(),
		Email:     user.Email(),
		Posts:     user.Posts(),
		Idlimit:   user.Idlimit(),
	}, nil
}
//---e-