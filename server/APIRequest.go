package server

//RequestType is the int value of the "Enum" Request Type
type RequestType int

//Determines the Type of Request
const (
	Overview   RequestType = iota // 0
	Video                         // 1
	BadRequest                    // 2
)

//APIRequest is the basic tagging of what kind of request a byte[] was before it can be processed.
type APIRequest struct {
	Type  RequestType
	Bytes []byte
}
