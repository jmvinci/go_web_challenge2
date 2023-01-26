package handler

import (
	loaddb "github.com/bootcamp-go/desafio-cierre-db.git/internal/load_db"
	"github.com/gin-gonic/gin"
)

type LoadDB struct {
	l loaddb.Service
}

func NewHandlerLoadDB(l loaddb.Service) *LoadDB {
	return &LoadDB{l}
}

func (l *LoadDB) LoadDB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := l.l.LoadDb()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, "Cargado")
	}
}

func (l *LoadDB) UpdateInvoices() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := l.l.UpdateInvoices()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, "Actualizado")
	}
}
