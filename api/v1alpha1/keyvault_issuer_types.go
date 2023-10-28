package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// KeyVaultIssuerSpec defines the desired state of KeyVaultIssuer
type KeyVaultIssuerSpec struct {
	// CertificateID is the URI of the certificate used to sign CertificateRequests.
	// The URI should not contain the extension for the certificate version,
	// as the issuer will always default to using the latest version
	// of the key to sign requests.
	CertificateID string `json:"certificateId"`

	// ServicePrincipalClientSecret is a reference to a Kubernetes
	// secret within the cluster that contains the client secret
	// when authenticating via a Service Principal in Azure.
	// +optional
	ServicePrincipalClientSecret *v1.SecretReference `json:"credential,omitempty"`

	// ServicePrincipalClientID is the client ID of the Entra ID service
	// principal used for authenticating requests against key vault.
	// +optional
	ServicePrincipalClientID string `json:"servicePrincipalClientId"`

	// ServicePrincipalTenantID is the tenant ID of the Entra ID tenant
	// in which the service principal exists.
	// +optional
	ServicePrincipalTenantID string `json:"servicePrincipalTenantId"`
}
