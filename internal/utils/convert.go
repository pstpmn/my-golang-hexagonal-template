package utils

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertStrToObjectId(value string, target *primitive.ObjectID) error {
	_, err := hex.DecodeString(value)
	if err != nil {
		return fmt.Errorf("invalid hexadecimal string: %v", err)
	}
	*target, err = primitive.ObjectIDFromHex(value)
	return err
}

func ConvertToHyphenated(input string) string {
	return strings.ReplaceAll(input, " ", "-")
}

func ConvertToSpaceSeparated(input string) string {
	return strings.ReplaceAll(input, "-", " ")
}

func SnippetFromContent(content string, maxLength int) string {
	htmlTagRegex := regexp.MustCompile(`<[^>]*>`)
	content = htmlTagRegex.ReplaceAllString(content, "")

	words := strings.Fields(content)
	var shortenedContent string
	if len(words) > maxLength {
		shortenedContent = strings.Join(words[:maxLength], " ")
		shortenedContent = shortenedContent + " ..."
	} else {
		shortenedContent = content
	}
	return shortenedContent
}
