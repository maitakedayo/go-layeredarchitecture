package user

import (
	userDomain "github.com/maitakedayo/go-layeredarchitecture/app/domain/user"
)

// DTO domain領域(User)からAppli領域にDTO変換
type FindAllUsersUseCaseDto struct {
	Users []*FindUseCaseDto
}
//---e-

// ---appli領域-s-
type FindAllUsersUseCase struct {
	userServ *userDomain.UserService // domain領域を依存注入でレンタル
}
//
func NewFindAllUsersUseCase(userServ *userDomain.UserService) *FindAllUsersUseCase {
	return &FindAllUsersUseCase{userServ: userServ}
}
//
func (uc *FindAllUsersUseCase) Run() (*FindAllUsersUseCaseDto, error) {
	users, err := uc.userServ.Repository.FindAllUsers()
	if err != nil {
		return nil, err
	}

	var userDtos []*FindUseCaseDto
	for _, user := range users {
		userDto := &FindUseCaseDto{
			ID:        user.ID(),
			LastName:  user.LastName(),
			FirstName: user.FirstName(),
			Email:     user.Email(),
			Posts:     user.Posts(),
			Idlimit:   user.Idlimit(),
		}
		userDtos = append(userDtos, userDto)
	}

	return &FindAllUsersUseCaseDto{Users: userDtos}, nil
}
//---e-
