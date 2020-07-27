module github.com/lucas-clemente/quic-go

go 1.16

replace golang.org/x/crypto => github.com/aperturerobotics/golang-x-crypto v0.0.0-20220321111526-87c0d0398f72 // gopherjs-compat

require (
	github.com/cheekybits/genny v1.0.0
	github.com/marten-seemann/qpack v0.2.1
	github.com/marten-seemann/qtls-go1-16 v0.1.5
	github.com/marten-seemann/qtls-go1-17 v0.1.1
	github.com/marten-seemann/qtls-go1-18 v0.1.1
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8
)
