package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
  "go-rest/db"
	docs "go-rest/docs"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type (
	RestUser struct {
		//		gorm.Model
		//ID     string `gorm:"type:uuid;default:gen_random_uuid();rimaryKey" json:"id"`
		ID   int    `gorm:"type:int;autoincrement;rimaryKey" json:"id"`
		Name string `gorm:"type:varchar(32)" json:"name"`
	}
)

var (
	lock = sync.Mutex{}
)

//----------
// Handlers
//----------

type API struct {
	e  *echo.Echo
	db *gorm.DB
}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	fmt.Printf("GOID:%v\n", id)
	return id
}

func NewRest(e *echo.Echo, da *db.Adapter) {

  db:=da.Get()
	db.AutoMigrate(&RestUser{})
	api := &API{e: e, db: db}
	e.PUT("/users/:id", api.updateUser)
	e.GET("/users/:id", api.getUser)
	e.GET("/users", api.getAllUsers)
	e.POST("/users", api.createUser)
	e.DELETE("/users/:id", api.deleteUser)
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger User API"
	docs.SwaggerInfo.Description = "This is a sample rest server. v1.1"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	fmt.Printf("docs.SwaggerInfo: %v\n", docs.SwaggerInfo.Description)
}

// Create User
// @Summary Create User.
// @Description Create User.
// @Tags users
// @ParametersIn 3333
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [post]
func (api API) createUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	user := RestUser{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := api.db.Create(&user).Error; err != nil {
		log.Printf("Create User Error:%v", err)
		return c.JSON(http.StatusBadRequest,err)
	}
	return c.JSON(http.StatusCreated, user)
}

// Get User By ID godoc
// @Summary Get User By ID.
// @Description Get User Info By ID.
// @Tags users
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/:id [get]
func (api API) getUser(c echo.Context) error {
	user := RestUser{}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := api.db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Printf("Get User Error:%v", err)
		return c.JSON(http.StatusNotFound,err)
	}
	return c.JSON(http.StatusOK, user)
}

// Update User Info By ID godoc
// @Summary Update User By ID.
// @Description Update User info By ID
// @Tags users
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/:id [put]
func (api API) updateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	user := RestUser{}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := api.db.Where("id = ?", id).Take(&user).Error; err != nil {
		log.Printf("Update User Error:%v", err)
		return c.JSON(http.StatusNotFound,err)
	}
	if user.ID == id {
		if err := c.Bind(&user); err != nil {
			return err
		}
		api.db.Save(&user)
	}
	//api.db.Model(&RestUser{}).Where("id = ?",id).Update("name", user.Name)
	return c.JSON(http.StatusOK, user)
}

// Delete User By ID godoc
// @Summary Delete User By ID.
// @Description et the status of server.
// @Tags users
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [delete]
func (api API) deleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	user := RestUser{}
	if err := api.db.Where("id = ?", id).Unscoped().Delete(&user).Error; err != nil {
		log.Printf("Delete User Error:%v", err)
		return c.JSON(http.StatusNotFound,err)
	}
	return c.JSON(http.StatusOK, user)
}

// Get All User
// @Summary Get All User.
// @Description Get All Users from server.
// @Tags users
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
func (api API) getAllUsers(c echo.Context) error {
	users := []RestUser{}
	if err := api.db.Find(&users).Error; err != nil {
		log.Printf("Get User Error:%v", err)
		return c.JSON(http.StatusNotFound,err)
	}
	return c.JSON(http.StatusOK, users)
}
