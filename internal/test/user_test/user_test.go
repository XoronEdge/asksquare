package user_test_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/XoronEdge/asksquare/graph/generated"
	graph "github.com/XoronEdge/asksquare/graph/resolvers"
	"github.com/XoronEdge/asksquare/initial"
	userRepo "github.com/XoronEdge/asksquare/internal/user/repo/postgres"
	userUsecase "github.com/XoronEdge/asksquare/internal/user/usecase"
	"github.com/labstack/echo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var cleaner func()
var _ = Describe("User", func() {
	BeforeEach(func() {
		cleaner = initial.CleanUpHook(db)
	})
	JustAfterEach(func() {
		defer cleaner()
	})
	Context("Create", func() {
		It("Should create user", func() {
			user := `{"query":"mutation createUser {\n  createUser(\n    input: {\n      id: 1\n      firstname: \"Joe\"\n      lastname: \"Doe\"\n      username: \"joedoe\"\n      phone: \"03204181540\"\n      password: \"password\"\n      email: \"joe@gmail.com\"\n      gender: \"male\"\n    }\n  ){\n      firstname\n      lastname\n      username\n      phone\n      password\n      email\n      gender\n  }\n}\n"}`
			var res_map map[string]interface{}
			e := echo.New()
			ur := userRepo.NewUserRepo(db)
			uc := userUsecase.NewUserUsecase(ur, time.Minute)
			graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Uc: uc}}))
			playgroundHandler := playground.Handler("GraphQL", "/query")
			e.POST("/query", func(c echo.Context) error {
				graphqlHandler.ServeHTTP(c.Response(), c.Request())
				return nil
			})
			e.GET("/playground", func(c echo.Context) error {
				playgroundHandler.ServeHTTP(c.Response(), c.Request())
				return nil
			})
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(user))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			res := httptest.NewRecorder()
			graphqlHandler.ServeHTTP(res, req)
			json.Unmarshal(res.Body.Bytes(), &res_map)
			Expect(res.Code).To(Equal(200))
		})
	})
})
