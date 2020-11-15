package user_test_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/XoronEdge/asksquare/domain"
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
			user := `{"query":"mutation createUser {\n  createUser(\n    input: {\n      id: 1\n      firstname: \"Joeee\"\n      lastname: \"Doeee\"\n      username: \"joedoe\"\n      phone: \"03204181540\"\n      password: \"password\"\n      email: \"joe@gmail.com\"\n      gender: \"Male\"\n    }\n  ){\n      firstname\n      lastname\n      username\n      phone\n      password\n      email\n      gender\n  }\n}\n"}`
			var resMap map[string]interface{}
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
			json.Unmarshal(res.Body.Bytes(), &resMap)
			Expect(res.Code).To(Equal(200))
		})
	})
	Context("Get user by ID", func() {
		It("Should get user by id", func() {
			userBufferString := `{"query":"mutation createUser {\n  createUser(\n    input: {\n      id: 1\n      firstname: \"Joeee\"\n      lastname: \"Doeee\"\n      username: \"joedoe\"\n      phone: \"03204181540\"\n      password: \"password\"\n      email: \"joe@gmail.com\"\n      gender: \"Male\"\n    }\n  ){\n      firstname\n      lastname\n      username\n      phone\n      password\n      email\n      gender\n  }\n}\n"}`
			var resMap map[string]interface{}
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
			reqOne := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userBufferString))
			reqOne.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			resOne := httptest.NewRecorder()
			graphqlHandler.ServeHTTP(resOne, reqOne)
			json.Unmarshal(resOne.Body.Bytes(), &resMap)
			var createUserMap map[string]interface{}
			for key, value := range resMap {
				if key == "data" {
					dataMap, _ := value.(map[string]interface{})
					for _, v := range dataMap {
						createUserMap, _ = v.(map[string]interface{})
					}
				}
			}
			var user domain.User
			db.First(&user)
			id := user.ID
			db.Where("id = ?", id).First(&user)
			Expect(createUserMap["email"]).To(Equal(user.Email))
			Expect(createUserMap["firstname"]).To(Equal(user.FirstName))
			Expect(createUserMap["lastname"]).To(Equal(user.LastName))
			Expect(createUserMap["gender"]).To(Equal(user.Gender))
			Expect(createUserMap["phone"]).To(Equal(user.Phone))
			Expect(createUserMap["username"]).To(Equal(user.Username))
		})
	})
	Context("Get all users", func() {
		It("Should get all users", func() {
			userBufferString := `{"query":"mutation createUser {\n  createUser(\n    input: {\n      id: 1\n      firstname: \"Joeee\"\n      lastname: \"Doeee\"\n      username: \"joedoe\"\n      phone: \"03204181540\"\n      password: \"password\"\n      email: \"joe@gmail.com\"\n      gender: \"Male\"\n    }\n  ){\n      firstname\n      lastname\n      username\n      phone\n      password\n      email\n      gender\n  }\n}\n"}`
			var resMap map[string]interface{}
			var expMap map[string]interface{}
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
			reqOne := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userBufferString))
			reqOne.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			resOne := httptest.NewRecorder()
			graphqlHandler.ServeHTTP(resOne, reqOne)
			json.Unmarshal(resOne.Body.Bytes(), &resMap)
			getAllUsers := `{"query":"query\n{\n  users{\n    firstname\n    lastname\n    username\n    gender\n    email\n    gender\n    phone\n    password\n  }\n}"}`
			reqTwo := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(getAllUsers))
			reqTwo.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			resTwo := httptest.NewRecorder()
			graphqlHandler.ServeHTTP(resTwo, reqTwo)
			json.Unmarshal(resOne.Body.Bytes(), &expMap)
			Expect(resMap).To(Equal(expMap))
		})
	})
})
