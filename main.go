package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Post struct {
	Id      string
	Content string
	Author  string
}

func main() {
	// creating csv file
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		{Id: "1", Content: "Golang Functions", Author: "Rob Pike"},
		{Id: "2", Content: "Go web progrqmming", Author: "Ken Thompson"},
		{Id: "3", Content: "Golang concurrency", Author: "Robert Greismar"},
		{Id: "4", Content: "Google internships", Author: "King Wood"},
	}

	//writing data to the csvFile
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{post.Id, post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	// Reading data from csv file
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post

	for _, item := range record {
		post := Post{Id: item[0], Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)

}
