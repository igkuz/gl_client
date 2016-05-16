package resources

import (
  "fmt"
  "encoding/json"
)

// User structure used to represent API response with object.
type User struct {
  Id          int64   `json:"id"`
  Name        string  `json:"name"`
  Username    string  `json:"username"`
  Email       string  `json:"email"`
}

type UserEmail struct {
  Id          int64   `json:"id"`
  Email       string  `json:"email"`
}

// parseUsers function parses response body from API. Unmarshals json to User structure.
func ParseUsers(body []byte) ([]User, error) {
  u := []User{}

  err := json.Unmarshal(body, &u)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  return u, err
}

func ParseUser(body []byte) (User, error) {
  u := User{}

  err := json.Unmarshal(body, &u)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  return u, err
}

func ParseUserEmails(body []byte) (*[]UserEmail, error) {
  ue := &[]UserEmail{}

  err := json.Unmarshal(body, ue)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  return ue, err
}
