package gobbs

// Stream wrapps a Generator and provides a channel, which will be fed with the generators output.
type Stream struct {
	g *Generator
	C chan uint
}

// Start starts a goroutine, which runs the generator and pipes its output to channel C.
func (s *Stream) Start() *Stream {

	s.C = make(chan uint, 512)
	go s.g.bytesLoop(s.C)
	return s
}
