package helpers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/shared"
)

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
