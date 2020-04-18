package comments

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"

	"database/sql"

	

	"github.com/aws/aws-lambda-go/events"
)

type BodyRequest struct {
	RequestName string `json:"name"`
}

type Response struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id          int    `json:"id"`
	Slug        string `json:"slug"`
	Body        string `json:"body"`
	Author      string `json:"author"`
	Posted      string `json:"posted"`
	Picture     string `json:"picture,omitempty"`
	Thumbs_up   int    `json:"thumbs_up,omitempty"`
	Thumbs_down int    `json:"thumbs_down,omitempty"`
	Heart       int    `json:"heart,omitempty"`
	Smile       int    `json:"smile,omitempty"`
}

type Vote struct {
	Id   int    `json:"id"`
	Vote string `json:"vote"`
}

// GetComments -
// Returns the comments for the given post
func GetComments(request events.APIGatewayProxyRequest, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	fmt.Println(request.QueryStringParameters["slug"])
	slug := request.QueryStringParameters["slug"]

	results, err := db.Query("SELECT * FROM comments WHERE slug = ?", slug)
	if err != nil {
		panic(err.Error())
	}

	var comments []Comment

	for results.Next() {
		var comment Comment
		err = results.Scan(&comment.Id, &comment.Body, &comment.Slug, &comment.Posted, &comment.Author, &comment.Picture, &comment.Thumbs_up, &comment.Thumbs_down, &comment.Heart, &comment.Smile)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%+v\n", comment)

		comments = append(comments, comment)
	}

	response := Response{
		Comments: comments,
	}

	fmt.Printf("%+v\n", comments)

	jsonResults, err := json.Marshal(response)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", string(jsonResults))

	return events.APIGatewayProxyResponse{
		Body:       string(jsonResults),
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// UpdateComment -
// Updates the comment
func UpdateComment(request events.APIGatewayProxyRequest, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	fmt.Println("Adding a Vote to a Comment")

	fmt.Printf("Request: %v\n", request)

	fmt.Println("Received body: ", request.Body)
	body, _ := base64.StdEncoding.DecodeString(request.Body)
	fmt.Println(string(body))

	var vote Vote

	err = json.Unmarshal(body, &vote)
	if err != nil {
		panic(err.Error())
	}

	switch vote.Vote {
	case "thumbs_up":
		stmt, err := db.Prepare("UPDATE comments SET thumbs_up = thumbs_up+1 WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()
		_, err = stmt.Exec(vote.Id)
		if err != nil {
			panic(err.Error())
		}
	case "thumbs_down":
		stmt, err := db.Prepare("UPDATE comments SET thumbs_down = thumbs_down+1 WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()
		_, err = stmt.Exec(vote.Id)
		if err != nil {
			panic(err.Error())
		}
	case "heart":
		stmt, err := db.Prepare("UPDATE comments SET heart = heart+1 WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()
		_, err = stmt.Exec(vote.Id)
		if err != nil {
			panic(err.Error())
		}
	case "smile":
		stmt, err := db.Prepare("UPDATE comments SET smile = smile+1 WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		defer stmt.Close()
		_, err = stmt.Exec(vote.Id)
		if err != nil {
			panic(err.Error())
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       "Successfull Vote!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// PostComment -
// Adds a new comment to the site
func PostComment(request events.APIGatewayProxyRequest, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	fmt.Println("Posting a Comment")

	fmt.Printf("Request: %v\n", request)

	fmt.Println("Received body: ", request.Body)
	body, err := base64.StdEncoding.DecodeString(request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))

	var comment Comment

	err = json.Unmarshal([]byte(request.Body), &comment)
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare("INSERT INTO comments(body, slug, posted, author, picture, thumbs_up, thumbs_down, heart, smile) VALUES (?, ?, current_timestamp(), ?, ?, 0, 0, 0, 0)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(comment.Body, comment.Slug, comment.Author, comment.Picture)
	if err != nil {
		panic(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       "Post Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// DeleteComment -
// Deletes the comment with the ID
func DeleteComment(request events.APIGatewayProxyRequest, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	return events.APIGatewayProxyResponse{
		Body:       "Delete Request!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

// AllComments -
// Returns all comments that have been posted to the site
func AllComments(request events.APIGatewayProxyRequest, db *sql.DB) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Getting Comments")

	return events.APIGatewayProxyResponse{
		Body:       "Get All Comments!",
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}
