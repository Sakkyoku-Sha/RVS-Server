package server

import (
	"encoding/binary"
)

//LangID is the Byte value of the "Enum" Request Type
type LangID byte

//Enum Specification for LangId
const (
	ENG LangID = iota // 0
	JPN               // 1
)

//VideoFileRequest is the []byte structure of a containing APIRequest which is of type "Video"
type VideoFileRequest struct {
	XResolution uint16 //2 bytes
	YResolution uint16 //2 bytes
	Language    LangID //1 byte
	SubTitle    LangID //1 byte
	FileName    string //Remaining Bytes
}

//NewVideoFileRequest Creates a new VideoFile request given the bytes from a properly tagged
func NewVideoFileRequest(apiRequest APIRequest, byteOrder binary.ByteOrder) VideoFileRequest {

	const offSet = APIRequestTCPOffset

	var TCPBytes = apiRequest.TCPBytes

	var xResolutionBytes = []byte{TCPBytes[offSet], TCPBytes[offSet+1]}
	var xResolution = binary.ByteOrder.Uint16(byteOrder, xResolutionBytes)

	var yResolutionBytes = []byte{TCPBytes[offSet+2], TCPBytes[offSet+3]}
	var yResolution = binary.ByteOrder.Uint16(byteOrder, yResolutionBytes)

	var language = LangID(TCPBytes[offSet+4])
	var subtitle = LangID(TCPBytes[offSet+5])

	var fileName = string(TCPBytes[offSet+6:])

	var returned = VideoFileRequest{xResolution, yResolution, language, subtitle, fileName}
	return returned
}
