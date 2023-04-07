package main

import (
  "encoding/json"
  "fmt"
)

type keyword struct {
	keyword string json:"keyword"
	Status  string json:"status"
	Prefix  string json:"prefix"
  }
  
  func main() {
	Keyword := []string{"bonfire", "bool"}
	result := make([]keyword, len(Keyword))
  
	for i, keyword := range Keyword {
	  result[i].keyword = keyword
	  if foundInServer(keyword) {
		result[i].Status = "found!"
		result[i].Prefix = getsmalluniquePrefix(keyword)
	  } else {
		result[i].Status = "not found!"
		result[i].Prefix = "not_applicable"
	  }
	}
  
	jsonResult, _ := json.Marshal(result)
	fmt.Print(string(jsonResult))
  }
  
  func foundInServer(keyword string) bool {
	return true
  }
  
  func getsmalluniquePrefix(keyword string) string {
    return "bonf"
  }