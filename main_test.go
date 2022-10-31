package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/murilorscampos/gin-api-rest/controllers"
	"github.com/murilorscampos/gin-api-rest/database"
	"github.com/murilorscampos/gin-api-rest/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	rotas := gin.Default()

	return rotas

}

func CriaAlunoMock() {

	aluno := models.Aluno{Nome: "Nome do aluno teste", CPF: "12345678901", RG: "123456789"}

	database.DB.Create(&aluno)

	ID = int(aluno.ID)

}

func ApagaAlunoMock() {

	var aluno models.Aluno

	database.DB.Delete(&aluno, ID)

}

func TestVerificaStatusCodeSaudacoes(t *testing.T) {

	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacoes)

	req, _ := http.NewRequest("GET", "/murilo", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

	mockDaResposta := `{"Message":"Olá, murilo. É um prazer tê-lo aqui."}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))

}

func TestListandoTodosAlunos(t *testing.T) {

	database.ConectaComBancoDeDados()

	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

	defer ApagaAlunoMock()

}

func TestListaPorCPF(t *testing.T) {

	database.ConectaComBancoDeDados()

	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/cpf/:cpf", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

	defer ApagaAlunoMock()

}

func TestBuscaAlunoPorID(t *testing.T) {

	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)

	req, _ := http.NewRequest("GET", "/alunos/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var AlunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &AlunoMock)

	assert.Equal(t, "Nome do aluno teste", AlunoMock.Nome)
	assert.Equal(t, "12345678901", AlunoMock.CPF)
	assert.Equal(t, "123456789", AlunoMock.RG)

	ApagaAlunoMock()
}
func TestDeletaAluno(t *testing.T) {

	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	req, _ := http.NewRequest("DELETE", "/alunos/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

}
func TestEditaAluno(t *testing.T) {

	database.ConectaComBancoDeDados()
	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)

	aluno := models.Aluno{Nome: "Nome do aluno teste", CPF: "10987654321", RG: "987654321"}
	valorJson, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PATCH", "/alunos/"+strconv.Itoa(ID), bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, "Nome do aluno teste", alunoMockAtualizado.Nome)
	assert.Equal(t, "10987654321", alunoMockAtualizado.CPF)
	assert.Equal(t, "987654321", alunoMockAtualizado.RG)

	defer ApagaAlunoMock()

}
