package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"it-order-app/backend/model"
	"it-order-app/backend/repository"
	"it-order-app/backend/service"
	"net/http"
)

// DTO (Data Transfer Object) untuk input registrasi
type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthHandler struct {
	userRepository repository.UserRepository
	jwtService     service.JWTService
}

func NewAuthHandler(userRepo repository.UserRepository, jwtService service.JWTService) *AuthHandler {
	return &AuthHandler{
		userRepository: userRepo,
		jwtService:     jwtService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input RegisterInput

	// Validasi input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Buat objek user baru
	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "user", // Default role
	}

	// Simpan ke database via repository
	savedUser, err := h.userRepository.Save(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	// Kirim response sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
		"data": gin.H{
			"id":    savedUser.ID,
			"name":  savedUser.Name,
			"email": savedUser.Email,
			"role":  savedUser.Role,
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input LoginInput

	// Validasi input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 1. Cari user berdasarkan email
	user, err := h.userRepository.FindByEmail(input.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// 2. Bandingkan password yang di-hash dengan password inputan
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		// Password tidak cocok
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// 3. Jika password cocok, generate token JWT
	token, err := h.jwtService.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 4. Kirim token sebagai response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
