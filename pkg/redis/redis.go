package redis

import (
	"log"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/jalilbengoufa/go-search/pkg/setting"
)

var RedisConn *redisearch.Client
var RedisAutocompleter *redisearch.Autocompleter

func Setup() {

	RedisConn = redisearch.NewClient(setting.RedisearchSetting.Host, "myIndex")
	RedisAutocompleter = redisearch.NewAutocompleter(setting.RedisearchSetting.Host, "myAutocompleter")

	// Create a schema
	sc := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("title", redisearch.TextFieldOptions{Weight: 5.0, Sortable: true})).
		AddField(redisearch.NewTextField("desc"))

	// Drop an existing index. If the index does not exist an error is returned
	RedisConn.Drop()

	// Create the index with the given schema
	if err := RedisConn.CreateIndex(sc); err != nil {
		log.Fatal(err)
	}
}

func Insert(title string, desc string, id int) error {

	// Create a document with an id and given score
	doc := redisearch.NewDocument(strconv.Itoa(id), 1.0)
	doc.Set("title", title).
		Set("des", desc)

	// Index the document. The API accepts multiple documents at a time
	if err := RedisConn.Index([]redisearch.Document{doc}...); err != nil {
		log.Fatal(err)
		return err
	}
	terms := make([]redisearch.Suggestion, 1)
	terms[0] = redisearch.Suggestion{Term: title, Score: 5.0}

	RedisAutocompleter.AddTerms(terms...)
	return nil
}

func Find(word string) (docs []redisearch.Document, total int, err error) {

	// Searching with limit and sorting
	docs, total, err = RedisConn.Search(redisearch.NewQuery(word).
		Limit(0, 5).
		SetReturnFields("title"))

	// Index the document. The API accepts multiple documents at a time
	if err != nil {
		log.Fatal(err)
		return nil, 0, err
	}

	return docs, total, nil
}

func Autocomplete(word string) (suggestions []redisearch.Suggestion, err error) {

	suggestions, err = RedisAutocompleter.SuggestOpts(word, redisearch.SuggestOptions{Num: 5, WithScores: true, Fuzzy: true})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return suggestions, nil
}

func SpellCheck(word string) (suggestions []redisearch.MisspelledTerm, total int, err error) {

	suggestions, total, err = RedisConn.SpellCheck(redisearch.NewQuery(word), &redisearch.SpellCheckOptions{Distance: 4})

	if err != nil {
		log.Fatal(err)
		return nil, 0, err
	}

	return suggestions, total, nil
}
