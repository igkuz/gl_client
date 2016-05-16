// GitLab client package. Client gives an ability to parse GitLab json API
// version 3 and response with proper structures.
package gl_client

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "strconv"
)

// GLClient structure – the client instance. Gives methods to parse GitLab API.
//    AuthToken - stores api token from GitLab
//    BaseUrl – stores base url for gitlab instance, like http(s)://git.exmaple.com/
type GLClient struct {
  AuthToken   string
  BaseUrl     string
  Client      *http.Client
  ApiVersion  string
}

// User structure used to represent API response with object.
type User struct {
  Id          int64   `json:"id"`
  Name        string  `json:"name"`
  Username    string  `json:"username"`
  Email       string  `json:"email"`
}

// NewClient returns GLClient instance.
func NewClient(token string, host string, apiVersion string, httpClient *http.Client) *GLClient {
  return &GLClient{
    AuthToken: token,
    BaseUrl: host,
    Client: httpClient,
    ApiVersion: apiVersion,
  }
}

// makeRequest - makes GET request to provided url. Returns byte sequence or error.
// Body represented as byte array can be next unmarshalled with json library to structure.
func (c *GLClient) makeRequest(url string) ([]byte, error) {
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

// parseUsers function parses response body from API. Unmarshals json to User structure.
func parseUsers(body []byte) ([]User, error) {
  u := []User{}

  err := json.Unmarshal(body, &u)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  return u, err
}

func parseUser(body []byte) (User, error) {
  u := User{}

  err := json.Unmarshal(body, &u)
  if err != nil {
    fmt.Println("Error: ", err)
  }

  return u, err
}

// 
func (c *GLClient) GetUsers() ([]User, error) {
  url := c.BaseUrl + "/api/" + c.ApiVersion + "/users"

  body, err := c.makeRequest(url)
  if err != nil {
    return nil, err
  }

  users, err := parseUsers(body)

  return users, err
}

// Get single user by ID.
func (c *GLClient) GetUser(id int64) (*User, error) {
  url := c.BaseUrl + "/api/" + c.ApiVersion + "/users/" + strconv.FormatInt(id, 10)

  body, err := c.makeRequest(url)
  if err != nil {
    return nil, err
  }

  user, err := parseUser(body)

  return &user, err
}
