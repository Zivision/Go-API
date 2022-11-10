package main

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
)

type account struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

// Fake DB with names and emails

var accounts = []account {
	{Name: "Vergil", Email: "Vergil@Motivation.com"},
	{Name: "Tavish", Email: "Tavish@RED.com"},
	{Name: "Jane", Email: "Jane@BLU.com"},
}

func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil);
}

func aboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil);
}

func getAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, accounts);
}

func getAccountByName(name string) (*account, error) {
	for i, acc := range accounts {
		if acc.Name == name {
			return &accounts[i], nil;
		}
	}

	return nil, errors.New("Account not Found");
}

func accountByName(c *gin.Context) {
	name := c.Param("name")
	account, err := getAccountByName(name);	

	if err != nil {
		c.String(http.StatusNotFound, "Account Not Found");
		return 
	}

	c.IndentedJSON(http.StatusOK, account)
}

func addAccount(c *gin.Context) {
	var newAccount account;

	if err := c.BindJSON(&newAccount); err != nil { return }

	accounts = append(accounts, newAccount);
	c.IndentedJSON(http.StatusCreated, newAccount);
}

func main() {
	router := gin.Default();
	
	router.LoadHTMLGlob("./templates/*.html");
	router.Static("/static/", "./static/");
	
	router.GET("/", indexPage);
	router.GET("/about", aboutPage);

	router.GET("/accounts", getAccounts);
	router.GET("/accounts/:name", accountByName)

	router.POST("/accounts", addAccount);

	router.Run("localhost:8080");
}
