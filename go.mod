module github.com/lucas-clemente/quic-go

go 1.16

replace golang.org/x/crypto => github.com/aperturerobotics/golang-x-crypto v0.0.0-20210728080812-5cca7994ea9f // gopherjs-compat

replace github.com/marten-seemann/qtls-go1-16 => github.com/paralin/qtls-go1-16 v0.1.5-0.20210728071944-419a2c247411 // gopherjs-compat

require (
	github.com/cheekybits/genny v1.0.0
	github.com/marten-seemann/qpack v0.2.1
	github.com/marten-seemann/qtls-go1-16 v0.1.4
	github.com/marten-seemann/qtls-go1-17 v0.1.0
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20210428140749-89ef3d95e781
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b
)
