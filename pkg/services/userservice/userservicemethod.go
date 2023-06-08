package userservice

import (
	"context"
	"errors"
	"time"

	models "github.com/Dazzler/My-RestServer/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceMethod struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceMethod{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (us *UserServiceMethod) CreateUser(user *models.User) error {
	// Check if the user already exists
	existingUser := &models.User{}
	err := us.userCollection.FindOne(us.ctx, bson.M{"username": user.Username}).Decode(existingUser)
	if err == nil {
		return errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Store the user in the database
	user.Password = string(hashedPassword)
	_, err = us.userCollection.InsertOne(us.ctx, user)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (us *UserServiceMethod) Login(user *models.User) (string, error) {
	// Find the user in the database
	existingUser := &models.User{}
	err := us.userCollection.FindOne(us.ctx, bson.M{"username": user.Username}).Decode(existingUser)
	if err != nil {
		return "", errors.New("username does not exist")
	}

	// Compare the provided password with the stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("fail to login")
	}

	// Generate a JWT token for the user
	token, err := generateToken(existingUser.ID.Hex())
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func generateToken(userID string) (string, error) {
	// Generate and return a JWT token
	// You can use a library like "github.com/dgrijalva/jwt-go" to generate JWT tokens
	// Here's an example of how you can generate a token:

	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return signedToken, nil

	// Replace the code above with your actual implementation using a JWT library
}
