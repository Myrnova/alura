package main

import (
	"api-alunos-gin/src/controllers"
	"api-alunos-gin/src/database"
	"api-alunos-gin/src/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestVerificarStatusCodeDaSaudacaoComParametroHandler(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/teste", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	resBody, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.Code, "Deveriam ser iguais")
	assert.Equal(t, `{"API diz:":"E ai teste, tudo beleza?"}`, string(resBody))
}

func BuildAluno(alunoParametro models.Aluno) models.Aluno {
	if alunoParametro.Nome == "" {
		alunoParametro.Nome = "Nome do Aluno Teste"
	}
	if alunoParametro.CPF == "" {
		alunoParametro.CPF = "12345678901"
	}
	if alunoParametro.RG == "" {
		alunoParametro.RG = "123456789"
	}
	return alunoParametro
}

func CriaAlunoMock(aluno *models.Aluno) { //I am creating a new aluno using a already existing struct instance so if I want to have the updated informations I need to pass the pointer to the memory reference
	database.DB.Create(&aluno)
}

func TruncateAlunos() {
	database.DB.Exec("TRUNCATE alunos RESTART IDENTITY;")
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	primeiroAluno := BuildAluno(models.Aluno{})
	segundoAluno := BuildAluno(models.Aluno{Nome: "Segundo aluno"})
	database.ConectaComBancoDeDados()

	CriaAlunoMock(&primeiroAluno)
	CriaAlunoMock(&segundoAluno)

	defer TruncateAlunos()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var resultAlunos []models.Aluno
	if err := json.NewDecoder(resp.Body).Decode(&resultAlunos); err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Aluno, '%v'", resp.Body, err)
	}

	expectedAlunos := []models.Aluno{
		primeiroAluno,
		segundoAluno,
	}

	assert.Equal(t, http.StatusOK, resp.Code, "Deveriam ser iguais")
	assert.Equal(t, expectedAlunos[0].ID, resultAlunos[0].ID)
	assert.Equal(t, expectedAlunos[1].ID, resultAlunos[1].ID)

}

func TestBuscarAlunoPorCPFHandler(t *testing.T) {
	alunoBuscado := BuildAluno(models.Aluno{Nome: "Segundo aluno", CPF: "12345678908"})
	alunoSobrante := BuildAluno(models.Aluno{})
	database.ConectaComBancoDeDados()

	CriaAlunoMock(&alunoBuscado)
	CriaAlunoMock(&alunoSobrante)

	defer TruncateAlunos()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678908", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var resultAluno models.Aluno
	if err := json.NewDecoder(resp.Body).Decode(&resultAluno); err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Aluno, '%v'", resp.Body, err)
	}

	expectedAlunos := []models.Aluno{
		alunoBuscado,
		alunoSobrante,
	}
	assert.Equal(t, http.StatusOK, resp.Code, "Deveriam ser iguais")

	assert.Equal(t, expectedAlunos[0].ID, resultAluno.ID)
	assert.Equal(t, expectedAlunos[0].ID, resultAluno.ID)
	assert.Equal(t, expectedAlunos[0].CPF, resultAluno.CPF)
	assert.Equal(t, expectedAlunos[0].Nome, resultAluno.Nome)

	assert.NotEqual(t, expectedAlunos[1].ID, resultAluno.ID, "Função deve apenas trazer o alunoBuscado cadastrado")

}

func TestBuscarAlunoPorIDHandler(t *testing.T) {
	alunoBuscado := BuildAluno(models.Aluno{})
	alunoSobrante := BuildAluno(models.Aluno{Nome: "Segundo aluno", CPF: "12345678908"})
	database.ConectaComBancoDeDados()

	CriaAlunoMock(&alunoBuscado)
	CriaAlunoMock(&alunoSobrante)

	defer TruncateAlunos()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/alunos/%d", alunoBuscado.ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var resultAluno models.Aluno

	if err := json.Unmarshal(resp.Body.Bytes(), &resultAluno); err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Aluno, '%v'", resp.Body, err)
	}

	expectedAlunos := []models.Aluno{
		alunoBuscado,
		alunoSobrante,
	}

	assert.Equal(t, http.StatusOK, resp.Code, "Deveriam ser iguais")

	assert.Equal(t, expectedAlunos[0].ID, resultAluno.ID)
	assert.Equal(t, expectedAlunos[0].CPF, resultAluno.CPF)
	assert.Equal(t, expectedAlunos[0].Nome, resultAluno.Nome)

	assert.NotEqual(t, expectedAlunos[1].ID, resultAluno.ID, "Função deve apenas trazer o alunoBuscado cadastrado")

}

func TestDeletarAlunoHandler(t *testing.T) {
	alunoDeletado := BuildAluno(models.Aluno{})
	alunoSobrante := BuildAluno(models.Aluno{Nome: "Segundo aluno", CPF: "12345678908"})
	database.ConectaComBancoDeDados()

	CriaAlunoMock(&alunoDeletado)
	CriaAlunoMock(&alunoSobrante)

	defer TruncateAlunos()

	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/alunos/%d", alunoDeletado.ID), nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	resBody, _ := ioutil.ReadAll(resp.Body)

	var alunos []models.Aluno
	database.DB.Find(&alunos)

	assert.Equal(t, http.StatusOK, resp.Code, "Deveriam ser iguais")

	assert.Equal(t, `{"data":"Aluno deletado com sucesso"}`, string(resBody))
	assert.Lenf(t, alunos, 1, "Deve trazer apenas um aluno depois da execução de deletar")
	assert.Equal(t, alunoSobrante.ID, alunos[0].ID, "Deve trazer apenas o alunoSobrante depois da execução de deletar")
}

func TestEditarAlunoHandler(t *testing.T) {
	alunoNovo := BuildAluno(models.Aluno{})
	database.ConectaComBancoDeDados()

	CriaAlunoMock(&alunoNovo)

	defer TruncateAlunos()

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)

	alunoEditado := BuildAluno(models.Aluno{Nome: "Segundo aluno", CPF: "12345678908"})
	alunoEditadoJson, _ := json.Marshal(alunoEditado)

	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/alunos/%d", alunoNovo.ID), bytes.NewBuffer(alunoEditadoJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var resultAluno models.Aluno
	if err := json.Unmarshal(resp.Body.Bytes(), &resultAluno); err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Aluno, '%v'", resp.Body, err)
	}

	assert.Equal(t, http.StatusOK, resp.Code, "Deveriam ser iguais")

	assert.Equal(t, alunoEditado.Nome, resultAluno.Nome)
	assert.Equal(t, alunoEditado.CPF, resultAluno.CPF)

}
