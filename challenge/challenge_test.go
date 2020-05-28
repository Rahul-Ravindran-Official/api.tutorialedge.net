package challenge_test

// func TestPostComments(t *testing.T) {
// 	fmt.Println("Testing Post Comments...")
// 	db, err := database.GetDBConn()
// 	if err != nil {
// 		t.Log(err)
// 		t.Error("Could not get DB Connection")
// 		return
// 	}

// 	body := []byte(`{
// 		"slug": "/challenges/go/01/",
// 		"body": "package main",
// 		"author": "Elliot Forbes",
// 		"picture": "https://lh3.googleusercontent.com/a-/AOh14GjvaxQsEaZVJi809M65ACo6BogR8nI-Ntl7FuBkEA",
// 		"user": {
// 			"sub": "google-oauth2|116080913132292461306",
// 			"given_name": "Elliot",
// 			"family_name": "Forbes",
// 			"nickname": "efgamercity",
// 			"name": "Elliot Forbes",
// 			"picture": "https://lh3.googleusercontent.com/a-/AOh14GjvaxQsEaZVJi809M65ACo6BogR8nI-Ntl7FuBkEA",
// 			"locale": "en-GB",
// 			"updated_at": "2020-05-04T08:03:18.105Z"
// 		}
// 	}`)

// 	request := events.APIGatewayProxyRequest{}
// 	request.Body = string(body)

// 	tokenInfo := auth.TokenInfo{
// 		Sub: "-1",
// 	}

// 	response, err := comments.PostComment(request, tokenInfo, db)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	fmt.Println(response)

// 	var comment comments.Comment
// 	db.Where("slug = ?", "/random-test-slug/").Find(&comment)

// 	if comment.Body != "Test" {
// 		t.Fail()
// 	}

// 	// 	db.Delete(&comment)
// }
