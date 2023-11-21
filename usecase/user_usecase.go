package usecase

import (
	"os"
	"time"
	"verbme-api/model"
	"verbme-api/repository"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// ユーザーに関するビジネスロジックを記述する
type IUserUseCase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUseCase struct {
	ur repository.IUserRepository
}

func NewUserUseCase(ur repository.IUserRepository) IUserUseCase {
	return &userUseCase{ur}
}

// インターフェースを満たすためのメソッド
func (uu *userUseCase) SignUp(user model.User) (model.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10) //ハッシュ化
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Name: user.Name, Password: string(hash)} //インスタンス化
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{ //レスポンス用のインスタンス化
		ID:   newUser.ID,
		Name: newUser.Name,
	}
	return resUser, nil
}

func (uu *userUseCase) Login(user model.User) (string, error) {
	// if err := uu.uv.UserValidate(user); err != nil {
	// 	return "", err
	// }
	storedUser := model.User{}
	if err := uu.ur.GetUserByName(&storedUser, user.Name); err != nil { //ユーザー名からユーザー情報を取得
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)) //ハッシュ化されたパスワードと入力されたパスワードを比較
	if err != nil {
		return "", err
	}
	//JWTの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
