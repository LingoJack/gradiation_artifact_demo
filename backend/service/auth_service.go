package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lingojack/taobao_clone/config"
	"github.com/lingojack/taobao_clone/model"
	"github.com/lingojack/taobao_clone/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db  *gorm.DB
	cfg *config.Config
	dao *repository.UsersDao
}

func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{
		db:  db,
		cfg: cfg,
		dao: repository.NewUsersDao(db),
	}
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=4,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Nickname string `json:"nickname"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type JWTClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthResponse struct {
	User  *model.Users `json:"user"`
	Token string       `json:"token"`
}

func (s *AuthService) Register(ctx context.Context, req *RegisterRequest) (*AuthResponse, error) {
	// 检查用户名是否已存在
	var count int64
	s.db.Model(&model.Users{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	s.db.Model(&model.Users{}).Where("email = ?", req.Email).Count(&count)
	if count > 0 {
		return nil, errors.New("邮箱已被注册")
	}

	// 检查手机号是否已存在
	s.db.Model(&model.Users{}).Where("phone = ?", req.Phone).Count(&count)
	if count > 0 {
		return nil, errors.New("手机号已被注册")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户
	now := time.Now()
	status := int8(1)
	user := &model.Users{
		Username:  req.Username,
		Email:     &req.Email,
		Phone:     &req.Phone,
		Password:  string(hashedPassword),
		Nickname:  &req.Nickname,
		Status:    &status,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := s.dao.Insert(ctx, user); err != nil {
		return nil, err
	}

	// 生成 token
	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{User: user, Token: token}, nil
}

func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*AuthResponse, error) {
	// 支持用户名/邮箱/手机号登录
	var user model.Users
	err := s.db.Where("username = ? OR email = ? OR phone = ?", req.Username, req.Username, req.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("密码错误")
	}

	// 检查用户状态
	if user.Status == nil || *user.Status != 1 {
		return nil, errors.New("用户已被禁用")
	}

	// 生成 token
	token, err := s.GenerateToken(&user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{User: &user, Token: token}, nil
}

func (s *AuthService) GenerateToken(user *model.Users) (string, error) {
	userID := uint64(0)
	if user.Id != nil {
		userID = *user.Id
	}
	claims := &JWTClaims{
		UserID:   userID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "taobao_clone",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.Middleware.Auth.JWTSecret))
}

func (s *AuthService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.Middleware.Auth.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func (s *AuthService) GetUserByID(ctx context.Context, userID uint64) (*model.Users, error) {
	var user model.Users
	err := s.db.First(&user, userID).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user, nil
}
