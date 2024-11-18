package tests

import (
	"bytes"
	"encoding/json"
	"go_test/controller"
	"go_test/dto"
	"go_test/models"
	"go_test/repository"
	"go_test/route"
	"go_test/service"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func sDBT() *gorm.DB{
	DB, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	DB.AutoMigrate(&models.User{},&models.Book{},&models.Author{})

	
	return DB
}

func routerS(db *gorm.DB) *gin.Engine{
	userRepo := repository.UserRepo(db)
	userService := service.UserServices(userRepo)
	userController := controller.UserControllerDI(userService)

	bookRepo := repository.BookRepo(db)
	bookService := service.BookServices(bookRepo)
	bookController := controller.BookControllerDI(bookService)

	authorRepo := repository.AuthorRepo(db)
	authorService := service.AuthorServices(authorRepo)
	authorController := controller.AuthorControllerDI(authorService)
	r := gin.Default()
	
	route.UserRoutes(r,userController)
	route.BookRoutes(r,bookController)
	route.AuthorRoutes(r,authorController)

	return r
}
func login(t *testing.T, r *gin.Engine) string{
	payload := map[string]string{
		"username":"test",
		"password":"test",
	}

	jsonP,_ := json.Marshal(payload)

	w:= httptest.NewRecorder()

	req,_ := http.NewRequest("POST","/auth/register",bytes.NewBuffer(jsonP))

	r.ServeHTTP(w,req)
	assert.Equal(t, http.StatusCreated,w.Code)

	w = httptest.NewRecorder()
	req,_ = http.NewRequest("POST","/auth/login",bytes.NewBuffer(jsonP))

	r.ServeHTTP(w,req)
	assert.Equal(t, http.StatusCreated,w.Code)

	var res map[string]string
	_ = json.Unmarshal(w.Body.Bytes(),&res)

	return res["token"]
 

}

func TestBook(t *testing.T){
	db := sDBT()
	r := routerS(db)
	layout:="2000-01-01"
	dateStr:="2000-01-02"
	dateP ,_:=time.Parse(layout,dateStr)

	author := models.Author{
		Name :"Test",
		Birthdate:dateP ,
	}
	db.Create(&author)

	token:= login(t,r)
	t.Run("CreateBook", func(t *testing.T) {
		book := map[string]interface{}{
			"title":    "Test",
			"isbn":     "testISBN",
			"authorid": author.ID,
		}
		bookJSON, _ := json.Marshal(book)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(bookJSON))
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "Book registered successfully")
	})

	
	t.Run("CreateBookInvalid", func(t *testing.T) {
		book := map[string]interface{}{
			"title":    "Test",
			"isbn":     "testISBN",
			"authorid": author.ID,
		}
		bookJSON, _ := json.Marshal(book)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(bookJSON))
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Failed to register book")
	})

	t.Run("CreateBookInvalid", func(t *testing.T) {
		book := map[string]interface{}{
			"title":    "Test",
			"isbn":     "testISBN",
			"authorid": 99999,
		}
		bookJSON, _ := json.Marshal(book)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(bookJSON))
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Failed to register book")
	})

	t.Run("CreateBookInvalid", func(t *testing.T) {
		book := map[string]interface{}{
			"title":    "Test",
			"isbn":     "testISBN",
			"authorid": 99999,
		}
		bookJSON, _ := json.Marshal(book)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(bookJSON))
		// req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "invalid")
	})


	t.Run("CreateBookInvalid", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books/", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "invalid")
	})
	t.Run("EditBook", func(t *testing.T) {
		book := map[string]interface{}{
			"title":    "Test3",
			"isbn":     "testISBN3",
			"authorid": author.ID,
		}
		bookJSON, _ := json.Marshal(book)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(bookJSON))
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "update")
	})
	


	t.Run("EditBook", func(t *testing.T) {
		book := map[string]interface{}{
			"title":    "Test3",
			"isbn":     "testISBN3",
			"authorid": author.ID,
		}
		bookJSON, _ := json.Marshal(book)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(bookJSON))
		// req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "invalid")
	})
	


	t.Run("GetBook", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books/1",nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Title")
	})
t.Run("GetBook", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books/",nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Title")
	})

	t.Run("GetBook", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books/2000",nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})

	t.Run("GetBook", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books/2000",nil)
		// req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "invalid")
	})

	t.Run("DeleteBook", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/books/1",nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "deleted")
	})
	t.Run("DeleteBookInvalid", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/books/2",nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})

t.Run("DeleteBookInvalid", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/books/2",nil)
		// req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "invalid")
	})

}

func TestAuthor(t *testing.T){
	db := sDBT()
	r := routerS(db)
	

	
	token:= login(t,r)
	t.Run("CreateAuthor", func(t *testing.T) {
		
		author := dto.ARequest{
		Name :"Test",
		Birthdate:"2000-01-02" ,
	}
		authorJSON, _ := json.Marshal(author)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/authors/", bytes.NewBuffer(authorJSON))
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "Author registered successfully")
	})

	
	t.Run("CreateAuthorInvalid", func(t *testing.T) {
		author := dto.ARequest{
		Name :"Test",
		Birthdate:"2000-01-02ksf" ,
	}
		authorJSON, _ := json.Marshal(author)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/authors/", bytes.NewBuffer(authorJSON))
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid")
	})

	t.Run("CreateAuthorInvalid", func(t *testing.T) {
		book := map[string]interface{}{
			"title":    "Test",
			"isbn":     "testISBN",
			"authorid": 99999,
		}
		bookJSON, _ := json.Marshal(book)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(bookJSON))
		// req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "invalid token")
	})

	t.Run("CreateAuthorInvalid", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books/", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "invalid")
	})
	t.Run("EditAuthor", func(t *testing.T) {
		author := dto.ARequest{
		Name :"TestUSD",
		Birthdate:"2000-01-02" ,
	}
		authorJSON, _ := json.Marshal(author)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/authors/1", bytes.NewBuffer(authorJSON))
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "update")
	})

	t.Run("EditAuthor", func(t *testing.T) {
		author := dto.ARequest{
		Name :"TestUSD",
		Birthdate:"2000-01-02" ,
	}
		authorJSON, _ := json.Marshal(author)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/authors/1", bytes.NewBuffer(authorJSON))
		// req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "invalid")
	})
	


	t.Run("GetAuthor", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/authors/1",nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "name")
	})

	t.Run("GetAuthor", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/authors/",nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "name")
	})

	t.Run("GetAuthor", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/authors/",nil)
		// req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "invalid")
	})

	t.Run("GetAuthor", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/authors/2000",nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})

	t.Run("DeleteAuthor", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/authors/1",nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "deleted")
	})
	t.Run("DeleteAuthorInvalid", func(t *testing.T) {
		

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/authors/2",nil)
		req.Header.Set("Authorization", "Bearer "+token)
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})

}

func TestUser(t *testing.T){
	db := sDBT()
	r := routerS(db)

	t.Run("CreateUser", func(t *testing.T) {
		
		user := models.User{
		Username :"Test",
		Password:"2000-01-02" ,
	}
		userJSON, _ := json.Marshal(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(userJSON))
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "successfully")
	})

	
	t.Run("CreateUserInvalid", func(t *testing.T) {
		
		user := models.User{
		Username :"",
		Password:"2000-01-02" ,
	}
		userJSON, _ := json.Marshal(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(userJSON))
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})

	t.Run("CreateUserInvalid", func(t *testing.T) {
		
		user := models.User{
		Username :"fd",
		Password:"" ,
	}
		userJSON, _ := json.Marshal(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(userJSON))
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})
	t.Run("CreateUserInvalid", func(t *testing.T) {
		
		user := models.User{
		Username :"",
		Password:"" ,
	}
		userJSON, _ := json.Marshal(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(userJSON))
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})


	t.Run("loginUser", func(t *testing.T) {
		
		user := models.User{
		Username :"Test",
		Password:"2000-01-02" ,
	}
		userJSON, _ := json.Marshal(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(userJSON))
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "token")
	})

	
	t.Run("LoginUserInvalid", func(t *testing.T) {
		
		user := models.User{
		Username :"",
		Password:"2000-01-02" ,
	}
		userJSON, _ := json.Marshal(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(userJSON))
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})

	t.Run("loginUserInvalid", func(t *testing.T) {
		
		user := models.User{
		Username :"fd",
		Password:"" ,
	}
		userJSON, _ := json.Marshal(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(userJSON))
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})
	t.Run("loginUserInvalid", func(t *testing.T) {
		
		user := models.User{
		Username :"",
		Password:"" ,
	}
		userJSON, _ := json.Marshal(user)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(userJSON))
		r.ServeHTTP(w, req)

		println(w.Body.String())

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})

}