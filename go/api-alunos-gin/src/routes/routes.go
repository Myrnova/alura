package routes

import (
	"api-alunos-gin/src/controllers"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func GinConfiguration() {
	r = gin.Default()
	r.Static("/assets", "./src/assets")
  r.LoadHTMLGlob("src/templates/*")
}

func HandleRequests() {
	GinConfiguration()
	r.GET("/", controllers.ExibePaginaIndex)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}
