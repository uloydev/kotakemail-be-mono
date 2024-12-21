package storage

import (
	"io"
	"time"

	appcontext "kotakemail.id/pkg/context"
)

type File struct {
	io.ReadCloser // Underlying data.
	Attributes
}

// Attributes represents the metadata of a File
// Inspired from gocloud.dev/blob.Attributes
type Attributes struct {
	// ContentType is the MIME type of the blob object. It will not be empty.
	ContentType string
	// ContentEncoding specifies the encoding used for the blob's content, if any.
	ContentEncoding string
	// Metadata holds key/value pairs associated with the blob.
	// Keys are guaranteed to be in lowercase, even if the backend provider
	// has case-sensitive keys (although note that Metadata written via
	// this package will always be lowercased). If there are duplicate
	// case-insensitive keys (e.g., "foo" and "FOO"), only one value
	// will be kept, and it is undefined which one.
	Metadata map[string]string
	// ModTime is the time the blob object was last modified.
	ModTime time.Time
	// CreationTime is the time the blob object was created.
	CreationTime time.Time
	// Size is the size of the object in bytes.
	Size int64
}

type ReaderOptions struct {
	// ReadCompressed controls whether the file must be uncompressed based on Content-Encoding.
	// Only respected by Google Cloud Storage: https://cloud.google.com/storage/docs/transcoding
	// Common pitfall: https://github.com/googleapis/google-cloud-go/issues/1743
	ReadCompressed bool
}

// WriterOptions are used to modify the behaviour of write operations.
// Inspired from gocloud.dev/blob.WriterOptions
// Not all options are supported by all FS
type WriterOptions struct {
	Attributes Attributes

	// BufferSize changes the default size in bytes of the chunks that
	// Writer will upload in a single request; larger blobs will be split into
	// multiple requests.
	//
	// This option may be ignored by some drivers.
	//
	// If 0, the driver will choose a reasonable default.
	//
	// If the Writer is used to do many small writes concurrently, using a
	// smaller BufferSize may reduce memory usage.
	BufferSize int
}

type Storage interface {
	Name() string
	Write(ctx *appcontext.AppContext, path string, options *WriterOptions) (io.WriteCloser, error)
	Read(ctx *appcontext.AppContext, path string, options *ReaderOptions) (*File, error)
	Delete(ctx *appcontext.AppContext, path string) error
	GetURL(ctx *appcontext.AppContext, path string) (string, error)
}
