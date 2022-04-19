package main

//import the packages needed for the program
import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// article represents data about a record article
type article struct {
	Id    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

// articles slice to seed record article data
var articles = []article{
	{
		Id:    "1",
		Title: "latest science shows that potato chips are better for you than sugar",
		Date:  "2016-09-22",
		Body:  "some text, potentially containing simple markup about how potato chips are great",
		Tags:  []string{"health", "fitness", "science"},
	},
}

// Tag response struct
type TagRes struct {
	Tag          string   `json:"tag"`
	Count        int      `json:"count"`
	Articles     []string `json: "articles"`
	Related_Tags []string `json:"related_tags"`
}

// This functions outputs json response from the slice of article structs
func getArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, articles)
}

//this function create new article from JSON received in the request body
func createArticle(c *gin.Context) {
	var newArticle article
	// BindJSON called to bind the received JSON to newArticle
	if err := c.BindJSON(&newArticle); err != nil {
		return
	}
	//add the new article to the slice
	articles = append(articles, newArticle)
	c.IndentedJSON(http.StatusCreated, newArticle)
}

//this function find article by the id
func ArticleById(c *gin.Context) {
	id := c.Param("id")
	book, err := getArticleById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// this function loop over the list of article if matches id it return the article
func getArticleById(id string) (*article, error) {
	for i, a := range articles {
		if a.Id == id {
			return &articles[i], nil
		}
	}

	return nil, errors.New("Article not found")
}

//this function  return the list of articles that have that tag name on the
//given date and some summary data about that tag for that day.
func getTags(c *gin.Context) {
	tag := c.Param("tagName")
	date := c.Param("date")
	//loop over the list of article to match the year
	tempCount := 0
	tempId := []string{}
	tempTags := []string{}
	for i, a := range articles {
		aDate := strings.Join(strings.Split(a.Date, "-"), "")
		log.Print("date: ", aDate+tag)

		if aDate == date {
			//if date matches in article
			log.Print(a.Id, a.Date)
			for _, t := range articles[i].Tags {
				if tag == t {
					//if date and tag name matches in article
					tempCount += 1
					tempId = append(tempId, a.Id)
					tempTags = append(tempTags, a.Tags...)
					log.Print(t, tempTags)
				}
			}

		}

	}
	if tempCount > 0 {
		log.Print(tempTags)
		uniqueTags := unique(tempTags, tag)
		log.Print(uniqueTags)
		tagRes := TagRes{Tag: tag, Count: tempCount, Articles: tempId, Related_Tags: uniqueTags}
		log.Print(tagRes)
		log.Print(tagRes)
		c.IndentedJSON(http.StatusOK, tagRes)
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found."})
	}
}

// this function removes the duplicate tags and return the unique relative tags
func unique(arr []string, tagName string) []string {
	occurred := map[string]bool{}
	result := []string{}
	for e := range arr {
		// if tag name is present in relative tags don't append.
		if arr[e] == tagName {
			log.Print("This is the tagName ", tagName+" . Do not add to result.")
		} else {
			// check if already  mapped variable is set to true or not
			if occurred[arr[e]] != true {
				occurred[arr[e]] = true

				// Append to result slice.
				result = append(result, arr[e])
			}
		}
	}

	return result
}

func main() {
	//setup the gin router
	router := gin.Default()                     //Initialize a Gin router using Default.
	router.GET("/articles", getArticles)        //list all the available articles
	router.GET("/articles/:id", ArticleById)    //show article if matches by article id
	router.POST("/articles", createArticle)     // create new article
	router.GET("/tags/:tagName/:date", getTags) //gets tag summary if date and tagname matches in article
	router.Run("localhost:8080")
}
