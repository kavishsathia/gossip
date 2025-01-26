package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/shared"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDatabase() (*gorm.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		print("NOT SET")
		return nil, fmt.Errorf("DATABASE_URL not set")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func OpenRedis() (*redis.Client, error) {
	dbURL := os.Getenv("REDIS_URL")
	if dbURL == "" {
		print("NOT SET")
		return nil, fmt.Errorf("DATABASE_URL not set")
	}

	return redis.NewClient(&redis.Options{
		Addr: dbURL,
		DB:   0,
	}), nil
}

func GetUserInfo(c *gin.Context) (interface{}, *User, error) {
	user, ok := c.Get("user")
	if !ok {
		return nil, nil, fmt.Errorf("failed to get user")
	}

	userInfo, ok := user.(*User)
	if !ok {
		return user, nil, fmt.Errorf("failed to cast userInfo")
	}

	return user, userInfo, nil
}

func VerifyJWT(tokenString string) (*User, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, errors.New("JWT secret key not found in environment")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("invalid token: " + err.Error())
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims format")
	}

	userID, ok := claims["userId"].(float64)
	if !ok {
		return nil, errors.New("userId claim not found or invalid")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("username claim not found or invalid")
	}

	return &User{
		Username: username,
		UserID:   int(userID),
	}, nil
}

func GenerateJWT(userId int, username string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		return "", errors.New("JWT private key not found in environment")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"username": username,
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		print(err.Error())
		return "", errors.New("failed to sign token")
	}

	return tokenString, nil
}

func Moderate(c *gin.Context, text string) (*string, error) {
	moderationClient := openai.NewClient()

	res, err := moderationClient.Moderations.New(c, openai.ModerationNewParams{
		Model: openai.F("omni-moderation-latest"),
		Input: openai.F[openai.ModerationNewParamsInputUnion](shared.UnionString(text)),
	})

	if err != nil {
		return nil, err
	}

	if !res.Results[0].Flagged {
		return nil, nil
	}

	var flagged map[string]bool
	json.Unmarshal([]byte(res.Results[0].Categories.JSON.RawJSON()), &flagged)

	for category, isFlagged := range flagged {
		if isFlagged {
			return &category, nil
		}
	}

	return nil, nil
}

func GenerateDescription(c *gin.Context, text string) (string, error) {
	descriptionClient := openai.NewClient()

	res, err := descriptionClient.Chat.Completions.New(c, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(`
				Provide a description for this post, 
				only return the description and nothing else. 
				Do not use markdown.
				The description should be short and concise, maybe only one short sentence
				The post: ` + text),
		}),
		Model: openai.F(openai.ChatModelGPT4oMini),
	})

	if err != nil {
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}

func GenerateCorrections(c *gin.Context, text string) ([]string, error) {
	factCheckClient := openai.NewClient()

	res, err := factCheckClient.Chat.Completions.New(c, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(`
				Fact check this piece of text. 
				Return a JSON object containing a key called corrections.
				The type for the value is an array of strings.
				Each string is a correction for the post.
				If the post is does not have critical errors, just
				return an empty array.
				The text: ` + text),
		}),
		Model: openai.F(openai.ChatModelGPT4oMini),
		ResponseFormat: openai.F[openai.ChatCompletionNewParamsResponseFormatUnion](shared.ResponseFormatTextParam{
			Type: openai.F(shared.ResponseFormatTextType(shared.ResponseFormatJSONObjectTypeJSONObject)),
		}),
	})

	if err != nil {
		return []string{}, err
	}

	var corrections map[string][]string

	json.Unmarshal([]byte(res.Choices[0].Message.Content), &corrections)

	return corrections["corrections"], nil
}
