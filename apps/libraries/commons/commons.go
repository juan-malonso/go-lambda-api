package commons


type Body = interface{}

type Config = interface{}

type Headers struct {
	Origin string `json:"origin,omitempty"`
}

type Event struct {
	StatusCode int `json:"statusCode,omitempty"`

	ConfigId string `json:"configId,omitempty"`
	Config Config `json:"config,omitempty"`
	
	Body Body `json:"body,omitempty"`
	Headers Headers `json:"headers,omitempty"`
}
