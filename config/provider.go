
package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/lschulte19/crossplane-test/config/artifacts"
	"github.com/lschulte19/crossplane-test/config/certificatesmanagement"
	"github.com/lschulte19/crossplane-test/config/containerengine"
	"github.com/lschulte19/crossplane-test/config/core"
	"github.com/lschulte19/crossplane-test/config/dns"
	"github.com/lschulte19/crossplane-test/config/events"
	"github.com/lschulte19/crossplane-test/config/filestorage"
	"github.com/lschulte19/crossplane-test/config/functions"
	"github.com/lschulte19/crossplane-test/config/healthchecks"
	"github.com/lschulte19/crossplane-test/config/identity"
	"github.com/lschulte19/crossplane-test/config/kms"
	"github.com/lschulte19/crossplane-test/config/loadbalancer"
	"github.com/lschulte19/crossplane-test/config/logging"
	"github.com/lschulte19/crossplane-test/config/monitoring"
	"github.com/lschulte19/crossplane-test/config/networkfirewall"
	"github.com/lschulte19/crossplane-test/config/networkloadbalancer"
	"github.com/lschulte19/crossplane-test/config/objectstorage"
	"github.com/lschulte19/crossplane-test/config/ons"
	"github.com/lschulte19/crossplane-test/config/servicemesh"
	"github.com/lschulte19/crossplane-test/config/streaming"
	"github.com/lschulte19/crossplane-test/config/vault"
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