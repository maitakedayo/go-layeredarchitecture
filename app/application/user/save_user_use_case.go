package user

import (
	"time"

	userDomain "github.com/maitakedayo/go-layeredarchitecture/app/domain/user"
)

// DTO domain領域(User)をappli領域にDTO変換
type SaveUseCaseDto struct {
	LastName  string
	FirstName string
	Email     string
	Posts     []string   // Change: Add a field for posts
	Idlimit   time.Time
}
//---e-

// ---appli領域-s-
type SaveUserUseCase struct {
	userServ *userDomain.UserService // domain領域を依存注入でレンタル
}
//
func NewSaveUserUseCase(userServ *userDomain.UserService) *SaveUserUseCase {
	return &SaveUserUseCase{userServ: userServ}
}
//
func (uc *SaveUserUseCase) Run(dto SaveUseCaseDto) error {
	user, err := userDomain.NewUser(dto.LastName, dto.FirstName, dto.Email, dto.Posts, dto.Idlimit)
	if err != nil {
		return err
	}
	return uc.userServ.Repository.Save(user)
}
//---e-
