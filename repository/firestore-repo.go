package repository

import (
	"context"
	"golang-rest-api/entity"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	//"cloud.google.com/go/firestore"  SE ENCHOLA NO SE TOCO COPIAR LA DE mtnaghibi
	//"google.golang.org/api/iterator"
)

// Implement the interface
type repo struct{}

// NewFirestoreRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

// g3notype-ska

const (
	// projectId string = "rare-lambda-415802" // cuenta free
	projectId      string = "g3notype-ska" // firebase trin
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {

	credentialsOption := option.WithCredentialsFile("./repository/g3notype-ska-firebase-adminsdk-jf6zl-99fea81c23.json")
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectId, credentialsOption)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {

	// credentialsOption := option.WithCredentialsJSON(decodedCreds)
	// credentialsOption := option.WithCredentialsJSON("./g3notype-ska-firebase-adminsdk-jf6zl-99fea81c23.json")  // requiere parametro en bytes

	credentialsOption := option.WithCredentialsFile("./repository/g3notype-ska-firebase-adminsdk-jf6zl-99fea81c23.json") // cannot read credentials file: open

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, credentialsOption)
	//client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to Create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post

	itr := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
