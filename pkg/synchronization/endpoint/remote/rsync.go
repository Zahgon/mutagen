package remote

import (
	"github.com/mutagen-io/mutagen/pkg/encoding"
	"github.com/mutagen-io/mutagen/pkg/stream"
	"github.com/mutagen-io/mutagen/pkg/synchronization/rsync"
)

// protobufRsyncEncoder implements rsync.Encoder using Protocol Buffers.
type protobufRsyncEncoder struct {
	// encoder is the underlying Protocol Buffers encoder.
	encoder *encoding.ProtobufEncoder
	// flusher flushes the underlying stream.
	flusher stream.Flusher
	// error stores any previously encountered transmission error.
	error error
}

// Encode implements rsync.Encoder.Encode.
func (e *protobufRsyncEncoder) Encode(transmission *rsync.Transmission) error {
	_ = "STUB: not implemented"
	// Check for previous errors.
	return nil
}

// Encode the transmission.

// Finalize implements rsync.Encoder.Finalize.
func (e *protobufRsyncEncoder) Finalize() error {
	_ = "STUB: not implemented"
	// If an error has occurred, then there's nothing to do.
	return nil
}

// Otherwise, attempt to flush the compressor.

// Success.

// protobufRsyncDecoder implements rsync.Decoder using Protocol Buffers.
type protobufRsyncDecoder struct {
	// decoder is the underlying Protocol Buffers decoder.
	decoder *encoding.ProtobufDecoder
}

// Decode implements rsync.Decoder.Decode.
func (d *protobufRsyncDecoder) Decode(transmission *rsync.Transmission) error {
	_ = "STUB: not implemented"
	// TODO: This is not particularly efficient because the Protocol Buffers
	// decoding implementation doesn't reuse existing capacity in operation data
	// buffers. This is something that needs to be fixed upstream, but we should
	// file an issue. Once it's done, nothing on our end needs to change except
	// to update the Protocol Buffers runtime.
	return nil
}

// Finalize implements rsync.Decoder.Finalize.
func (d *protobufRsyncDecoder) Finalize() error { _ = "STUB: not implemented"; return nil }
