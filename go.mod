module github.com/lucas-clemente/quic-go

go 1.14

// aperture: use 1.3.x based fork for compatibility
replace github.com/golang/protobuf => github.com/aperturerobotics/go-protobuf-1.3.x v0.0.0-20200726220404-fa7f51c52df0 // aperture-1.3.x

require (
	github.com/cheekybits/genny v1.0.0
	github.com/francoispqt/gojay v1.2.13
	github.com/golang/mock v1.5.0
	github.com/marten-seemann/qpack v0.2.1
	github.com/marten-seemann/qtls-go1-15 v0.1.4
	github.com/marten-seemann/qtls-go1-16 v0.1.3
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc
	golang.org/x/net v0.0.0-20210420210106-798c2154c571
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210420072515-93ed5bcd2bfe
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
