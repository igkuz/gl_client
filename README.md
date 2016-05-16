### Client for GitLab API written in GO

Package provides methods for pasing GitLab API and return structures that are easy to work with.
Currently supports only last api version – v3.

# Initializing

```go
apiToken := "123"
baseUrl := "https://git.example.com"
apiVersion := "v3"
client := &http.Client{}
gitLabClient := gl_client.NewClient(apiToken, baseUrl, apiVersion, client)
users, err := gitLabClient.GetUsers()
```
