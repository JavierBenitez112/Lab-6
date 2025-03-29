package main

import (
	"log"
	"os"

	"laliga-api/internal/config"
	"laliga-api/internal/controllers"

	"github.com/gin-gonic/gin"
)

// @title           La Liga API
// @version         1.0
// @description     API para gestionar partidos de La Liga
// @host           localhost:8086
// @BasePath       /api

func main() {
	// Inicializar la base de datos
	if err := config.InitDB(); err != nil {
		log.Fatal("Error al inicializar la base de datos:", err)
	}

	// Crear el router de Gin
	router := gin.Default()

	// Configurar CORS global
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Servir archivos estáticos
	router.StaticFile("/LaLigaTracker.html", "./LaLigaTracker.html")
	router.StaticFile("/swagger.yaml", "./swagger.yaml")

	// Crear el controlador de partidos
	matchController := controllers.NewMatchController()

	// Grupo de rutas API
	api := router.Group("/api")
	{
		// Rutas de partidos
		api.GET("/matches", matchController.GetMatches)
		api.GET("/matches/:id", matchController.GetMatch)
		api.POST("/matches", matchController.CreateMatch)
		api.PUT("/matches/:id", matchController.UpdateMatch)
		api.DELETE("/matches/:id", matchController.DeleteMatch)

		// Rutas PATCH para eventos del partido
		api.PATCH("/matches/:id/goals", matchController.RegisterGoal)
		api.PATCH("/matches/:id/yellowcards", matchController.RegisterYellowCard)
		api.PATCH("/matches/:id/redcards", matchController.RegisterRedCard)
		api.PATCH("/matches/:id/extratime", matchController.SetExtraTime)
	}

	// Obtener el puerto de las variables de entorno o usar el puerto por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar el servidor
	log.Printf("Servidor iniciado en el puerto %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
