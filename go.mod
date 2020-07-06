module github.com/lucas-clemente/quic-go

go 1.14

// aperture: use 1.3.x based fork for compatibility
replace github.com/golang/protobuf => github.com/aperturerobotics/go-protobuf-1.3.x v0.0.0-20200706003739-05fb54d407a9 // aperture-1.3.x

require (
	github.com/cheekybits/genny v1.0.0
	github.com/francoispqt/gojay v1.2.13
	github.com/golang/mock v1.4.0
	github.com/golang/protobuf v1.4.0
	github.com/marten-seemann/qpack v0.1.0
	github.com/marten-seemann/qtls v0.10.0
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	golang.org/x/crypto v0.0.0-20200423211502-4bdfaf469ed5
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/protobuf v1.23.0
)
