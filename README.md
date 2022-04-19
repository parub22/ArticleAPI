# ArticleAPI


Steps to run the code:
1. Install the latest version of go
2. Add Gin module as a dependency 
  $ go get github.com/gin-gonic/gin 
3. From the command line in the directory containing main.go, run the code
   $ go run .
4. From a new command line window, use curl to make a request to your running web service.

      i. Get the articles
      $ curl http://localhost:8080/articles
      The command should display the data you seeded the service with.
	  ```json
      [
          {
              "id": "1",
              "title": "latest science shows that potato chips are better for you than sugar",
              "date": "2016-09-22",
              "body": "some text, potentially containing simple markup about how potato chips are great",
              "tags": [
                  "health",
                  "fitness",
                  "science"
              ]
          }
      ] 
	  ```

      ii. Create new article
      For the post request, we have created body.json file which has content to be added in articles. While sending the post request we need to include the request          body as shown below.

      $ curl localhost:8080/articles --include --header "Content-Type: application/json" -d @body.json --request "POST"

      The command should display headers and JSON for the added article.

      HTTP/1.1 201 Created
      Content-Type: application/json; charset=utf-8
      Date: Tue, 19 Apr 2022 04:12:22 GMT
      Content-Length: 291
	  ```json
      {
          "id": "51",
          "title": "Eight more COVID-related deaths in NSW",
          "date": "2016-09-22",
          "body": "NSW Health is today reporting the deaths of eight people with COVID-19 - five women and three men.",
          "tags": [
              "science",
              "health",
              "lifestyle"
          ]
      }
	  ```



      iii. Get a specific article using the id
      When the client makes a request to GET /articles/{id}, article whose ID matches the id path parameter.
      
      $ curl http://localhost:8080/articles/51
```json
      {
          "id": "51",
          "title": "Eight more COVID-related deaths in NSW",
          "date": "2016-09-22",
          "body": "NSW Health is today reporting the deaths of eight people with COVID-19 - five women and three men.",
          "tags": [
              "science",
              "health",
              "lifestyle"
          ]
      }
```

      iv. The final endpoint, GET /tags/{tagName}/{date} will return the list of articles that have that tag name on the given date and some summary data about that tag for that day.
      $ curl http://localhost:8080/tags/health/20160922    
	  
  
 ```json
  	 {
          "tag": "health",
          "count": 2,
          "Articles": [
              "1",
              "51"
          ],
          "related_tags": [
              "fitness",
              "science",
              "lifestyle"
          ]
      }
 ```
