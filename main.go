package main

import (
	"fmt"
	"net/http"
	"time"

	oapi "github.com/Tanaka-Junsei/OgiriGenerator_backend/generated"
	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// apiControllerは、APIエンドポイントのハンドラーを実装します。
type apiController struct{}

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

// OpenAPI で定義された (GET /answers/byLike) の実装
func (a apiController) GetAnswersByLike(ctx echo.Context, params oapi.GetAnswersByLikeParams) error {
	// 仮のデータを返す
	answers := []oapi.Answer{
		{
			AnswerId:   "1",
			Answer:     "Answer 1",
			Question:   "Question 1",
			UserId:     "1",
			QuestionId: "1",
			Timestamp:  time.Now(),
			IsPublic:   true,
			Likes:      10,
		},
		{
			AnswerId:   "2",
			Answer:     "Answer 2",
			Question:   "Question 2",
			UserId:     "2",
			QuestionId: "2",
			Timestamp:  time.Now(),
			IsPublic:   false,
			Likes:      5,
		},
	}
	return ctx.JSON(http.StatusOK, answers)
}

// OpenAPI で定義された (GET /answers/byQuestionId) の実装
func (a apiController) GetAnswersByQuestionId(ctx echo.Context, params oapi.GetAnswersByQuestionIdParams) error {
	// 仮のデータを返す
	answers := []oapi.Answer{
		{
			AnswerId:   "1",
			Answer:     "Answer 1",
			Question:   "Question 1",
			UserId:     "1",
			QuestionId: params.QuestionId,
			Timestamp:  time.Now(),
			IsPublic:   true,
			Likes:      10,
		},
		{
			AnswerId:   "2",
			Answer:     "Answer 2",
			Question:   "Question 2",
			UserId:     "2",
			QuestionId: params.QuestionId,
			Timestamp:  time.Now(),
			IsPublic:   false,
			Likes:      5,
		},
	}
	return ctx.JSON(http.StatusOK, answers)
}

// OpenAPI で定義された (GET /answers/byUserId) の実装
func (a apiController) GetAnswersByUserId(ctx echo.Context, params oapi.GetAnswersByUserIdParams) error {
	// 仮のデータを返す
	answers := []oapi.Answer{
		{
			AnswerId:   "1",
			Answer:     "Answer 1",
			Question:   "Question 1",
			UserId:     params.UserId,
			QuestionId: "1",
			Timestamp:  time.Now(),
			IsPublic:   true,
			Likes:      10,
		},
		{
			AnswerId:   "2",
			Answer:     "Answer 2",
			Question:   "Question 2",
			UserId:     params.UserId,
			QuestionId: "2",
			Timestamp:  time.Now(),
			IsPublic:   false,
			Likes:      5,
		},
	}
	return ctx.JSON(http.StatusOK, answers)
}

// OpenAPI で定義された (POST /likes) の実装
func (a apiController) PostLike(ctx echo.Context) error {
	// リクエストボディを構造体にバインド
	like := new(oapi.Like)
	if err := ctx.Bind(like); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	fmt.Println(like)
	// 200 ステータスのみ返す
	return ctx.NoContent(http.StatusOK)
}

// OpenAPI で定義された (GET /likes/byUserId) の実装
func (a apiController) GetLikesByUserId(ctx echo.Context, params oapi.GetLikesByUserIdParams) error {
	// 仮のデータを返す
	likes := []oapi.Like{
		{
			LikeId:     "1",
			AnswerId:   "1",
			QuestionId: "1",
			UserId:     params.UserId,
			Timestamp:  time.Now(),
		},
		{
			LikeId:     "2",
			AnswerId:   "2",
			QuestionId: "2",
			UserId:     params.UserId,
			Timestamp:  time.Now(),
		},
	}
	return ctx.JSON(http.StatusOK, likes)
}

// OpenAPI で定義された (POST /questions) の実装
func (a apiController) PostQuestion(ctx echo.Context) error {
	// リクエストボディを構造体にバインド
	question := new(oapi.Question)
	if err := ctx.Bind(question); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
	fmt.Println(question)
	// 200 ステータスのみ返す
	return ctx.NoContent(http.StatusOK)
}

// OpenAPI で定義された (GET /questions/byQuestionId) の実装
func (a apiController) GetQuestionsByQuestionId(ctx echo.Context, params oapi.GetQuestionsByQuestionIdParams) error {
	// 仮のデータを返す
	question := oapi.Question{
		QuestionId:  params.QuestionId,
		Question:    "This is a sample question",
		UserId:      "1",
		Timestamp:   time.Now(),
		CreatedByAI: false,
	}
	return ctx.JSON(http.StatusOK, question)
}

// OpenAPI で定義された (GET /questions/byUserId) の実装
func (a apiController) GetQuestionsByUserId(ctx echo.Context, params oapi.GetQuestionsByUserIdParams) error {
	// 仮のデータを返す
	questions := []oapi.Question{
		{
			QuestionId:  "1",
			Question:    "Question 1",
			UserId:      params.UserId,
			Timestamp:   time.Now(),
			CreatedByAI: false,
		},
		{
			QuestionId:  "2",
			Question:    "Question 2",
			UserId:      params.UserId,
			Timestamp:   time.Now(),
			CreatedByAI: true,
		},
	}
	return ctx.JSON(http.StatusOK, questions)
}

// OpenAPI で定義された (GET /users) の実装
func (a apiController) GetUser(ctx echo.Context) error {
	// OpenAPI で生成された User モデルを使ってレスポンスを返す
	return ctx.JSON(http.StatusOK, &oapi.User{
		UserId: "1",
		Name:   "Sample User",
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
