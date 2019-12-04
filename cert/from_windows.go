package cert

import (
	"crypto/x509"
	"log"
	"syscall"
	"unsafe"
)

// AppendFromSystem builds a CA pool with all certificates known by the host Windows operating system.
func (pool Pool) AppendFromSystem() error {
	storeHandle, err := syscall.CertOpenSystemStore(0, syscall.StringToUTF16Ptr("Root"))
	if err != nil {
		return err
	}

	var cert *syscall.CertContext

	for {
		cert, err = syscall.CertEnumCertificatesInStore(storeHandle, cert)
		if err != nil {
			if errno, ok := err.(syscall.Errno); ok {
				if errno == 0x80092004 {
					return nil
				}
			}
			log.Println(syscall.GetLastError())
		}
		if cert == nil {
			return nil
		}

		buf := (*[1 << 20]byte)(unsafe.Pointer(cert.EncodedCert))[:]
		buf2 := make([]byte, cert.Length)
		copy(buf2, buf)
		if c, err := x509.ParseCertificate(buf2); err == nil {
			pool.Certs.AddCert(c)
		}
	}
}
