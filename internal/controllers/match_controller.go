package controllers

import (
	"net/http"
	"strconv"
	"time"

	"laliga-api/internal/config"
	"laliga-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MatchController struct {
	db *gorm.DB
}

func NewMatchController() *MatchController {
	return &MatchController{
		db: config.DB,
	}
}

// GetMatches obtiene todos los partidos
func (c *MatchController) GetMatches(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	var matches []models.Match
	if err := c.db.Find(&matches).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los partidos"})
		return
	}
	ctx.JSON(http.StatusOK, matches)
}

// GetMatch obtiene un partido por ID
func (c *MatchController) GetMatch(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var match models.Match
	if err := c.db.First(&match, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, match)
}

// CreateMatch crea un nuevo partido
func (c *MatchController) CreateMatch(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	var input struct {
		HomeTeam  string `json:"homeTeam" binding:"required"`
		AwayTeam  string `json:"awayTeam" binding:"required"`
		MatchDate string `json:"matchDate" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Convertir la fecha string a time.Time
	matchDate, err := time.Parse("2006-01-02", input.MatchDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha inválido. Use el formato YYYY-MM-DD"})
		return
	}

	match := models.Match{
		HomeTeam:  input.HomeTeam,
		AwayTeam:  input.AwayTeam,
		MatchDate: matchDate,
	}

	if err := c.db.Create(&match).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el partido"})
		return
	}
	ctx.JSON(http.StatusCreated, match)
}

// UpdateMatch actualiza un partido existente
func (c *MatchController) UpdateMatch(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PUT, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		HomeTeam  string `json:"homeTeam" binding:"required"`
		AwayTeam  string `json:"awayTeam" binding:"required"`
		MatchDate string `json:"matchDate" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Convertir la fecha string a time.Time
	matchDate, err := time.Parse("2006-01-02", input.MatchDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fecha inválido. Use el formato YYYY-MM-DD"})
		return
	}

	var match models.Match
	if err := c.db.First(&match, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	match.HomeTeam = input.HomeTeam
	match.AwayTeam = input.AwayTeam
	match.MatchDate = matchDate

	if err := c.db.Save(&match).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el partido"})
		return
	}
	ctx.JSON(http.StatusOK, match)
}

// DeleteMatch elimina un partido
func (c *MatchController) DeleteMatch(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.db.Delete(&models.Match{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el partido"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Partido eliminado correctamente"})
}

// RegisterYellowCard registra una tarjeta amarilla en un partido
func (c *MatchController) RegisterYellowCard(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PATCH, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var match models.Match
	if err := c.db.First(&match, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	match.YellowCards++
	if err := c.db.Save(&match).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar la tarjeta amarilla"})
		return
	}
	ctx.JSON(http.StatusOK, match)
}

// RegisterRedCard registra una tarjeta roja en un partido
func (c *MatchController) RegisterRedCard(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PATCH, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var match models.Match
	if err := c.db.First(&match, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	match.RedCards++
	if err := c.db.Save(&match).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar la tarjeta roja"})
		return
	}
	ctx.JSON(http.StatusOK, match)
}

// SetExtraTime establece el tiempo extra para un partido
func (c *MatchController) SetExtraTime(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PATCH, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Minutes int `json:"minutes" binding:"required,min=1,max=30"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	var match models.Match
	if err := c.db.First(&match, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	match.ExtraTime = input.Minutes
	if err := c.db.Save(&match).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al establecer el tiempo extra"})
		return
	}
	ctx.JSON(http.StatusOK, match)
}

// RegisterGoal registra un gol en un partido
func (c *MatchController) RegisterGoal(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "PATCH, OPTIONS")
	ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Team string `json:"team" binding:"required,oneof=home away"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	var match models.Match
	if err := c.db.First(&match, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
		return
	}

	if input.Team == "home" {
		match.HomeGoals++
	} else {
		match.AwayGoals++
	}

	if err := c.db.Save(&match).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el gol"})
		return
	}
	ctx.JSON(http.StatusOK, match)
}
