package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/gin-api-rest/database"
	"github.com/murilorscampos/gin-api-rest/models"
)

func ExibeTodosAlunos(c *gin.Context) {

	var alunos []models.Aluno

	database.DB.Order("nome ASC").Find(&alunos)
	c.JSON(http.StatusOK, alunos)

}

func BuscaAlunoPorID(c *gin.Context) {

	var aluno models.Aluno

	id := c.Params.ByName("id")

	database.DB.Order("nome ASC").First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Mesage:": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)

}

func Saudacao(c *gin.Context) {

	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})

}

func CriaNovoAluno(c *gin.Context) {

	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Mesage": err.Error()})

		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)

}
func DeletaAluno(c *gin.Context) {

	var aluno models.Aluno

	id := c.Params.ByName("id")

	linhasAfetadas := database.DB.Delete(&aluno, id).RowsAffected

	if linhasAfetadas == 0 {

		c.JSON(http.StatusOK, gin.H{"Mesage": "Aluno ID: " + id + " não encontrado."})

		return

	}

	c.JSON(http.StatusOK, gin.H{"Mesage:": "Aluno excluído."})

}
func EditaAluno(c *gin.Context) {

	var aluno models.Aluno

	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"Mesage": err.Error()})

		return

	}

	linhasAfetadas := database.DB.Model(&aluno).UpdateColumns(aluno).RowsAffected

	if linhasAfetadas == 0 {

		c.JSON(http.StatusOK, gin.H{"Mesage": "Aluno ID: " + id + " não encontrado."})

		return

	}

	c.JSON(http.StatusOK, gin.H{"Mesage:": "Aluno alterado."})

}

func BuscaAlunoPorCPF(c *gin.Context) {

	var aluno models.Aluno

	cpf := c.Param("cpf")
	linhaAfetada := database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno).RowsAffected

	if linhaAfetada == 0 {

		c.JSON(http.StatusOK, gin.H{
			"Mesage": "Aluno não encontrado.",
		})

		return
	}

	c.JSON(http.StatusOK, aluno)

}
