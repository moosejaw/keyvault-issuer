package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// ChainCertificateData defines a source of certificate
// data to be included in a ChainCertificate in order
// to form a completed certificate chain.
// Valid sources of certificate data include a ConfigMap
// or a Secret.
type ChainCertificateData struct {
	// ConfigMapKeyRef is a reference to an existing key
	// within a ConfigMap.
	// +optional
	ConfigMapKeyRef *v1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`

	// SecretKeyRef is a reference to a key within an
	// existing secret in the KeyVaultIssuer's namespace.
	// +optional
	SecretKeyRef *v1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// ChainCertificate defines a certificate to be included in a
// fully-formed certificate chain to be prepended to the issued
// certificate's data.
type ChainCertificate struct {
	// Number is the certificate's position in the chain.
	// The chain is formed automatically by taking each
	// certificate in ascending order, so the value assigned
	// to Number is inconsequential, provided, for example,
	// the root certificate's Number is smaller than the
	// intermediate certificate's Number.
	Number uint `json:"number"`

	// CertificateData is a reference to a source of public
	// certificate data to be used when forming the full
	// certificate chain for issued certificates.
	CertificateData ChainCertificateData `json:"chainCertificateData"`
}

// KeyVaultIssuerSpec defines the desired state of KeyVaultIssuer
type KeyVaultIssuerSpec struct {
	// KeyID is the URI of the key used to sign CertificateRequests.
	// The URI should not contain the extension for the key version,
	// as the issuer will always default to using the latest version
	// of the key to sign requests.
	KeyID string `json:"keyId"`

	// ServicePrincipalCredential is a reference to a Kubernetes
	// secret within the cluster that contains the required fields
	// when authenticating via a Service Principal in Azure.
	// +optional
	ServicePrincipalCredential *v1.SecretReference `json:"credential,omitempty"`

	// CertificateChain is a representation the other certificates
	// in the full chain.
	// +optional
	CertificateChain *[]ChainCertificate `json:"certificateChain,omitempty"`
}
