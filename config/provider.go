
package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/lschulte19/crossplane-test/config/objectstorage"

)

const (
	resourcePrefix = "oci"
	modulePath     = "github.com/oracle/provider-oci"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		objectstorage.Configure,
		identity.Configure,
		core.Configure,
		kms.Configure,
		containerengine.Configure,
		artifacts.Configure,
		ons.Configure,
		networkloadbalancer.Configure,
		dns.Configure,
		healthchecks.Configure,
		functions.Configure,
		networkfirewall.Configure,
		logging.Configure,
		monitoring.Configure,
		loadbalancer.Configure,
		servicemesh.Configure,
		certificatesmanagement.Configure,
		filestorage.Configure,
		events.Configure,
		vault.Configure,
		events.Configure,
		streaming.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}