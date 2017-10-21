/*
author: foolbread
file: storage/buffer.go
date: 2017/10/21
*/
package storage

const (
	BUFFER_2MB   = 1
	BUFFER_4MB   = 2
	BUFFER_8MB   = 3
	BUFFER_16MB  = 4
	BUFFER_32MB  = 5
	BUFFER_64MB  = 6
	BUFFER_128MB = 7
	BUFFER_256MB = 8
	BUFFER_MAX   = 9
)

const (
	BUFFER_SIZE_1MB   = 1024*1024
	BUFFER_SIZE_2MB   = BUFFER_SIZE_1MB * 2
	BUFFER_SIZE_4MB   = BUFFER_SIZE_1MB * 4
	BUFFER_SIZE_8MB   = BUFFER_SIZE_1MB * 8
	BUFFER_SIZE_16MB  = BUFFER_SIZE_1MB * 16
	BUFFER_SIZE_32MB  = BUFFER_SIZE_1MB * 32
	BUFFER_SIZE_64MB  = BUFFER_SIZE_1MB * 64
	BUFFER_SIZE_128MB = BUFFER_SIZE_1MB * 128
	BUFFER_SIZE_256MB = BUFFER_SIZE_1MB * 256
)
type BUFFER_TYPE int

type fbFTPBufferManager struct {

}

type fbFTPBuffer struct {
	t	BUFFER_TYPE
	buf []byte
}

func newfbFTPBuffer(t BUFFER_TYPE)*fbFTPBuffer{
	r := new(fbFTPBuffer)
	r.t = t

	switch t {
	case BUFFER_2MB:
		r.buf = make([]byte,BUFFER_SIZE_2MB,BUFFER_SIZE_2MB)
	case BUFFER_4MB:
		r.buf = make([]byte,BUFFER_SIZE_4MB,BUFFER_SIZE_4MB)
	case BUFFER_8MB:
		r.buf = make([]byte,BUFFER_SIZE_8MB,BUFFER_SIZE_8MB)
	case BUFFER_16MB:
		r.buf = make([]byte,BUFFER_SIZE_16MB,BUFFER_SIZE_16MB)
	case BUFFER_32MB:
		r.buf = make([]byte,BUFFER_SIZE_32MB,BUFFER_SIZE_32MB)
	case BUFFER_64MB:
		r.buf = make([]byte,BUFFER_SIZE_64MB,BUFFER_SIZE_64MB)
	case BUFFER_128MB:
		r.buf = make([]byte,BUFFER_SIZE_128MB,BUFFER_SIZE_128MB)
	case BUFFER_256MB:
		r.buf = make([]byte,BUFFER_SIZE_256MB,BUFFER_SIZE_256MB)
	default:
		r.t = BUFFER_256MB
		r.buf = make([]byte,BUFFER_SIZE_256MB,BUFFER_SIZE_256MB)
	}

	return r
}
