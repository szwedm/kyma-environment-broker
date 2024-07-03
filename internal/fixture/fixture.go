package fixture

import (
	"fmt"
	"time"

	"github.com/kyma-project/control-plane/components/provisioner/pkg/gqlschema"
	"github.com/kyma-project/kyma-environment-broker/common/gardener"
	"github.com/kyma-project/kyma-environment-broker/common/orchestration"
	"github.com/kyma-project/kyma-environment-broker/internal"
	"github.com/kyma-project/kyma-environment-broker/internal/ptr"
	"github.com/pivotal-cf/brokerapi/v8/domain"
)

const (
	ServiceId                   = "47c9dcbf-ff30-448e-ab36-d3bad66ba281"
	ServiceName                 = "kymaruntime"
	PlanId                      = "4deee563-e5ec-4731-b9b1-53b42d855f0c"
	TrialPlan                   = "7d55d31d-35ae-4438-bf13-6ffdfa107d9f"
	PlanName                    = "azure"
	GlobalAccountId             = "e8f7ec0a-0cd6-41f0-905d-5d1efa9fb6c4"
	SubscriptionGlobalAccountID = ""
	Region                      = "westeurope"
	ServiceManagerUsername      = "u"
	ServiceManagerPassword      = "p"
	ServiceManagerURL           = "https://service-manager.local"
	InstanceDashboardURL        = "https://dashboard.local"
	XSUAADataXSAppName          = "XSApp"
	MonitoringUsername          = "username"
	MonitoringPassword          = "password"
)

type SimpleInputCreator struct {
	Labels            map[string]string
	ShootName         *string
	ShootDomain       string
	shootDnsProviders gardener.DNSProvidersData
	CloudProvider     internal.CloudProvider
	RuntimeID         string
	Config            *internal.ConfigForPlan
}

func FixServiceManagerEntryDTO() *internal.ServiceManagerEntryDTO {
	return &internal.ServiceManagerEntryDTO{
		Credentials: internal.ServiceManagerCredentials{
			BasicAuth: internal.ServiceManagerBasicAuth{
				Username: ServiceManagerUsername,
				Password: ServiceManagerPassword,
			},
		},
		URL: ServiceManagerURL,
	}
}

func FixERSContext(id string) internal.ERSContext {
	var (
		tenantID     = fmt.Sprintf("Tenant-%s", id)
		subAccountId = fmt.Sprintf("SA-%s", id)
		userID       = fmt.Sprintf("User-%s", id)
		licenseType  = ""
	)

	return internal.ERSContext{
		TenantID:        tenantID,
		SubAccountID:    subAccountId,
		GlobalAccountID: GlobalAccountId,
		Active:          ptr.Bool(true),
		UserID:          userID,
		LicenseType:     &licenseType,
	}
}

func FixProvisioningParametersDTO() internal.ProvisioningParametersDTO {
	trialCloudProvider := internal.Azure

	return internal.ProvisioningParametersDTO{
		Name:         "cluster-test",
		VolumeSizeGb: ptr.Integer(50),
		MachineType:  ptr.String("Standard_D8_v3"),
		Region:       ptr.String(Region),
		Purpose:      ptr.String("Purpose"),
		LicenceType:  ptr.String("LicenceType"),
		Zones:        []string{"1"},
		AutoScalerParameters: internal.AutoScalerParameters{
			AutoScalerMin:  ptr.Integer(3),
			AutoScalerMax:  ptr.Integer(10),
			MaxSurge:       ptr.Integer(4),
			MaxUnavailable: ptr.Integer(1),
		},
		Provider: &trialCloudProvider,
	}
}

func FixProvisioningParameters(id string) internal.ProvisioningParameters {
	return internal.ProvisioningParameters{
		PlanID:         PlanId,
		ServiceID:      ServiceId,
		ErsContext:     FixERSContext(id),
		Parameters:     FixProvisioningParametersDTO(),
		PlatformRegion: Region,
	}
}

func FixInstanceDetails(id string) internal.InstanceDetails {
	var (
		runtimeId    = fmt.Sprintf("runtime-%s", id)
		subAccountId = fmt.Sprintf("SA-%s", id)
		shootName    = fmt.Sprintf("Shoot-%s", id)
		shootDomain  = fmt.Sprintf("shoot-%s.domain.com", id)
	)

	monitoringData := internal.MonitoringData{
		Username: MonitoringUsername,
		Password: MonitoringPassword,
	}

	return internal.InstanceDetails{
		Avs:                   internal.AvsLifecycleData{},
		EventHub:              internal.EventHub{Deleted: false},
		SubAccountID:          subAccountId,
		RuntimeID:             runtimeId,
		ShootName:             shootName,
		ShootDomain:           shootDomain,
		ShootDNSProviders:     FixDNSProvidersConfig(),
		Monitoring:            monitoringData,
		KymaResourceNamespace: "kyma-system",
		KymaResourceName:      runtimeId,
	}
}

func FixInstance(id string) internal.Instance {
	var (
		runtimeId    = fmt.Sprintf("runtime-%s", id)
		subAccountId = fmt.Sprintf("SA-%s", id)
	)

	return internal.Instance{
		InstanceID:                  id,
		RuntimeID:                   runtimeId,
		GlobalAccountID:             GlobalAccountId,
		SubscriptionGlobalAccountID: SubscriptionGlobalAccountID,
		SubAccountID:                subAccountId,
		ServiceID:                   ServiceId,
		ServiceName:                 ServiceName,
		ServicePlanID:               PlanId,
		ServicePlanName:             PlanName,
		DashboardURL:                InstanceDashboardURL,
		Parameters:                  FixProvisioningParameters(id),
		ProviderRegion:              Region,
		Provider:                    internal.Azure,
		InstanceDetails:             FixInstanceDetails(id),
		CreatedAt:                   time.Now(),
		UpdatedAt:                   time.Now().Add(time.Minute * 5),
		Version:                     0,
	}
}

func FixOperation(id, instanceId string, opType internal.OperationType) internal.Operation {
	var (
		description     = fmt.Sprintf("Description for operation %s", id)
		orchestrationId = fmt.Sprintf("Orchestration-%s", id)
	)

	return internal.Operation{
		InstanceDetails:        FixInstanceDetails(instanceId),
		ID:                     id,
		Type:                   opType,
		Version:                0,
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now().Add(time.Hour * 48),
		InstanceID:             instanceId,
		ProvisionerOperationID: "",
		State:                  domain.Succeeded,
		Description:            description,
		ProvisioningParameters: FixProvisioningParameters(id),
		OrchestrationID:        orchestrationId,
		FinishedStages:         []string{"prepare", "check_provisioning"},
	}
}

func FixInputCreator(provider internal.CloudProvider) *SimpleInputCreator {
	return &SimpleInputCreator{
		Labels:        make(map[string]string),
		ShootName:     ptr.String("ShootName"),
		CloudProvider: provider,
	}
}

func FixProvisioningOperation(operationId, instanceId string) internal.Operation {
	o := FixOperation(operationId, instanceId, internal.OperationTypeProvision)
	o.InputCreator = FixInputCreator(internal.Azure)
	o.DashboardURL = "https://console.kyma.org"
	return o
}

func FixUpdatingOperation(operationId, instanceId string) internal.UpdatingOperation {
	o := FixOperation(operationId, instanceId, internal.OperationTypeUpdate)
	o.InputCreator = FixInputCreator(internal.Azure)
	o.UpdatingParameters = internal.UpdatingParametersDTO{
		OIDC: &internal.OIDCConfigDTO{
			ClientID:       "clinet-id-oidc",
			GroupsClaim:    "groups",
			IssuerURL:      "issuer-url",
			SigningAlgs:    []string{"signingAlgs"},
			UsernameClaim:  "sub",
			UsernamePrefix: "",
		},
	}
	return internal.UpdatingOperation{
		Operation: o,
	}
}

func FixProvisioningOperationWithProvider(operationId, instanceId string, provider internal.CloudProvider) internal.Operation {
	o := FixOperation(operationId, instanceId, internal.OperationTypeProvision)
	o.InputCreator = FixInputCreator(provider)
	o.DashboardURL = "https://console.kyma.org"
	return o
}

func FixDeprovisioningOperation(operationId, instanceId string) internal.DeprovisioningOperation {
	return internal.DeprovisioningOperation{
		Operation: FixDeprovisioningOperationAsOperation(operationId, instanceId),
	}
}

func FixDeprovisioningOperationAsOperation(operationId, instanceId string) internal.Operation {
	o := FixOperation(operationId, instanceId, internal.OperationTypeDeprovision)
	o.Temporary = false
	return o
}

func FixSuspensionOperationAsOperation(operationId, instanceId string) internal.Operation {
	o := FixOperation(operationId, instanceId, internal.OperationTypeDeprovision)
	o.Temporary = true
	o.ProvisioningParameters.PlanID = TrialPlan
	return o
}

func FixRuntime(id string) orchestration.Runtime {
	var (
		instanceId   = fmt.Sprintf("Instance-%s", id)
		subAccountId = fmt.Sprintf("SA-%s", id)
	)

	return orchestration.Runtime{
		InstanceID:             instanceId,
		RuntimeID:              id,
		GlobalAccountID:        GlobalAccountId,
		SubAccountID:           subAccountId,
		ShootName:              "ShootName",
		MaintenanceWindowBegin: time.Now().Truncate(time.Millisecond).Add(time.Hour),
		MaintenanceWindowEnd:   time.Now().Truncate(time.Millisecond).Add(time.Minute).Add(time.Hour),
	}
}

func FixRuntimeOperation(operationId string) orchestration.RuntimeOperation {
	return orchestration.RuntimeOperation{
		Runtime: FixRuntime(operationId),
		ID:      operationId,
		DryRun:  false,
	}
}

func FixUpgradeClusterOperation(operationId, instanceId string) internal.UpgradeClusterOperation {
	o := FixOperation(operationId, instanceId, internal.OperationTypeUpgradeCluster)
	o.RuntimeOperation = FixRuntimeOperation(operationId)
	o.InputCreator = FixInputCreator(internal.Azure)
	return internal.UpgradeClusterOperation{
		Operation: o,
	}
}

func FixOrchestration(id string) internal.Orchestration {
	return internal.Orchestration{
		OrchestrationID: id,
		State:           orchestration.Succeeded,
		Description:     "",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now().Add(time.Hour * 1),
		Parameters:      orchestration.Parameters{},
	}
}

func FixOIDCConfigDTO() internal.OIDCConfigDTO {
	return internal.OIDCConfigDTO{
		ClientID:       "9bd05ed7-a930-44e6-8c79-e6defeb7dec9",
		GroupsClaim:    "groups",
		IssuerURL:      "https://kymatest.accounts400.ondemand.com",
		SigningAlgs:    []string{"RS256"},
		UsernameClaim:  "sub",
		UsernamePrefix: "-",
	}
}

func FixDNSProvidersConfig() gardener.DNSProvidersData {
	return gardener.DNSProvidersData{
		Providers: []gardener.DNSProviderData{
			{
				DomainsInclude: []string{"devtest.kyma.ondemand.com"},
				Primary:        true,
				SecretName:     "aws_dns_domain_secrets_test_incustom",
				Type:           "route53_type_test",
			},
		},
	}
}

func FixRuntimeState(id, runtimeID, operationID string) internal.RuntimeState {
	disabled := false
	return internal.RuntimeState{
		ID:          id,
		CreatedAt:   time.Now(),
		RuntimeID:   runtimeID,
		OperationID: operationID,
		KymaConfig:  gqlschema.KymaConfigInput{},
		ClusterConfig: gqlschema.GardenerConfigInput{
			ShootNetworkingFilterDisabled: &disabled,
		},
	}
}

// SimpleInputCreator implements ProvisionerInputCreator interface
func (c *SimpleInputCreator) SetProvisioningParameters(params internal.ProvisioningParameters) internal.ProvisionerInputCreator {
	return c
}

func (c *SimpleInputCreator) SetShootName(name string) internal.ProvisionerInputCreator {
	c.ShootName = &name
	return c
}

func (c *SimpleInputCreator) SetShootDomain(name string) internal.ProvisionerInputCreator {
	c.ShootDomain = name
	return c
}

func (c *SimpleInputCreator) SetShootDNSProviders(providers gardener.DNSProvidersData) internal.ProvisionerInputCreator {
	c.shootDnsProviders = providers
	return c
}

func (c *SimpleInputCreator) SetLabel(key, val string) internal.ProvisionerInputCreator {
	c.Labels[key] = val
	return c
}

func (c *SimpleInputCreator) SetKubeconfig(_ string) internal.ProvisionerInputCreator {
	return c
}

func (c *SimpleInputCreator) SetClusterName(_ string) internal.ProvisionerInputCreator {
	return c
}

func (c *SimpleInputCreator) SetInstanceID(kcfg string) internal.ProvisionerInputCreator {
	return c
}

func (c *SimpleInputCreator) SetRuntimeID(runtimeID string) internal.ProvisionerInputCreator {
	c.RuntimeID = runtimeID
	return c
}

func (c *SimpleInputCreator) SetOIDCLastValues(oidcConfig gqlschema.OIDCConfigInput) internal.ProvisionerInputCreator {
	return c
}

func (c *SimpleInputCreator) CreateProvisionClusterInput() (gqlschema.ProvisionRuntimeInput, error) {
	return gqlschema.ProvisionRuntimeInput{}, nil
}

func (c *SimpleInputCreator) CreateProvisionRuntimeInput() (gqlschema.ProvisionRuntimeInput, error) {
	return gqlschema.ProvisionRuntimeInput{}, nil
}

func (c *SimpleInputCreator) CreateUpgradeRuntimeInput() (gqlschema.UpgradeRuntimeInput, error) {
	return gqlschema.UpgradeRuntimeInput{}, nil
}

func (c *SimpleInputCreator) CreateUpgradeShootInput() (gqlschema.UpgradeShootInput, error) {
	return gqlschema.UpgradeShootInput{}, nil
}

func (c *SimpleInputCreator) Provider() internal.CloudProvider {
	return c.CloudProvider
}

func (c *SimpleInputCreator) Configuration() *internal.ConfigForPlan {
	return c.Config
}
