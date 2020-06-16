package catalog

import (
	"github.com/open-service-mesh/osm/pkg/certificate"
	"github.com/open-service-mesh/osm/pkg/service"
)

// GetCertificateForService returns the certificate the given proxy uses for mTLS to the XDS server.
func (mc *MeshCatalog) GetCertificateForService(nsService service.NamespacedService) (certificate.Certificater, error) {
	cn := nsService.GetCommonName()

	cert, err := mc.certManager.GetCertificate(cn)
	if err != nil {
		// Certificate was not found in CertManager's cache, issue one
		newCert, err := mc.certManager.IssueCertificate(nsService.GetCommonName(), nil)
		if err != nil {
			log.Error().Err(err).Msgf("Error issuing a new certificate for service:%s, CN: %s", nsService, cn)
			return nil, err
		}
		return newCert, nil
	}
	return cert, nil
}
