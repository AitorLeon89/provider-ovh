/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
        "ovh_cloud_project": config.IdentifierFromProvider,
        "ovh_cloud_project_user": config.IdentifierFromProvider,
        "ovh_cloud_project_user_s3_credential": config.IdentifierFromProvider,
        "ovh_cloud_project_user_s3_policy": config.IdentifierFromProvider,
        "ovh_me_identity_user": config.IdentifierFromProvider,
        "ovh_me_api_oauth2_client": config.ParameterAsIdentifier("client_id"),
        "ovh_me_identity_group": config.IdentifierFromProvider,
        "ovh_me_ssh_key": config.IdentifierFromProvider,
        "ovh_vrack_dedicated_server": config.IdentifierFromProvider,
        "ovh_iam_policy": config.IdentifierFromProvider,
        "ovh_ip_service": config.IdentifierFromProvider,
        "ovh_ip_reverse": config.IdentifierFromProvider,
        "ovh_cloud_project_failover_ip_attach": config.IdentifierFromProvider,
        "ovh_dedicated_ceph_acl": config.IdentifierFromProvider,
        "ovh_me_ipxe_script": config.IdentifierFromProvider,
        "ovh_domain_zone": config.IdentifierFromProvider,
        "ovh_domain_zone_redirection": config.IdentifierFromProvider,
        "ovh_iploadbalancing_vrack_network": config.IdentifierFromProvider,
        "ovh_iploadbalancing_refresh": config.IdentifierFromProvider,
        "ovh_iam_resource_group": config.IdentifierFromProvider,
	"ovh_cloud_project_workflow_backup": config.IdentifierFromProvider,
        "ovh_vrack_dedicated_server_interface": config.IdentifierFromProvider,
        "ovh_vrack_ip": config.IdentifierFromProvider,
        "ovh_vrack_iploadbalancing": config.IdentifierFromProvider,
        "ovh_dbaas_logs_input": config.IdentifierFromProvider,
        "ovh_dbaas_logs_cluster": config.TemplatedStringAsIdentifier("cluster_id","{{ .parameter.service_name }}/{{ .externalName }}"),
        "ovh_dbaas_logs_graylog_output_stream": config.IdentifierFromProvider,
        "ovh_dedicated_nasha_partition": config.TemplatedStringAsIdentifier("name","{{ .parameter.service_name }}/{{ .externalName }}"),
        "ovh_cloud_project_region_storage_presign": config.IdentifierFromProvider,
        "ovh_hosting_private_database": config.ParameterAsIdentifier("service_name"),
        "ovh_hosting_privatedatabase_database": config.TemplatedStringAsIdentifier("database_name","{{ .parameter.service_name }}/{{ .externalName }}"),
        "ovh_hosting_privatedatabase_user": config.TemplatedStringAsIdentifier("user_name","{{ .parameter.service_name }}/{{ .externalName }}"),
        "ovh_hosting_privatedatabase_user_grant": config.IdentifierFromProvider,
        "ovh_hosting_privatedatabase_whitelist": config.IdentifierFromProvider,
        "ovh_dedicated_server_install_task": config.IdentifierFromProvider,
        "ovh_dedicated_server_reboot_task": config.IdentifierFromProvider,
        "ovh_dedicated_server_update": config.IdentifierFromProvider,
        "ovh_me_installation_template": config.TemplatedStringAsIdentifier("template_name","{{ .parameter.base_template_name }}/{{ .externalName }}"),
        "ovh_me_installation_template_partition_scheme": config.ParameterAsIdentifier("id"),
        "ovh_me_installation_template_partition_scheme_hardware_raid": config.ParameterAsIdentifier("id"),
        "ovh_me_installation_template_partition_scheme_partition": config.ParameterAsIdentifier("id"),
        "ovh_iploadbalancing": config.IdentifierFromProvider,
        "ovh_iploadbalancing_http_farm": config.IdentifierFromProvider,
        "ovh_iploadbalancing_http_farm_server": config.IdentifierFromProvider,
        "ovh_iploadbalancing_http_frontend": config.IdentifierFromProvider,
        "ovh_iploadbalancing_http_route": config.IdentifierFromProvider,
        "ovh_iploadbalancing_http_route_rule": config.IdentifierFromProvider,
        "ovh_iploadbalancing_tcp_farm": config.IdentifierFromProvider,
        "ovh_iploadbalancing_tcp_farm_server": config.IdentifierFromProvider,
        "ovh_iploadbalancing_tcp_frontend": config.IdentifierFromProvider,
        "ovh_iploadbalancing_tcp_route": config.IdentifierFromProvider,
        "ovh_iploadbalancing_tcp_route_rule": config.IdentifierFromProvider,
        "ovh_cloud_project_kube": config.IdentifierFromProvider,
        "ovh_cloud_project_kube_iprestrictions": config.IdentifierFromProvider,
        "ovh_cloud_project_kube_nodepool": config.IdentifierFromProvider,
        "ovh_cloud_project_kube_oidc": config.IdentifierFromProvider,
        "ovh_cloud_project_containerregistry": config.IdentifierFromProvider,
        "ovh_cloud_project_containerregistry_oidc": config.TemplatedStringAsIdentifier("registry_id","{{ .parameter.service_name }}/{{ .externalName }}"),
        "ovh_cloud_project_containerregistry_user": config.IdentifierFromProvider,
        "ovh_dedicated_nasha_partition_access": config.IdentifierFromProvider,
        "ovh_dedicated_nasha_partition_snapshot": config.IdentifierFromProvider,
        "ovh_cloud_project_network_private": config.IdentifierFromProvider,
        "ovh_cloud_project_network_private_subnet": config.IdentifierFromProvider,
        "ovh_vrack": config.ParameterAsIdentifier("order_id"),
        "ovh_vrack_cloudproject": config.IdentifierFromProvider,
        "ovh_cloud_project_database": config.IdentifierFromProvider,
        "ovh_cloud_project_database_database": config.IdentifierFromProvider,
        "ovh_cloud_project_database_integration": config.IdentifierFromProvider,
        "ovh_cloud_project_database_ip_restriction": config.IdentifierFromProvider,
        "ovh_cloud_project_database_kafka_acl": config.IdentifierFromProvider,
        "ovh_cloud_project_database_kafka_schemaregistryacl": config.IdentifierFromProvider,
        "ovh_cloud_project_database_kafka_topic": config.IdentifierFromProvider,
        "ovh_cloud_project_database_m3db_namespace": config.IdentifierFromProvider,
        "ovh_cloud_project_database_m3db_user": config.IdentifierFromProvider,
        "ovh_cloud_project_database_mongodb_user": config.IdentifierFromProvider,
        "ovh_cloud_project_database_opensearch_pattern": config.IdentifierFromProvider,
        "ovh_cloud_project_database_opensearch_user": config.IdentifierFromProvider,
        "ovh_cloud_project_database_postgresql_user": config.IdentifierFromProvider,
        "ovh_cloud_project_database_redis_user": config.IdentifierFromProvider,
        "ovh_cloud_project_database_user": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
