# Go Client

TagShelf API Client Library for Go

## Usage

-   Import lib

```
import "github.com/tagshelf-api/go/tagshelf"
```

-   Get Client

Sets up the client configuration

```
// for APP API KEY client
config := tagshelf.Config{AppApiKey: "YOU_API_KEY"}
```

or

```
// for HMAC client
config := tagshelf.Config{
	SecretKey: "SECRET_KEY",
	ApiKey:    "API_KEY",
}
```

or

```
// for OAuth client
config := tagshelf.Config{
	GrantType: "GRANT_TYPE",
	User:      "USERNAME",
	Pass:      "PASSWORD",
}
```

Gets client

```
client, err := tagshelf.New(config)
if err != nil {
	// handle error
}
```

-   Use client

```
// ========= Service Status
status, err := client.Status()
if err != nil {
	// handle error
}

// ========= Authenticated user details
me, _ := client.WhoAmI()

// ========= Ping service
pong, _ := client.Ping()

// ========= File Detail
file, _ := client.FileDetail("some-file-id")

// ========= File Upload
upload := tagshelf.NewFileUpload()
upload.Add("some-link-to-file")
// or we can add multiple urls like this:
// upload.Add("some-link-to-file", "some-link-to-file", "some-link-to-file",...)
// Upload Merge flag
upload.Merge = false // true/false optional value
// Upload Metadata
meta := tagshelf.FileMetadata{"name": "John"}
upload.AddMeta(meta)
fileUpload, _ := client.FileUpload(upload)

// ========= Job Detail
jobDetail, _ := client.JobDetail("some-job-id")
```
