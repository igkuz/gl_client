package main

import (
  "fmt"
  "net/http"
  "os"
  "crypto/tls"
  "encoding/json"
  "io/ioutil"
)

type GLClient struct {
  AuthToken   string
  BaseUrl     string
  Client      *http.Client
  ApiVersion  string
}

type User struct {
  Id          int64   `json:"id"`
  Name        string  `json:"name"`
  Username    string  `json:"username"`
  Email       string  `json:"email"`
}

func (c *GLClient) MakeRequest(url string) ([]byte, error) {
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("PRIVATE-TOKEN", c.AuthToken)
  resp, err := c.Client.Do(req)

  defer resp.Body.Close()

  if err != nil {
    fmt.Println("Error occured during request: ", err)
    return nil, err
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  return []byte(body), err
}

func ParseUsers(body []byte) ([]User, error) {
  u := []User{}

  err := json.Unmarshal(body, &u)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  return u, err
}

func (c *GLClient) getUsers() ([]User, error) {
  url := c.BaseUrl + "/api/" + c.ApiVersion + "/users"

  body, err := c.MakeRequest(url)
  if err != nil {
    return nil, err
  }

  users, err := ParseUsers(body)

  return users, err
}

func main() {
  transport := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }

  gc := &GLClient{AuthToken: os.Getenv("GL_TOKEN"),
    BaseUrl: os.Getenv("GL_HOST"), 
    Client: &http.Client{Transport: transport}, 
    ApiVersion: "v3",
  }
  users, err := gc.getUsers()
  fmt.Println("Users: ", users, "Error: ", err)
}
