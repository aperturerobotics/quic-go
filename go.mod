module github.com/lucas-clemente/quic-go

go 1.15

// aperture: use 1.3.x based fork for compatibility
replace github.com/golang/protobuf => github.com/aperturerobotics/go-protobuf-1.3.x v0.0.0-20200726220404-fa7f51c52df0 // aperture-1.3.x

require (
	github.com/cheekybits/genny v1.0.0
	github.com/francoispqt/gojay v1.2.13
	github.com/golang/mock v1.5.0
	github.com/marten-seemann/qpack v0.2.1
	github.com/marten-seemann/qtls-go1-15 v0.1.4
	github.com/marten-seemann/qtls-go1-16 v0.1.3
	github.com/marten-seemann/qtls-go1-17 v0.1.0-alpha.1
	github.com/onsi/ginkgo v1.16.2
	github.com/onsi/gomega v1.12.0
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
