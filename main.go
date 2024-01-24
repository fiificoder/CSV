package main

import (
	"encoding/csv"
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
		{Id: "4", Content: "Dev internships", Author: "Max Wood"},
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

	_ = record
	


}
