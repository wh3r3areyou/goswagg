package tag

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"log"
)

// Tag the tag model
type Tag struct {
	ID   int64
	Name string
}

// GetTags Get tags in swagger doc
func GetTags() []Tag {
	strTags, err := jsoniter.Marshal(viper.Get("tags"))

	if err != nil {
		log.Fatal("can`t marshal tags in swagger. Check ur swagger documentation for valid")
	}

	var tags []Tag

	err = jsoniter.Unmarshal(strTags, &tags)
	if err != nil {
		log.Fatal("Can`t get tags. Check ur swagger documentation for valid")
	}

	return tags
}
