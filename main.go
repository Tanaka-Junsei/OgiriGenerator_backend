package main

import (
	"fmt"
	"net/http"

	oapi "github.com/Tanaka-Junsei/OgiriGenerator_backend/generated"

	oapiMiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type apiController struct{}

// OpenAPI で定義された (GET /users) の実装
func (a apiController) GetUser(ctx echo.Context) error {
	// OpenApi で生成された User モデルを使ってレスポンスを返す
	return ctx.JSON(http.StatusOK, &oapi.User{
		UserId: 1,
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
	// 仮のデータを返す
	return ctx.JSON(http.StatusOK, &oapi.Answer{
		AnswerId: params.AnswerId,
		Answer:   "This is a sample answer",
		Question: "What is the question?",
		UserId:   1,
	})
}

// OpenAPI で定義された (GET /answers/byUserId) の実装
func (a apiController) GetAnswersByUserId(ctx echo.Context, params oapi.GetAnswersByUserIdParams) error {
	// 仮のデータを返す
	answers := []oapi.Answer{
		{
			AnswerId: 1,
			Answer:   "Answer 1",
			Question: "Question 1",
			UserId:   params.UserId,
		},
		{
			AnswerId: 2,
			Answer:   "Answer 2",
			Question: "Question 2",
			UserId:   params.UserId,
		},
	}
	return ctx.JSON(http.StatusOK, answers)
}

func main() {
	// Echo のインスタンス作成
	e := echo.New()

	// OpenApi 仕様に沿ったリクエストかバリデーションをするミドルウェアを設定
	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}
	e.Use(oapiMiddleware.OapiRequestValidator(swagger))
	// ロガーのミドルウェアを設定
	e.Use(middleware.Logger())
	// APIがエラーで落ちてもリカバーするミドルウェアを設定
	e.Use(middleware.Recover())

	// OpenAPI の仕様を満たす構造体をハンドラーとして登録する
	api := apiController{}
	oapi.RegisterHandlers(e, api)

	// 8080ポートで Echo サーバ起動
	e.Logger.Fatal(e.Start(":8080"))
}
