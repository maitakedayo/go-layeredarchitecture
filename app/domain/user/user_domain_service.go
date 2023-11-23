package user

// --- domain領域 ドメインサービス用型(メソッドのみ) -s-
type UserService struct {
	Repository IUserRepository // 抽象化 infra領域を抽象依存注入でレンタル
}
//
func NewUserService(repository IUserRepository) *UserService {
	return &UserService{Repository: repository}
}
//---e-