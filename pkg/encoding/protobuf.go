package encoding

import (
	"io"

	"google.golang.org/protobuf/proto"

	"github.com/mutagen-io/mutagen/pkg/stream"
)

const (
	// protobufEncoderInitialBufferSize is the initial buffer size for encoders.
	protobufEncoderInitialBufferSize = 32 * 1024

	// protobufEncoderMaximumPersistentBufferSize is the maximum buffer size
	// that the encoder will keep allocated.
	protobufEncoderMaximumPersistentBufferSize = 1024 * 1024

	// protobufDecoderInitialBufferSize is the initial buffer size for decoders.
	protobufDecoderInitialBufferSize = 32 * 1024

	// protobufDecoderMaximumAllowedMessageSize is the maximum message size that
	// we'll attempt to read from the wire.
	protobufDecoderMaximumAllowedMessageSize = 100 * 1024 * 1024

	// protobufDecoderMaximumPersistentBufferSize is the maximum buffer size
	// that the decoder will keep allocated.
	protobufDecoderMaximumPersistentBufferSize = 1024 * 1024
)

// LoadAndUnmarshalProtobuf loads data from the specified path and decodes it
// into the specified Protocol Buffers message.
func LoadAndUnmarshalProtobuf(path string, message proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// MarshalAndSaveProtobuf marshals the specified Protocol Buffers message and
// saves it to the specified path.
func MarshalAndSaveProtobuf(path string, message proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// ProtobufEncoder is a stream encoder for Protocol Buffers messages.
type ProtobufEncoder struct {
	// writer is the underlying writer.
	writer io.Writer
	// buffer is a reusable encoding buffer.
	buffer []byte
	// sizer is a Protocol Buffers marshaling configuration for computing sizes.
	sizer proto.MarshalOptions
	// encoder is a Protocol Buffers marshaling configuration for encoding.
	encoder proto.MarshalOptions
}

// NewProtobufEncoder creates a new Protocol Buffers stream encoder.
func NewProtobufEncoder(writer io.Writer) *ProtobufEncoder { _ = "STUB: not implemented"; return nil }

// Encode encodes and writes a length-prefixed Protocol Buffers message to the
// underlying stream. If this fails, the encoder should be considered corrupted.
func (e *ProtobufEncoder) Encode(message proto.Message) error {
	_ = "STUB: not implemented"
	// Always make sure that the buffer's capacity stays within the limit of
	// what we're willing to carry around once we're done.
	return nil
}

// Encode the message size.

// Encode the message.

// Write the data to the wire.

// Success.

// ProtobufDecoder is a stream decoder for Protocol Buffers messages.
type ProtobufDecoder struct {
	// reader is the underlying reader.
	reader stream.DualModeReader
	// buffer is a reusable receive buffer for decoding messages.
	buffer []byte
}

// NewProtobufDecoder creates a new Protocol Buffers stream decoder.
func NewProtobufDecoder(reader stream.DualModeReader) *ProtobufDecoder {
	_ = "STUB: not implemented"
	return nil
}

// bufferWithSize returns a buffer with the specified size, opting to reuse a
// cached buffer if possible.
func (d *ProtobufDecoder) bufferWithSize(size int) []byte {
	_ = "STUB: not implemented"
	// If we can satisfy this request with our existing buffer, then use that.
	return nil
}

// Otherwise allocate a new buffer.

// If this buffer doesn't exceed the maximum size that we're willing to keep
// around in memory, then store it.

// Done.

// Decode decodes a length-prefixed Protocol Buffers message from the underlying
// stream. If this fails, the decoder should be considered corrupted.
func (d *ProtobufDecoder) Decode(message proto.Message) error {
	_ = "STUB: not implemented"
	// Read the next message length.
	return nil
}

// Check if the message is too long to read.

// Grab a buffer to read the message.

// Read the message bytes.

// Unmarshal the message.

// Success.

// EncodeProtobuf encodes a single Protocol Buffers message that can be read by
// ProtobufDecoder or DecodeProtobuf. It is a useful shorthand for creating a
// ProtobufEncoder and writing a single message. For multiple message sends, it
// is far more efficient to use a ProtobufEncoder directly and repeatedly.
func EncodeProtobuf(writer io.Writer, message proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// DecodeProtobuf reads and decodes a single Protocol Buffers message as written
// by ProtobufEncoder or EncodeProtobuf. It is a useful shorthand for creating a
// ProtobufDecoder and reading a single message. For multiple message reads, it
// is far more efficient to use a ProtobufDecoder directly and repeatedly.
func DecodeProtobuf(reader stream.DualModeReader, message proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}
