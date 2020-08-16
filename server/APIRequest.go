package server

//RequestType is the byte value of the "Enum" Request Type
type RequestType byte

//Enum Specification of RequestType
const (
	Overview   RequestType = iota // 0
	Video                         // 1
	BadRequest                    // 2
)

//NewAPIRequest Creates a new API request from the TCP bytes
func NewAPIRequest(TCPBytes []byte) *APIRequest {

	if len(TCPBytes) == 0 {
		var badRequest = APIRequest{BadRequest, TCPBytes}
		return &badRequest
	}

	var requestType RequestType
	requestType = RequestType(TCPBytes[0])
	if requestType >= BadRequest {
		requestType = BadRequest
	}

	var apiRequest = APIRequest{requestType, TCPBytes}
	return &apiRequest
}

//APIRequestTCPOffset is the Offset ApiRequestHeader Data
const APIRequestTCPOffset uint32 = 1

//APIRequest is the basic tagging of what kind of request a byte[] was before it can be processed.
type APIRequest struct {
	Type     RequestType // 1 Byte
	TCPBytes []byte
}
