package tagshelf

import "io"

// Requester defines tagshelf interaction contract
type Requester interface {
	Clienter

	Status() (Responder, error)
	WhoAmI() (Responder, error)
	Ping() (Responder, error)
	FileUpload(*File) (Responder, error)
	FileDetail(string) (Responder, error)
	JobDetail(string) (Responder, error)
	CompanyInbox(string) (Responder, error)
}

// Authorizer defines the behavior that authorizes a request
type Authorizer interface {
	Auth() (Requester, error)
}

// Signer defines the behavior of a signed request
type Signer interface {
	Sign(string, string, io.Reader) error
}

// Responder tagshelf object response contract
type Responder interface {
	Status() int
	Body() Payload
}

type Clienter interface {
	SetQuery(query map[string]string)
	Query() map[string]string
}
