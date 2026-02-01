package models

type Team struct {
  ID      string `json:"id"`
  Name    string `json:"name"`
  Coach   string `json:"coach"`
  Stadium string `json:"stadium"`
  Info    string `json:"info"`
}