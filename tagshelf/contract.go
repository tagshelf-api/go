package tagshelf

// Requester defines tagshelf interaction contract
type Requester interface {
	Status() (Responder, error)
	WhoAmI() (Responder, error)
	// Ping() (Responder, error)
	// FileUpload(File) (Responder, error)
	// FileDetail(string) (Responder, error)
	// JobDetail(string) (Responder, error)
}

type Authorizer interface {
	Auth() (Requester, error)
}

// Responder tagshelf object response contract
type Responder interface {
	Status() int
	Body() interface{}
}
