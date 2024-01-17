package meinstallationtemplatepartitionschemehardwareraid

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("ovh_me_installation_template_partition_scheme_hardware_raid", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "meinstallationtemplatepartitionschemehardwareraid"
    })
}
