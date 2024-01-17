// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	project "github.com/AitorLeon89/provider-ovh/internal/controller/cloudproject/project"
	projectcontainerregistry "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectcontainerregistry/projectcontainerregistry"
	projectcontainerregistryoidc "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectcontainerregistryoidc/projectcontainerregistryoidc"
	projectcontainerregistryuser "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectcontainerregistryuser/projectcontainerregistryuser"
	projectdatabase "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabase/projectdatabase"
	projectdatabaseuser "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabase_user/projectdatabaseuser"
	projectdatabasedatabase "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabasedatabase/projectdatabasedatabase"
	projectdatabaseintegration "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabaseintegration/projectdatabaseintegration"
	projectdatabaseiprestriction "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabaseiprestriction/projectdatabaseiprestriction"
	projectdatabasekafkaacl "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabasekafkaacl/projectdatabasekafkaacl"
	projectdatabasekafkaschemaregistryacl "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabasekafkaschemaregistryacl/projectdatabasekafkaschemaregistryacl"
	projectdatabasekafkatopic "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabasekafkatopic/projectdatabasekafkatopic"
	projectdatabasem3dbnamespace "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabasem3dbnamespace/projectdatabasem3dbnamespace"
	projectdatabasem3dbuser "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabasem3dbuser/projectdatabasem3dbuser"
	projectdatabasemongodbuser "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabasemongodbuser/projectdatabasemongodbuser"
	projectdatabaseopensearchpattern "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabaseopensearchpattern/projectdatabaseopensearchpattern"
	projectdatabaseopensearchuser "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabaseopensearchuser/projectdatabaseopensearchuser"
	projectdatabasepostgresqluser "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabasepostgresqluser/projectdatabasepostgresqluser"
	projectdatabaseredisuser "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectdatabaseredisuser/projectdatabaseredisuser"
	projectfailoveripattach "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectfailoveripattach/projectfailoveripattach"
	projectkube "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectkube/projectkube"
	projectkubeiprestrictions "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectkubeiprestrictions/projectkubeiprestrictions"
	projectkubenodepool "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectkubenodepool/projectkubenodepool"
	projectkubeoidc "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectkubeoidc/projectkubeoidc"
	projectnetworkprivate "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectnetworkprivate/projectnetworkprivate"
	projectnetworkprivatesubnet "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectnetworkprivatesubnet/projectnetworkprivatesubnet"
	projectregionstoragepresign "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectregionstoragepresign/projectregionstoragepresign"
	projectuser "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectuser/projectuser"
	projectusers3credential "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectusers3credential/projectusers3credential"
	projectusers3policy "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectusers3policy/projectusers3policy"
	projectworkflowbackup "github.com/AitorLeon89/provider-ovh/internal/controller/cloudprojectworkflowbackup/projectworkflowbackup"
	logscluster "github.com/AitorLeon89/provider-ovh/internal/controller/dbaaslogscluster/logscluster"
	logsinput "github.com/AitorLeon89/provider-ovh/internal/controller/dbaaslogsinput/logsinput"
	cephacl "github.com/AitorLeon89/provider-ovh/internal/controller/dedicatedcephacl/cephacl"
	nashapartition "github.com/AitorLeon89/provider-ovh/internal/controller/dedicatednashapartition/nashapartition"
	nashapartitionaccess "github.com/AitorLeon89/provider-ovh/internal/controller/dedicatednashapartitionaccess/nashapartitionaccess"
	nashapartitionsnapshot "github.com/AitorLeon89/provider-ovh/internal/controller/dedicatednashapartitionsnapshot/nashapartitionsnapshot"
	serverinstalltask "github.com/AitorLeon89/provider-ovh/internal/controller/dedicatedserverinstalltask/serverinstalltask"
	serverreboottask "github.com/AitorLeon89/provider-ovh/internal/controller/dedicatedserverreboottask/serverreboottask"
	serverupdate "github.com/AitorLeon89/provider-ovh/internal/controller/dedicatedserverupdate/serverupdate"
	zone "github.com/AitorLeon89/provider-ovh/internal/controller/domainzone/zone"
	zoneredirection "github.com/AitorLeon89/provider-ovh/internal/controller/domainzoneredirection/zoneredirection"
	privatedatabasedatabase "github.com/AitorLeon89/provider-ovh/internal/controller/hostingprivatedatabasedatabase/privatedatabasedatabase"
	privatedatabaseuser "github.com/AitorLeon89/provider-ovh/internal/controller/hostingprivatedatabaseuser/privatedatabaseuser"
	privatedatabaseusergrant "github.com/AitorLeon89/provider-ovh/internal/controller/hostingprivatedatabaseusergrant/privatedatabaseusergrant"
	privatedatabasewhitelist "github.com/AitorLeon89/provider-ovh/internal/controller/hostingprivatedatabasewhitelist/privatedatabasewhitelist"
	policy "github.com/AitorLeon89/provider-ovh/internal/controller/iampolicy/policy"
	iploadbalancing "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancing/iploadbalancing"
	httpfarm "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancinghttpfarm/httpfarm"
	httpfarmserver "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancinghttpfarmserver/httpfarmserver"
	httpfrontend "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancinghttpfrontend/httpfrontend"
	httproute "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancinghttproute/httproute"
	httprouterule "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancinghttprouterule/httprouterule"
	refresh "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancingrefresh/refresh"
	tcpfarm "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancingtcpfarm/tcpfarm"
	tcpfarmserver "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancingtcpfarmserver/tcpfarmserver"
	tcpfrontend "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancingtcpfrontend/tcpfrontend"
	tcproute "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancingtcproute/tcproute"
	tcprouterule "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancingtcprouterule/tcprouterule"
	vracknetwork "github.com/AitorLeon89/provider-ovh/internal/controller/iploadbalancingvracknetwork/vracknetwork"
	reverse "github.com/AitorLeon89/provider-ovh/internal/controller/ipreverse/reverse"
	service "github.com/AitorLeon89/provider-ovh/internal/controller/ipservice/service"
	identitygroup "github.com/AitorLeon89/provider-ovh/internal/controller/meidentitygroup/identitygroup"
	identityuser "github.com/AitorLeon89/provider-ovh/internal/controller/meidentityuser/identityuser"
	installationtemplate "github.com/AitorLeon89/provider-ovh/internal/controller/meinstallationtemplate/installationtemplate"
	installationtemplatepartitionscheme "github.com/AitorLeon89/provider-ovh/internal/controller/meinstallationtemplatepartitionscheme/installationtemplatepartitionscheme"
	installationtemplatepartitionschemehardwareraid "github.com/AitorLeon89/provider-ovh/internal/controller/meinstallationtemplatepartitionschemehardwareraid/installationtemplatepartitionschemehardwareraid"
	installationtemplatepartitionschemepartition "github.com/AitorLeon89/provider-ovh/internal/controller/meinstallationtemplatepartitionschemepartition/installationtemplatepartitionschemepartition"
	ipxescript "github.com/AitorLeon89/provider-ovh/internal/controller/meipxescript/ipxescript"
	sshkey "github.com/AitorLeon89/provider-ovh/internal/controller/messhkey/sshkey"
	providerconfig "github.com/AitorLeon89/provider-ovh/internal/controller/providerconfig"
	cloudproject "github.com/AitorLeon89/provider-ovh/internal/controller/vrack/cloudproject"
	vrack "github.com/AitorLeon89/provider-ovh/internal/controller/vrack/vrack"
	dedicatedserver "github.com/AitorLeon89/provider-ovh/internal/controller/vrackdedicatedserver/dedicatedserver"
	dedicatedserverinterface "github.com/AitorLeon89/provider-ovh/internal/controller/vrackdedicatedserverinterface/dedicatedserverinterface"
	ip "github.com/AitorLeon89/provider-ovh/internal/controller/vrackip/ip"
	iploadbalancingvrackiploadbalancing "github.com/AitorLeon89/provider-ovh/internal/controller/vrackiploadbalancing/iploadbalancing"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.Setup,
		projectcontainerregistry.Setup,
		projectcontainerregistryoidc.Setup,
		projectcontainerregistryuser.Setup,
		projectdatabase.Setup,
		projectdatabaseuser.Setup,
		projectdatabasedatabase.Setup,
		projectdatabaseintegration.Setup,
		projectdatabaseiprestriction.Setup,
		projectdatabasekafkaacl.Setup,
		projectdatabasekafkaschemaregistryacl.Setup,
		projectdatabasekafkatopic.Setup,
		projectdatabasem3dbnamespace.Setup,
		projectdatabasem3dbuser.Setup,
		projectdatabasemongodbuser.Setup,
		projectdatabaseopensearchpattern.Setup,
		projectdatabaseopensearchuser.Setup,
		projectdatabasepostgresqluser.Setup,
		projectdatabaseredisuser.Setup,
		projectfailoveripattach.Setup,
		projectkube.Setup,
		projectkubeiprestrictions.Setup,
		projectkubenodepool.Setup,
		projectkubeoidc.Setup,
		projectnetworkprivate.Setup,
		projectnetworkprivatesubnet.Setup,
		projectregionstoragepresign.Setup,
		projectuser.Setup,
		projectusers3credential.Setup,
		projectusers3policy.Setup,
		projectworkflowbackup.Setup,
		logscluster.Setup,
		logsinput.Setup,
		cephacl.Setup,
		nashapartition.Setup,
		nashapartitionaccess.Setup,
		nashapartitionsnapshot.Setup,
		serverinstalltask.Setup,
		serverreboottask.Setup,
		serverupdate.Setup,
		zone.Setup,
		zoneredirection.Setup,
		privatedatabasedatabase.Setup,
		privatedatabaseuser.Setup,
		privatedatabaseusergrant.Setup,
		privatedatabasewhitelist.Setup,
		policy.Setup,
		iploadbalancing.Setup,
		httpfarm.Setup,
		httpfarmserver.Setup,
		httpfrontend.Setup,
		httproute.Setup,
		httprouterule.Setup,
		refresh.Setup,
		tcpfarm.Setup,
		tcpfarmserver.Setup,
		tcpfrontend.Setup,
		tcproute.Setup,
		tcprouterule.Setup,
		vracknetwork.Setup,
		reverse.Setup,
		service.Setup,
		identitygroup.Setup,
		identityuser.Setup,
		installationtemplate.Setup,
		installationtemplatepartitionscheme.Setup,
		installationtemplatepartitionschemehardwareraid.Setup,
		installationtemplatepartitionschemepartition.Setup,
		ipxescript.Setup,
		sshkey.Setup,
		providerconfig.Setup,
		cloudproject.Setup,
		vrack.Setup,
		dedicatedserver.Setup,
		dedicatedserverinterface.Setup,
		ip.Setup,
		iploadbalancingvrackiploadbalancing.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
