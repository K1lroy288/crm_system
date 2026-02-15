package handler

import (
	"log"
	"net/http"
	"user-service/model"
	"user-service/service"
	"user-service/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req model.UserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at login request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	user, err := h.service.GetUserByUsername(req.Username)
	if err != nil {
		log.Printf("Invalid username or password: %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(req.Password)); err != nil {
		log.Printf("Invalid username or password: %v", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		log.Printf("JWT generation failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
		return
	}

	response := map[string]string{"token": token}
	ctx.JSON(http.StatusOK, response)
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req model.UserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at register request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON at register request"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Hashed password generation failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process registration"})
		return
	}

	user := model.User{
		Username:     req.Username,
		PasswordHash: hashedPassword,
	}
	exist, err := h.service.CreateUser(&user)
	if err != nil {
		log.Printf("Exist user check failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	if exist {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User with such username already exists"})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *UserHandler) GetUserByLastname(ctx *gin.Context) {
	lastname := ctx.Param("userid")

	user, err := h.service.GetUserByLastname(lastname)
	if err != nil {
		log.Printf("error get user by lastname: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error get user by lastname"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) GetMasters(ctx *gin.Context) {
	masters, err := h.service.GetMasters()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, masters)
}

func (h *UserHandler) GetMastersByIDs(ctx *gin.Context) {
	var req []uint
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at login request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	masters, err := h.service.GetMastersByIDs(req)
	if err != nil {
		log.Printf("Error get masters from database: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, masters)
}
