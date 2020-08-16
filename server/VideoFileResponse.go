package server

import "encoding/binary"

//INITIAL RESPONSE =================================================================================================

//VideoFileType is a "enum" for what kind of file is being dealt with.
type VideoFileType byte

//VideoFileType Enum Specification (Only MKV to start)
const (
	MKV VideoFileType = iota // 0
	MP4                      // 1
)

const sizeOfInitResponse = 12

//VideoFileInitialResponse is the definition of the data that is sent back from a VideoFileRequest
type VideoFileInitialResponse struct {
	FileFound byte          //1 Byte
	FileType  VideoFileType //1 Byte
	ChunkSize uint64        //8 Bytes
	FPS       uint16        //2 Bytes
}

// InitResponseToBytes converts the VideoFileInitResponse to a byte slice
func InitResponseToBytes(response VideoFileInitialResponse, byteOrder binary.ByteOrder) [sizeOfInitResponse]byte {

	var bytes [sizeOfInitResponse]byte

	bytes[0] = response.FileFound
	bytes[1] = byte(response.FileType)

	binary.ByteOrder.PutUint64(byteOrder, bytes[2:], response.ChunkSize)
	binary.ByteOrder.PutUint16(byteOrder, bytes[10:], response.FPS)

	return bytes
}

//STREAMING CHUNKS
