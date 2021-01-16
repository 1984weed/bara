package rest_test

import (
	"bara/auth"
	"bara/mocks"
	"bara/model"
	"bara/problem/domain"
	"bara/problem/rest"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"gotest.tools/v3/assert"
)

// JWT_SECRET ...
const JWT_SECRET = "JWT_SECRET"

func TestCreateProblem(t *testing.T) {
	mockProblemUC := new(mocks.ProblemUsecase)

	api := rest.NewProblemRestApi(mockProblemUC)

	m := chi.NewRouter()
	m.Use(auth.Middleware("JWT_SECRET"))
	m.Post("/", api.CreateProblem)

	ts := httptest.NewServer(m)

	var jsonStr = []byte(`{
		"problem": {
			"title": "Test from rest api",
			"slug": "test-from-rest-api",
			"description": "test descriptionです",
			"functionName": "testFromRestApi",
			"outputType": "string",
			"args": [
				{
					"name": "people",
					"type": "int"
				},
				{
					"name": "name",
					"type": "string"
				}
			],
			"testcases": [{
				"inputs": ["1", "steve"],
				"output": "testtest"
			}]
		}
	}`)

	expectedNewProblem := &domain.Problem{
		Slug:          "test-from-rest-api",
		Title:         "Test from rest api",
		Description:   "It's a test description",
		FunctionName:  "testFromRestApi",
		LanguageSlugs: []model.CodeLanguageSlug{model.JavaScript},
		ProblemArgs: []domain.ProblemArgs{
			{Name: "target", VarType: "int"},
			{Name: "num", VarType: "int[]"},
		},
		ProblemTestcases: []domain.Testcase{
			{
				InputArray: []string{"1"},
				Input:      "input",
				Output:     "output",
			},
		},
		OutputType: "int",
	}

	t.Run("success", func(t *testing.T) {
		mockProblemUC.On("CreateProblem", mock.Anything, mock.Anything, int64(1)).Return(expectedNewProblem, nil).Once()

		if _, body := testRequest(t, ts, "POST", "/", GetAdminJWT(), bytes.NewBuffer(jsonStr)); body != "" {
			expectedJSON := `{"problem":{"title":"Test from rest api","slug":"test-from-rest-api","description":"It's a test description","outputType":"int","functionName":"testFromRestApi","args":[{"name":"target","type":"int"},{"name":"num","type":"int[]"}],"testcases":[{"inputs":["1"],"output":"output"}]}}`
			assert.Equal(t, body, expectedJSON)
		}
	})
	t.Run("access authority is rejected", func(t *testing.T) {
		mockProblemUC.On("CreateProblem", mock.Anything, mock.Anything, int64(3)).Return(expectedNewProblem, nil).Once()

		if res, body := testRequest(t, ts, "POST", "/", GetGeneralUserJWT(), bytes.NewBuffer(jsonStr)); body != "" {
			assert.Equal(t, res.StatusCode, http.StatusUnauthorized)
		}
	})
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, jwt string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	req.Header.Add("Authorization", fmt.Sprintf("Bear %s", jwt))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, strings.Trim(string(respBody), "\t \n")
}

func GetAdminJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "1",
		"role": "admin",
	})

	// It never fails
	tokenString, _ := token.SignedString([]byte(JWT_SECRET))

	return tokenString
}

func GetGeneralUserJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "3",
		"role": "general",
	})

	// It never fails
	tokenString, _ := token.SignedString([]byte(JWT_SECRET))

	return tokenString
}
