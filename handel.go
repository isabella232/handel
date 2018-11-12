package handel

import "sync"

// Handel is the principal struct that performs the large scale multi-signature
// aggregation protocol. Handel is thread-safe.
type Handel struct {
	sync.Mutex
	// Config holding parameters to Handel
	c *Config
	// Network enabling external communication with other Handel nodes
	net Network
	// Registry holding access to all Handel node's identities
	reg Registry
	// signature scheme used for this Handel protocol
	scheme SignatureScheme
	// Message that is being signed during the Handel protocol
	msg []byte
	// increamental aggregated signature cached by this handel node
	aggregate Signature
	// channel to exposes multi-signatures to the user
	out chan MultiSignature
}

// NewHandel returns a Handle interface that uses the given network and
// registry. The signature scheme is the one to use for this Handel protocol,
// and the message is the message to multi-sign.The first config in the slice is
// taken if not nil. Otherwise, the default config generated by DefaultConfig()
// is used.
func NewHandel(n Network, r Registry, s SignatureScheme, msg []byte,
	conf ...*Config) (*Handel, error) {
	h := &Handel{
		net:    n,
		reg:    r,
		scheme: s,
		msg:    msg,
	}

	if len(conf) > 0 && conf[0] != nil {
		h.c = mergeWithDefault(conf[0], r.Size())
	} else {
		h.c = DefaultConfig(r.Size())
	}

	ms, err := s.Sign(msg, nil)
	if err != nil {
		return nil, err
	}
	h.aggregate = ms
	return h, nil
}

// NewPacket implements the Listener interface for the network.
func (h *Handel) NewPacket(p *Packet) {
	h.Lock()
	defer h.Unlock()
}

// Start the Handel protocol
func (h *Handel) Start() {
	h.Lock()
	defer h.Unlock()
}
