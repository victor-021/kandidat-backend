package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var (
	dbPool      *pgxpool.Pool
	serverURL   string
	databaseURL string
)

type Community struct {
	CommunityID int
	Name        string
}

type User struct {
	UserID      int
	Name        string
	PhoneNumber int
	Address     string
}

type Review struct {
	ReviewID  int
	UserID    int
	ProductID int
	Rating    int
	Content   string
}

type Product struct {
	ProductID   int
	Name        string
	Service     bool
	Price       int
	UploadDate  string
	Description string
	UserID      int
}

func setupConfig() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("POSTGRES_DB")
	databaseUser := os.Getenv("POSTGRES_USER")
	databasePassword := os.Getenv("POSTGRES_PASSWORD")

	// Change empty config values to default values
	if serverHost == "" {
		serverHost = "localhost"
	}

	if serverPort == "" {
		serverPort = "8080"
	}

	if databaseHost == "" {
		databaseHost = "localhost"
	}

	if databasePort == "" {
		databasePort = "5432"
	}

	if databaseName == "" {
		databaseName = "backend-db"
	}

	if databaseUser == "" {
		databaseUser = "dbuser"
	}

	if databasePassword == "" {
		databasePassword = "kandidat-backend"
	}

	serverURL = serverHost + ":" + serverPort
	databaseURL = "postgres://" + databaseUser + ":" + databasePassword + "@" + databaseHost + ":" + databasePort + "/" + databaseName
}

func setupDBPool() *pgxpool.Pool {
	dbpool, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return dbpool
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func setupRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()
	// Log to stdout.
	gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.GET("/ping", ping)
	router.GET("/communities", getCommunities)
	router.GET("/communityname", getCommunityName)
	router.GET("/user/:userid/communities", getUsersCommunities)
	router.GET("/user/:userid", getUser)
	router.GET("/products/:productid", getProductID)
	router.POST("/users", createUser)
	return router
}

func testDB() {
	var greeting string
	err := dbPool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}

func main() {
	setupConfig()

	dbPool = setupDBPool()
	defer dbPool.Close()

	testDB()

	router := setupRouter()
	err := router.Run(serverURL)

	if err != nil {
		fmt.Println(err)
	}
}

func getCommunities(c *gin.Context) {
	query := "SELECT * FROM Community"
	rows, err := dbPool.Query(c, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var communities []Community
	for rows.Next() {
		var community Community
		err := rows.Scan(&community.CommunityID, &community.Name)
		if err != nil {
			panic(err)
		}
		communities = append(communities, community)
	}

	c.JSON(http.StatusOK, communities)
}

func getUsersCommunities(c *gin.Context) {
	user := c.Param("userid")
	query := "SELECT * from Community WHERE community_id = (SELECT fk_community_id FROM User_Community WHERE fk_user_id = $1)"
	rows, err := dbPool.Query(c, query, user)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var communities []Community
	for rows.Next() {
		var community Community
		err := rows.Scan(&community.CommunityID, &community.Name)
		if err != nil {
			panic(err)
		}
		communities = append(communities, community)
	}

	c.JSON(http.StatusOK, communities)
}

func getUser(c *gin.Context) {
	var result User
	user := c.Param("userid")
	query := "SELECT * from Users WHERE user_id = $1"
	err := dbPool.QueryRow(c, query, user).Scan(&result.UserID, &result.Name, &result.PhoneNumber, &result.Address)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, result)
}

func getProductID(c *gin.Context) {
	var result Product
	productId := c.Param("productid")
	query := "SELECT * FROM Product WHERE product_id = $1"
	err := dbPool.QueryRow(c, query, productId).Scan(&result)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	c.JSON(http.StatusOK, result)
}

//Useless?
func getNewCommunities(c *gin.Context) {
	user_id := 3 // TEST
	var result int
	query := "SELECT fk_community_id FROM User_Community WHERE fk_user_id != $1"
	err := dbPool.QueryRow(c, query, user_id).Scan(&result)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	c.JSON(http.StatusOK, result)
}

func getCommunityName(c *gin.Context) {
	community_id := 1 //TEST
	var result string
	query := "SELECT name FROM Community WHERE community_id = $1"
	err := dbPool.QueryRow(c, query, community_id).Scan(&result)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	c.JSON(http.StatusOK, result)
}

func createUser(c *gin.Context) {
	name := c.PostForm("n")
	address := c.PostForm("a")
	phone_nr, _ := strconv.Atoi(c.PostForm("p"))

	query := "INSERT INTO Users(name, phone_nr, address) VALUES($1,$2, $3)"
	_, err := dbPool.Exec(c, query, name, phone_nr, address)

	if err != nil {
		log.Fatal(err)
	}
	c.Status(http.StatusOK)
}
