### Client for GitLab API written in GO

Package provides methods for pasing GitLab API and return structures that are easy to work with.
Currently supports only last api version – v3.

Full API documentation with request params and response examples can be found [here](http://docs.gitlab.com/ce/api/)

### Initializing

```go
apiToken := "123"
baseUrl := "https://git.example.com"
apiVersion := "v3"
client := &http.Client{}

gitLabClient := gl_client.NewClient(apiToken, baseUrl, apiVersion, client)
```

Inhouse instances of GitLab sometimes use self signed SSL certificates. For proper work with this case you should skip verifying of certificate.
So initializing process will look slightly different.

```go
transport := &http.Transport{
  TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
apiToken := "123"
baseUrl := "https://git.example.com"
apiVersion := "v3"
client := &http.Client{Transport: transport}

gitLabClient := gl_client.NewClient(apiToken, baseUrl, apiVersion, client)
```

### Get users

```
GET /users
```

```go
users, err := gitLabClient.GetUsers()
```

Response:

```
[{1 John Smith j_smith j_smith@exmaple.com} {2 John Galt j_galt j_galt@exmaple.com}]
```

### Single user

```
GET /users/:id
```

```go
user, err := gitLabClient.GetUser(1)
```

Response:

```
{ 1 John Smith j_smith j_smith@example.com }
```

#### Get current authenticated user

```
GET /user
```

```go
user, err := gitLabClient.GetCurrentUser()
```

Response:

```
{ 1 John Smith j_smith j_smith@example.com }
```
