### Client for GitLab API written in GO

Package provides methods for pasing GitLab API and return structures that are easy to work with.
Currently supports only last api version – v3.

### Initializing

```go
apiToken := "123"
baseUrl := "https://git.example.com"
apiVersion := "v3"
client := &http.Client{}

gitLabClient := gl_client.NewClient(apiToken, baseUrl, apiVersion, client)
users, err := gitLabClient.GetUsers()
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
users, err := gitLabClient.GetUsers()
```
