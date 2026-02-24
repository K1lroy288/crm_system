package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"user-service/model"
	"user-service/service"
	"user-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	var req model.AuthDTO
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

/* func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req model.UserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at register request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON at register request"})
		return
	}

} */

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req model.UserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at register request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON at register request"})
		return
	}

	err := h.service.CreateUser(req)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			log.Printf("Duplicate username: %v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}

		log.Printf("create user error: %v", err)
		ctx.Status(http.StatusInternalServerError)
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

func (h *UserHandler) GetUserInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error parse id url param: %v", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserInfo(uint(idInt))
	if err != nil {
		log.Printf("error get user info: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var req model.UserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at update user info request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error parse id url param: %v", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	req.ID = uint(idInt)

	err = h.service.UpdateUser(req)
	if err != nil {
		log.Printf("Error update user info: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandler) ChangePassword(ctx *gin.Context) {
	var req model.PasswordDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid JSON at update user info request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	claims, err := utils.ValidateJWT(ctx)
	if err != nil {
		log.Printf("Invalid jwt token change password: %v", err)
		ctx.Status(http.StatusUnauthorized)
		return
	}

	err = h.service.ChangePassword(claims.UserID, req)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			log.Printf("Invalid current password")
			ctx.Status(http.StatusBadRequest)
			return
		}
		log.Printf("change password error: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *UserHandler) GetRoles(ctx *gin.Context) {
	roles, err := h.service.GetRoles()
	if err != nil {
		log.Printf("get roles error: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, roles)
}

func (h *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		log.Printf("get users error: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := h.service.DeleteUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Status(http.StatusNotFound)
			return
		}
		log.Printf("delete user error: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusNoContent)
}
