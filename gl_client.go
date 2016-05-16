// GitLab client package. Client gives an ability to parse GitLab json API
// version 3 and response with proper structures.
package gl_client

import (
  "fmt"
  "net/http"
  "strconv"
  "io/ioutil"
  r "github.com/igkuz/gl_client/resources"
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

// 
func (c *GLClient) GetUsers() ([]r.User, error) {
  url := c.BaseUrl + "/api/" + c.ApiVersion + "/users"

  body, err := c.makeRequest(url)
  if err != nil {
    return nil, err
  }

  users, err := r.ParseUsers(body)

  return users, err
}

// Get single user by ID.
func (c *GLClient) GetUser(id int64) (*r.User, error) {
  url := c.BaseUrl + "/api/" + c.ApiVersion + "/users/" + strconv.FormatInt(id, 10)

  body, err := c.makeRequest(url)
  if err != nil {
    return nil, err
  }

  user, err := r.ParseUser(body)

  return &user, err
}

// Get current authenticated user.
func (c *GLClient) GetCurrentUser() (*r.User, error) {
  url := c.BaseUrl + "/api/" + c.ApiVersion + "/user"

  body, err := c.makeRequest(url)
  if err != nil {
    return nil, err
  }

  user, err := r.ParseUser(body)

  return &user, err
}

// Get current authenticated user emails.
func (c *GLClient) GetCurrentUserEmails() (*[]r.UserEmail, error) {
  url := c.BaseUrl + "/api/" + c.ApiVersion + "/user/emails"

  body, err := c.makeRequest(url)
  if err != nil {
    return nil, err
  }

  userEmail, err := r.ParseUserEmails(body)

  return userEmail, err
}
