package user

import (
	_ "embed"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/user/dto"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserServiceImpl struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{DB: db}
}

func (service UserServiceImpl) Login(requestBody *dto.LoginDTO) *dto.TokenDTO {
	user := User{Email: requestBody.Email}
	tx := service.DB.Where("email = ?", requestBody.Email).Find(&user)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: "UNAUTHORIZED"})
	}
	password := salt + requestBody.Password + salt
	log.Printf("password %s", password)
	log.Printf("user hased pass %s", user.Password)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		panic(exception.ValidationException{Message: "UNAUTHORIZED"})
	}

	return &dto.TokenDTO{
		AccessToken:  generateToken(&user, false),
		RefreshToken: generateToken(&user, true),
	}
}

//go:embed salt
var salt string

func (service UserServiceImpl) Register(reqBody *dto.RegisterDTO) *dto.TokenDTO {
	email := reqBody.Email
	password := reqBody.Password
	log.Printf("Email %s, Password %s", email, password)
	password = salt + password + salt
	log.Printf("Salted password %s", password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	utils.LogIfError(err)

	id, err := uuid.NewUUID()
	utils.PanicIfError(err)
	user := User{
		Id:       id.String(),
		Email:    email,
		Password: string(hashedPassword),
	}
	tx := service.DB.Create(&user)
	if tx.Error != nil {
		panic(exception.ValidationException{Message: "REGISTRATION_FAILED"})
	}
	return &dto.TokenDTO{
		AccessToken:  generateToken(&user, false),
		RefreshToken: generateToken(&user, true),
	}
}

func (service UserServiceImpl) Get(userId string) *dto.ProfileDTO {
	user := User{Id: userId}
	tx := service.DB.Find(&user)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: "UNAUTHORIZED"})
	}

	profileDTO := dto.ProfileDTO{
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		MobilePhoneNumber: user.MobilePhoneNumber,
	}
	return &profileDTO
}

func (service UserServiceImpl) Update(userId string, dto *dto.ProfileDTO) {
	user := User{Id: userId}
	tx := service.DB.Find(&user)
	if tx.RowsAffected < 1 {
		panic(exception.ValidationException{Message: "UNAUTHORIZED"})
	}
	log.Printf("user id %s", userId)
	log.Printf("name %s", dto.Name)
	log.Printf("mobile phone number %s", dto.MobilePhoneNumber)
	log.Printf("username %s", dto.Username)
	user.Name = dto.Name
	user.MobilePhoneNumber = dto.MobilePhoneNumber
	user.Email = dto.Email
	user.Username = dto.Username
	save := service.DB.Save(&user)
	utils.PanicIfError(save.Error)
}

func (service UserServiceImpl) Token(reqBody dto.TokenDTO) *dto.TokenDTO {
	log.Printf("Refresh token %s", reqBody.RefreshToken)
	token := VerifyToken(reqBody.RefreshToken)
	claims := token.Claims.(jwt.MapClaims)
	i := claims["userId"]
	user := User{Id: i.(string)}
	return &dto.TokenDTO{
		AccessToken:  generateToken(&user, false),
		RefreshToken: generateToken(&user, true),
	}
}

func VerifyToken(tokenStr string) *jwt.Token {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("INVALID_ALGORITHM")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		panic(exception.AuthenticationException{Message: "UNAUTHORIZED"})
	}
	return token
}

//go:embed jwtsecret
var secret string

func generateToken(user *User, isRefreshToken bool) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	if isRefreshToken {
		claims["exp"] = time.Now().Add(48 * time.Hour).UnixNano()
	} else {
		claims["exp"] = time.Now().Add(10 * time.Minute).UnixNano()
	}
	//claims["authorized"] = true
	claims["userId"] = user.Id

	signedString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Printf("key invalid %s", secret)
	}
	return signedString
}
