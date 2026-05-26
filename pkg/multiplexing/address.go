package multiplexing

// multiplexerAddress implements net.Addr for Multiplexer.
type multiplexerAddress struct {
	// even indicates whether or not this is the even-valued multiplexer.
	even bool
}

// Network implements net.Addr.Network.
func (a *multiplexerAddress) Network() string { _ = "STUB: not implemented"; return "" }

// String implements net.Addr.String.
func (a *multiplexerAddress) String() string { _ = "STUB: not implemented"; return "" }

// streamAddress implements net.Addr for Stream.
type streamAddress struct {
	// remote indicates whether or not the address is remote.
	remote bool
	// identifier is the stream identifier.
	identifier uint64
}

// Network implements net.Addr.Network.
func (a *streamAddress) Network() string { _ = "STUB: not implemented"; return "" }

// String implements net.Addr.String.
func (a *streamAddress) String() string { _ = "STUB: not implemented"; return "" }
