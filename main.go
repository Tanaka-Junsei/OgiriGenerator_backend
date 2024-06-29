package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	oapi "github.com/Tanaka-Junsei/OgiriGenerator_backend/generated"
	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// apiControllerは、APIエンドポイントのハンドラーを実装します。
type apiController struct{}

// OpenAPI で定義された (GET /users) の実装
func (a apiController) GetUser(ctx echo.Context) error {
	// OpenAPI で生成された User モデルを使ってレスポンスを返す
	return ctx.JSON(http.StatusOK, &oapi.User{
		UserId: "1",
	})
}

// OpenAPI で定義された (POST /users) の実装
func (a apiController) PostUser(ctx echo.Context) error {
	// リクエストボディを構造体にバインド
	user := new(oapi.User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	fmt.Println(user)
	// 200 ステータスのみ返す
	return ctx.NoContent(http.StatusOK)
}

// OpenAPI で定義された (POST /answers) の実装
func (a apiController) PostAnswer(ctx echo.Context) error {
	// リクエストボディを構造体にバインド
	answer := new(oapi.Answer)
	if err := ctx.Bind(answer); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	fmt.Println(answer)
	// 200 ステータスのみ返す
	return ctx.NoContent(http.StatusOK)
}

// OpenAPI で定義された (GET /answers/byAnswerId) の実装
func (a apiController) GetAnswerByAnswerId(ctx echo.Context, params oapi.GetAnswerByAnswerIdParams) error {
	// stringをintに変換
	answerId, err := strconv.Atoi(params.AnswerId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid answerId"})
	}
	// 仮のデータを返す
	return ctx.JSON(http.StatusOK, &oapi.Answer{
		AnswerId:   answerId,
		Answer:     "This is a sample answer",
		Question:   "What is the question?",
		UserId:     "1",
		QuestionId: 1,
	})
}

// OpenAPI で定義された (GET /answers/byUserId) の実装
func (a apiController) GetAnswersByUserId(ctx echo.Context, params oapi.GetAnswersByUserIdParams) error {
	// 仮のデータを返す
	answers := []oapi.Answer{
		{
			AnswerId:   1,
			Answer:     "Answer 1",
			Question:   "Question 1",
			UserId:     params.UserId,
			QuestionId: 1,
		},
		{
			AnswerId:   2,
			Answer:     "Answer 2",
			Question:   "Question 2",
			UserId:     params.UserId,
			QuestionId: 2,
		},
	}
	return ctx.JSON(http.StatusOK, answers)
}

// OpenAPI で定義された (POST /questions/generate) の実装
func (a apiController) GenerateQuestion(ctx echo.Context) error {
	// PythonのAPIにリクエストを送信
	resp, err := http.Post("http://127.0.0.1:8000/add-data", "application/json", nil)
	if err != nil {
		log.Print(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to contact Python API"})
	}
	defer resp.Body.Close()

	// レスポンスを読み込む
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read response from Python API"})
	}

	// レスポンスのデータを解析
	var generatedQuestion oapi.Question
	if err := json.Unmarshal(body, &generatedQuestion); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse response from Python API"})
	}

	// クライアントに返す
	return ctx.JSON(http.StatusOK, generatedQuestion)
}

// OpenAPI で定義された (GET /questions) の実装
func (a apiController) GetQuestion(ctx echo.Context) error {
	// 仮のデータを返す
	return ctx.JSON(http.StatusOK, &oapi.Question{
		QuestionId: 1,
		Question:   "This is a sample question",
	})
}

func main() {
	e := echo.New()
	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}
	e.Use(oapiMiddleware.OapiRequestValidator(swagger))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := apiController{}
	oapi.RegisterHandlers(e, api)

	e.Logger.Fatal(e.Start(":8080"))
}
