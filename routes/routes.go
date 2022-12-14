package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/gin-api-rest/controllers"
)

func HandleRequests() {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/", controllers.RedirecionaPagina)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.GET("/:nome", controllers.Saudacoes)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.NoRoute(controllers.RotaNaoEncontrada)

	r.Run() // listen and serve on localhost:8080

}
