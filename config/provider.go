/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

        "github.com/AitorLeon89/provider-ovh/config/cloudproject"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectuser"
        "github.com/AitorLeon89/provider-ovh/config/meidentityuser"
        "github.com/AitorLeon89/provider-ovh/config/meidentitygroup"
        "github.com/AitorLeon89/provider-ovh/config/messhkey"
        "github.com/AitorLeon89/provider-ovh/config/vrackdedicatedserver"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectusers3credential"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectusers3policy"
        "github.com/AitorLeon89/provider-ovh/config/iampolicy"
        "github.com/AitorLeon89/provider-ovh/config/ipservice"
        "github.com/AitorLeon89/provider-ovh/config/ipreverse"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectfailoveripattach"
        "github.com/AitorLeon89/provider-ovh/config/dedicatedcephacl"
        "github.com/AitorLeon89/provider-ovh/config/meipxescript"
        "github.com/AitorLeon89/provider-ovh/config/domainzone"
        "github.com/AitorLeon89/provider-ovh/config/domainzoneredirection"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancing"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancingvracknetwork"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancingrefresh"
        "github.com/AitorLeon89/provider-ovh/config/iamresourcegroup"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectworkflowbackup"
        "github.com/AitorLeon89/provider-ovh/config/vrackdedicatedserverinterface"
        "github.com/AitorLeon89/provider-ovh/config/vrackip"
        "github.com/AitorLeon89/provider-ovh/config/vrackiploadbalancing"
        "github.com/AitorLeon89/provider-ovh/config/dbaaslogsinput"
        "github.com/AitorLeon89/provider-ovh/config/dbaaslogscluster"
        "github.com/AitorLeon89/provider-ovh/config/dbaaslogsgraylogoutputstream"
        "github.com/AitorLeon89/provider-ovh/config/dedicatednashapartition"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectregionstoragepresign"
        "github.com/AitorLeon89/provider-ovh/config/hostingprivatedatabase"
        "github.com/AitorLeon89/provider-ovh/config/hostingprivatedatabasedatabase"
        "github.com/AitorLeon89/provider-ovh/config/hostingprivatedatabaseuser"
        "github.com/AitorLeon89/provider-ovh/config/hostingprivatedatabaseusergrant"
        "github.com/AitorLeon89/provider-ovh/config/meapioauth2client"
        "github.com/AitorLeon89/provider-ovh/config/dedicatedserverinstalltask"
        "github.com/AitorLeon89/provider-ovh/config/dedicatedserverreboottask"
        "github.com/AitorLeon89/provider-ovh/config/dedicatedserverupdate"
        "github.com/AitorLeon89/provider-ovh/config/meinstallationtemplate"
        "github.com/AitorLeon89/provider-ovh/config/meinstallationtemplatepartitionscheme"
        "github.com/AitorLeon89/provider-ovh/config/meinstallationtemplatepartitionschemepartition"
        "github.com/AitorLeon89/provider-ovh/config/meinstallationtemplatepartitionschemehardwareraid"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancinghttpfarm"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancinghttpfarmserver"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancinghttpfrontend"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancinghttproute"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancinghttprouterule"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancingtcpfarm"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancingtcpfarmserver"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancingtcpfrontend"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancingtcproute"
        "github.com/AitorLeon89/provider-ovh/config/iploadbalancingtcprouterule"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectkube"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectkubeiprestrictions"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectkubenodepool"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectkubeoidc"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectcontainerregistry"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectcontainerregistryoidc"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectcontainerregistryuser"
        "github.com/AitorLeon89/provider-ovh/config/dedicatednashapartitionaccess"
        "github.com/AitorLeon89/provider-ovh/config/dedicatednashapartitionsnapshot"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectnetworkprivate"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectnetworkprivatesubnet"
        "github.com/AitorLeon89/provider-ovh/config/hostingprivatedatabasewhitelist"
        "github.com/AitorLeon89/provider-ovh/config/vrack"
        "github.com/AitorLeon89/provider-ovh/config/vrackcloudproject"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabase"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabasedatabase"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabaseintegration"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabaseiprestriction"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabasekafkaacl"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabasekafkaschemaregistryacl"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabasekafkatopic"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabasem3dbnamespace"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabasem3dbuser"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabasemongodbuser"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabaseopensearchpattern"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabaseopensearchuser"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabasepostgresqluser"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabaseredisuser"
        "github.com/AitorLeon89/provider-ovh/config/cloudprojectdatabaseuser"
)




const (
	resourcePrefix = "ovh"
	modulePath     = "github.com/AitorLeon89/provider-ovh"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
                cloudproject.Configure,
                cloudprojectuser.Configure,
                cloudprojectusers3credential.Configure,
                cloudprojectusers3policy.Configure,
                meidentityuser.Configure,
                meidentitygroup.Configure,
                messhkey.Configure,
                vrackdedicatedserver.Configure,
                iampolicy.Configure,
                ipservice.Configure,
                ipreverse.Configure,
                cloudprojectfailoveripattach.Configure,
                dedicatedcephacl.Configure,
                meipxescript.Configure,
                domainzone.Configure,
                domainzoneredirection.Configure,
                iploadbalancing.Configure,
                iploadbalancingvracknetwork.Configure,
                iploadbalancingrefresh.Configure,
                dedicatedcephacl.Configure,
                iamresourcegroup.Configure,
                cloudprojectworkflowbackup.Configure,
                vrackdedicatedserverinterface.Configure,
                vrackip.Configure,
                vrackiploadbalancing.Configure,
                dbaaslogsinput.Configure,
                dbaaslogscluster.Configure,
                dbaaslogsgraylogoutputstream.Configure,
                dedicatednashapartition.Configure,
                cloudprojectregionstoragepresign.Configure,
                hostingprivatedatabase.Configure,
                hostingprivatedatabasedatabase.Configure,
                hostingprivatedatabaseuser.Configure,
                hostingprivatedatabaseusergrant.Configure,
                meapioauth2client.Configure,
                dedicatedserverinstalltask.Configure,
                dedicatedserverreboottask.Configure,
                dedicatedserverupdate.Configure,
                meinstallationtemplate.Configure,
                meinstallationtemplatepartitionscheme.Configure,
                meinstallationtemplatepartitionschemepartition.Configure,
                meinstallationtemplatepartitionschemehardwareraid.Configure,
                iploadbalancinghttpfarm.Configure,
                iploadbalancinghttpfarmserver.Configure,
                iploadbalancinghttpfrontend.Configure,
                iploadbalancinghttproute.Configure,
                iploadbalancinghttprouterule.Configure,
                iploadbalancingtcpfarm.Configure,
                iploadbalancingtcpfarmserver.Configure,
                iploadbalancingtcpfrontend.Configure,
                iploadbalancingtcproute.Configure,
                iploadbalancingtcprouterule.Configure,
                cloudprojectkube.Configure,
                cloudprojectkubeiprestrictions.Configure,
                cloudprojectkubenodepool.Configure,
                cloudprojectkubeoidc.Configure,
                cloudprojectcontainerregistry.Configure,
                cloudprojectcontainerregistryoidc.Configure,
                cloudprojectcontainerregistryuser.Configure,
                dedicatednashapartitionaccess.Configure,
                dedicatednashapartitionsnapshot.Configure,
                cloudprojectnetworkprivate.Configure,
                cloudprojectnetworkprivatesubnet.Configure,
                hostingprivatedatabasewhitelist.Configure,
                vrack.Configure,
                vrackcloudproject.Configure,
                cloudprojectdatabase.Configure,
                cloudprojectdatabasedatabase.Configure,
                cloudprojectdatabaseintegration.Configure,
                cloudprojectdatabaseiprestriction.Configure,
                cloudprojectdatabasekafkaacl.Configure,
                cloudprojectdatabasekafkaschemaregistryacl.Configure,
                cloudprojectdatabasekafkatopic.Configure,
                cloudprojectdatabasem3dbnamespace.Configure,
                cloudprojectdatabasem3dbuser.Configure,
                cloudprojectdatabasemongodbuser.Configure,
                cloudprojectdatabaseopensearchpattern.Configure,
                cloudprojectdatabaseopensearchuser.Configure,
                cloudprojectdatabasepostgresqluser.Configure,
                cloudprojectdatabaseredisuser.Configure,
                cloudprojectdatabaseuser.Configure,
 	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
