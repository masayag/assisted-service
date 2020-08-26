// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"

	"github.com/openshift/assisted-service/restapi/operations"
	"github.com/openshift/assisted-service/restapi/operations/events"
	"github.com/openshift/assisted-service/restapi/operations/installer"
	"github.com/openshift/assisted-service/restapi/operations/managed_domains"
	"github.com/openshift/assisted-service/restapi/operations/manifests"
	"github.com/openshift/assisted-service/restapi/operations/versions"
)

type contextKey string

const AuthKey contextKey = "Auth"

//go:generate mockery -name EventsAPI -inpkg

/* EventsAPI  */
type EventsAPI interface {
	/* ListEvents Lists events for a cluster. */
	ListEvents(ctx context.Context, params events.ListEventsParams) middleware.Responder
}

//go:generate mockery -name InstallerAPI -inpkg

/* InstallerAPI  */
type InstallerAPI interface {
	/* CancelInstallation Cancels an ongoing installation. */
	CancelInstallation(ctx context.Context, params installer.CancelInstallationParams) middleware.Responder

	/* CompleteInstallation Agent API to mark a finalizing installation as complete. */
	CompleteInstallation(ctx context.Context, params installer.CompleteInstallationParams) middleware.Responder

	/* DeregisterCluster Deletes an OpenShift bare metal cluster definition. */
	DeregisterCluster(ctx context.Context, params installer.DeregisterClusterParams) middleware.Responder

	/* DeregisterHost Deregisters an OpenShift bare metal host. */
	DeregisterHost(ctx context.Context, params installer.DeregisterHostParams) middleware.Responder

	/* DisableHost Disables a host for inclusion in the cluster. */
	DisableHost(ctx context.Context, params installer.DisableHostParams) middleware.Responder

	/* DownloadClusterFiles Downloads files relating to the installed/installing cluster. */
	DownloadClusterFiles(ctx context.Context, params installer.DownloadClusterFilesParams) middleware.Responder

	/* DownloadClusterISO Downloads the OpenShift per-cluster Discovery ISO. */
	DownloadClusterISO(ctx context.Context, params installer.DownloadClusterISOParams) middleware.Responder

	/* DownloadClusterKubeconfig Downloads the kubeconfig file for this cluster. */
	DownloadClusterKubeconfig(ctx context.Context, params installer.DownloadClusterKubeconfigParams) middleware.Responder

	/* DownloadClusterLogs Download cluster logs. */
	DownloadClusterLogs(ctx context.Context, params installer.DownloadClusterLogsParams) middleware.Responder

	/* DownloadHostLogs Download host logs. */
	DownloadHostLogs(ctx context.Context, params installer.DownloadHostLogsParams) middleware.Responder

	/* EnableHost Enables a host for inclusion in the cluster. */
	EnableHost(ctx context.Context, params installer.EnableHostParams) middleware.Responder

	/* GenerateClusterISO Creates a new OpenShift per-cluster Discovery ISO. */
	GenerateClusterISO(ctx context.Context, params installer.GenerateClusterISOParams) middleware.Responder

	/* GetCluster Retrieves the details of the OpenShift bare metal cluster. */
	GetCluster(ctx context.Context, params installer.GetClusterParams) middleware.Responder

	/* GetClusterInstallConfig Get the cluster's install config YAML. */
	GetClusterInstallConfig(ctx context.Context, params installer.GetClusterInstallConfigParams) middleware.Responder

	/* GetCredentials Get the cluster admin credentials. */
	GetCredentials(ctx context.Context, params installer.GetCredentialsParams) middleware.Responder

	/* GetDiscoveryIgnition Get the cluster discovery ignition config */
	GetDiscoveryIgnition(ctx context.Context, params installer.GetDiscoveryIgnitionParams) middleware.Responder

	/* GetFreeAddresses Retrieves the free address list for a network. */
	GetFreeAddresses(ctx context.Context, params installer.GetFreeAddressesParams) middleware.Responder

	/* GetHost Retrieves the details of the OpenShift bare metal host. */
	GetHost(ctx context.Context, params installer.GetHostParams) middleware.Responder

	/* GetHostRequirements Get minimum host requirements. */
	GetHostRequirements(ctx context.Context, params installer.GetHostRequirementsParams) middleware.Responder

	/* GetNextSteps Retrieves the next operations that the host agent needs to perform. */
	GetNextSteps(ctx context.Context, params installer.GetNextStepsParams) middleware.Responder

	/* GetPresignedForClusterFiles Retrieves a pre-signed S3 URL for downloading cluster files. */
	GetPresignedForClusterFiles(ctx context.Context, params installer.GetPresignedForClusterFilesParams) middleware.Responder

	/* InstallCluster Installs the OpenShift bare metal cluster. */
	InstallCluster(ctx context.Context, params installer.InstallClusterParams) middleware.Responder

	/* InstallHosts Installs the OpenShift bare metal cluster. */
	InstallHosts(ctx context.Context, params installer.InstallHostsParams) middleware.Responder

	/* ListClusters Retrieves the list of OpenShift bare metal clusters. */
	ListClusters(ctx context.Context, params installer.ListClustersParams) middleware.Responder

	/* ListHosts Retrieves the list of OpenShift bare metal hosts. */
	ListHosts(ctx context.Context, params installer.ListHostsParams) middleware.Responder

	/* PostStepReply Posts the result of the operations from the host agent. */
	PostStepReply(ctx context.Context, params installer.PostStepReplyParams) middleware.Responder

	/* RegisterAddHostsCluster Creates a new OpenShift bare metal cluster definition for adding nodes to and existing OCP cluster. */
	RegisterAddHostsCluster(ctx context.Context, params installer.RegisterAddHostsClusterParams) middleware.Responder

	/* RegisterCluster Creates a new OpenShift bare metal cluster definition. */
	RegisterCluster(ctx context.Context, params installer.RegisterClusterParams) middleware.Responder

	/* RegisterHost Registers a new OpenShift bare metal host. */
	RegisterHost(ctx context.Context, params installer.RegisterHostParams) middleware.Responder

	/* ResetCluster Resets a failed installation. */
	ResetCluster(ctx context.Context, params installer.ResetClusterParams) middleware.Responder

	/* UpdateCluster Updates an OpenShift bare metal cluster definition. */
	UpdateCluster(ctx context.Context, params installer.UpdateClusterParams) middleware.Responder

	/* UpdateClusterInstallConfig Override values in the install config. */
	UpdateClusterInstallConfig(ctx context.Context, params installer.UpdateClusterInstallConfigParams) middleware.Responder

	/* UpdateDiscoveryIgnition Override values in the discovery ignition config */
	UpdateDiscoveryIgnition(ctx context.Context, params installer.UpdateDiscoveryIgnitionParams) middleware.Responder

	/* UpdateHostInstallProgress Update installation progress. */
	UpdateHostInstallProgress(ctx context.Context, params installer.UpdateHostInstallProgressParams) middleware.Responder

	/* UploadClusterIngressCert Transfer the ingress certificate for the cluster. */
	UploadClusterIngressCert(ctx context.Context, params installer.UploadClusterIngressCertParams) middleware.Responder

	/* UploadHostLogs Agent API to upload logs. */
	UploadHostLogs(ctx context.Context, params installer.UploadHostLogsParams) middleware.Responder

	/* UploadLogs Agent API to upload logs. */
	UploadLogs(ctx context.Context, params installer.UploadLogsParams) middleware.Responder
}

//go:generate mockery -name ManagedDomainsAPI -inpkg

/* ManagedDomainsAPI  */
type ManagedDomainsAPI interface {
	/* ListManagedDomains List of managed DNS domains. */
	ListManagedDomains(ctx context.Context, params managed_domains.ListManagedDomainsParams) middleware.Responder
}

//go:generate mockery -name ManifestsAPI -inpkg

/* ManifestsAPI  */
type ManifestsAPI interface {
	/* CreateClusterManifest Creates a manifest for customizing cluster installation. */
	CreateClusterManifest(ctx context.Context, params manifests.CreateClusterManifestParams) middleware.Responder

	/* DeleteClusterManifest Deletes a manifest from the cluster. */
	DeleteClusterManifest(ctx context.Context, params manifests.DeleteClusterManifestParams) middleware.Responder

	/* DownloadClusterManifest Downloads cluster manifest. */
	DownloadClusterManifest(ctx context.Context, params manifests.DownloadClusterManifestParams) middleware.Responder

	/* ListClusterManifests Lists manifests for customizing cluster installation. */
	ListClusterManifests(ctx context.Context, params manifests.ListClusterManifestsParams) middleware.Responder
}

//go:generate mockery -name VersionsAPI -inpkg

/* VersionsAPI  */
type VersionsAPI interface {
	/* ListComponentVersions List of component versions. */
	ListComponentVersions(ctx context.Context, params versions.ListComponentVersionsParams) middleware.Responder
}

// Config is configuration for Handler
type Config struct {
	EventsAPI
	InstallerAPI
	ManagedDomainsAPI
	ManifestsAPI
	VersionsAPI
	Logger func(string, ...interface{})
	// InnerMiddleware is for the handler executors. These do not apply to the swagger.json document.
	// The middleware executes after routing but before authentication, binding and validation
	InnerMiddleware func(http.Handler) http.Handler

	// Authorizer is used to authorize a request after the Auth function was called using the "Auth*" functions
	// and the principal was stored in the context in the "AuthKey" context value.
	Authorizer func(*http.Request) error

	// AuthAgentAuth Applies when the "X-Secret-Key" header is set
	AuthAgentAuth func(token string) (interface{}, error)

	// AuthUserAuth Applies when the "Authorization" header is set
	AuthUserAuth func(token string) (interface{}, error)

	// Authenticator to use for all APIKey authentication
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// Authenticator to use for all Bearer authentication
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// Authenticator to use for all Basic authentication
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator
}

// Handler returns an http.Handler given the handler configuration
// It mounts all the business logic implementers in the right routing.
func Handler(c Config) (http.Handler, error) {
	h, _, err := HandlerAPI(c)
	return h, err
}

// HandlerAPI returns an http.Handler given the handler configuration
// and the corresponding *AssistedInstall instance.
// It mounts all the business logic implementers in the right routing.
func HandlerAPI(c Config) (http.Handler, *operations.AssistedInstallAPI, error) {
	spec, err := loads.Analyzed(swaggerCopy(SwaggerJSON), "")
	if err != nil {
		return nil, nil, fmt.Errorf("analyze swagger: %v", err)
	}
	api := operations.NewAssistedInstallAPI(spec)
	api.ServeError = errors.ServeError
	api.Logger = c.Logger

	if c.APIKeyAuthenticator != nil {
		api.APIKeyAuthenticator = c.APIKeyAuthenticator
	}
	if c.BasicAuthenticator != nil {
		api.BasicAuthenticator = c.BasicAuthenticator
	}
	if c.BearerAuthenticator != nil {
		api.BearerAuthenticator = c.BearerAuthenticator
	}

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer
	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()
	api.AgentAuthAuth = func(token string) (interface{}, error) {
		if c.AuthAgentAuth == nil {
			return token, nil
		}
		return c.AuthAgentAuth(token)
	}

	api.UserAuthAuth = func(token string) (interface{}, error) {
		if c.AuthUserAuth == nil {
			return token, nil
		}
		return c.AuthUserAuth(token)
	}

	api.APIAuthorizer = authorizer(c.Authorizer)
	api.InstallerCancelInstallationHandler = installer.CancelInstallationHandlerFunc(func(params installer.CancelInstallationParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.CancelInstallation(ctx, params)
	})
	api.InstallerCompleteInstallationHandler = installer.CompleteInstallationHandlerFunc(func(params installer.CompleteInstallationParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.CompleteInstallation(ctx, params)
	})
	api.ManifestsCreateClusterManifestHandler = manifests.CreateClusterManifestHandlerFunc(func(params manifests.CreateClusterManifestParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.ManifestsAPI.CreateClusterManifest(ctx, params)
	})
	api.ManifestsDeleteClusterManifestHandler = manifests.DeleteClusterManifestHandlerFunc(func(params manifests.DeleteClusterManifestParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.ManifestsAPI.DeleteClusterManifest(ctx, params)
	})
	api.InstallerDeregisterClusterHandler = installer.DeregisterClusterHandlerFunc(func(params installer.DeregisterClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.DeregisterCluster(ctx, params)
	})
	api.InstallerDeregisterHostHandler = installer.DeregisterHostHandlerFunc(func(params installer.DeregisterHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.DeregisterHost(ctx, params)
	})
	api.InstallerDisableHostHandler = installer.DisableHostHandlerFunc(func(params installer.DisableHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.DisableHost(ctx, params)
	})
	api.InstallerDownloadClusterFilesHandler = installer.DownloadClusterFilesHandlerFunc(func(params installer.DownloadClusterFilesParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.DownloadClusterFiles(ctx, params)
	})
	api.InstallerDownloadClusterISOHandler = installer.DownloadClusterISOHandlerFunc(func(params installer.DownloadClusterISOParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.DownloadClusterISO(ctx, params)
	})
	api.InstallerDownloadClusterKubeconfigHandler = installer.DownloadClusterKubeconfigHandlerFunc(func(params installer.DownloadClusterKubeconfigParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.DownloadClusterKubeconfig(ctx, params)
	})
	api.InstallerDownloadClusterLogsHandler = installer.DownloadClusterLogsHandlerFunc(func(params installer.DownloadClusterLogsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.DownloadClusterLogs(ctx, params)
	})
	api.ManifestsDownloadClusterManifestHandler = manifests.DownloadClusterManifestHandlerFunc(func(params manifests.DownloadClusterManifestParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.ManifestsAPI.DownloadClusterManifest(ctx, params)
	})
	api.InstallerDownloadHostLogsHandler = installer.DownloadHostLogsHandlerFunc(func(params installer.DownloadHostLogsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.DownloadHostLogs(ctx, params)
	})
	api.InstallerEnableHostHandler = installer.EnableHostHandlerFunc(func(params installer.EnableHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.EnableHost(ctx, params)
	})
	api.InstallerGenerateClusterISOHandler = installer.GenerateClusterISOHandlerFunc(func(params installer.GenerateClusterISOParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GenerateClusterISO(ctx, params)
	})
	api.InstallerGetClusterHandler = installer.GetClusterHandlerFunc(func(params installer.GetClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetCluster(ctx, params)
	})
	api.InstallerGetClusterInstallConfigHandler = installer.GetClusterInstallConfigHandlerFunc(func(params installer.GetClusterInstallConfigParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetClusterInstallConfig(ctx, params)
	})
	api.InstallerGetCredentialsHandler = installer.GetCredentialsHandlerFunc(func(params installer.GetCredentialsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetCredentials(ctx, params)
	})
	api.InstallerGetDiscoveryIgnitionHandler = installer.GetDiscoveryIgnitionHandlerFunc(func(params installer.GetDiscoveryIgnitionParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetDiscoveryIgnition(ctx, params)
	})
	api.InstallerGetFreeAddressesHandler = installer.GetFreeAddressesHandlerFunc(func(params installer.GetFreeAddressesParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetFreeAddresses(ctx, params)
	})
	api.InstallerGetHostHandler = installer.GetHostHandlerFunc(func(params installer.GetHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetHost(ctx, params)
	})
	api.InstallerGetHostRequirementsHandler = installer.GetHostRequirementsHandlerFunc(func(params installer.GetHostRequirementsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetHostRequirements(ctx, params)
	})
	api.InstallerGetNextStepsHandler = installer.GetNextStepsHandlerFunc(func(params installer.GetNextStepsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetNextSteps(ctx, params)
	})
	api.InstallerGetPresignedForClusterFilesHandler = installer.GetPresignedForClusterFilesHandlerFunc(func(params installer.GetPresignedForClusterFilesParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetPresignedForClusterFiles(ctx, params)
	})
	api.InstallerInstallClusterHandler = installer.InstallClusterHandlerFunc(func(params installer.InstallClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.InstallCluster(ctx, params)
	})
	api.InstallerInstallHostsHandler = installer.InstallHostsHandlerFunc(func(params installer.InstallHostsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.InstallHosts(ctx, params)
	})
	api.ManifestsListClusterManifestsHandler = manifests.ListClusterManifestsHandlerFunc(func(params manifests.ListClusterManifestsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.ManifestsAPI.ListClusterManifests(ctx, params)
	})
	api.InstallerListClustersHandler = installer.ListClustersHandlerFunc(func(params installer.ListClustersParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.ListClusters(ctx, params)
	})
	api.VersionsListComponentVersionsHandler = versions.ListComponentVersionsHandlerFunc(func(params versions.ListComponentVersionsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.VersionsAPI.ListComponentVersions(ctx, params)
	})
	api.EventsListEventsHandler = events.ListEventsHandlerFunc(func(params events.ListEventsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.EventsAPI.ListEvents(ctx, params)
	})
	api.InstallerListHostsHandler = installer.ListHostsHandlerFunc(func(params installer.ListHostsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.ListHosts(ctx, params)
	})
	api.ManagedDomainsListManagedDomainsHandler = managed_domains.ListManagedDomainsHandlerFunc(func(params managed_domains.ListManagedDomainsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.ManagedDomainsAPI.ListManagedDomains(ctx, params)
	})
	api.InstallerPostStepReplyHandler = installer.PostStepReplyHandlerFunc(func(params installer.PostStepReplyParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.PostStepReply(ctx, params)
	})
	api.InstallerRegisterAddHostsClusterHandler = installer.RegisterAddHostsClusterHandlerFunc(func(params installer.RegisterAddHostsClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.RegisterAddHostsCluster(ctx, params)
	})
	api.InstallerRegisterClusterHandler = installer.RegisterClusterHandlerFunc(func(params installer.RegisterClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.RegisterCluster(ctx, params)
	})
	api.InstallerRegisterHostHandler = installer.RegisterHostHandlerFunc(func(params installer.RegisterHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.RegisterHost(ctx, params)
	})
	api.InstallerResetClusterHandler = installer.ResetClusterHandlerFunc(func(params installer.ResetClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.ResetCluster(ctx, params)
	})
	api.InstallerUpdateClusterHandler = installer.UpdateClusterHandlerFunc(func(params installer.UpdateClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.UpdateCluster(ctx, params)
	})
	api.InstallerUpdateClusterInstallConfigHandler = installer.UpdateClusterInstallConfigHandlerFunc(func(params installer.UpdateClusterInstallConfigParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.UpdateClusterInstallConfig(ctx, params)
	})
	api.InstallerUpdateDiscoveryIgnitionHandler = installer.UpdateDiscoveryIgnitionHandlerFunc(func(params installer.UpdateDiscoveryIgnitionParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.UpdateDiscoveryIgnition(ctx, params)
	})
	api.InstallerUpdateHostInstallProgressHandler = installer.UpdateHostInstallProgressHandlerFunc(func(params installer.UpdateHostInstallProgressParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.UpdateHostInstallProgress(ctx, params)
	})
	api.InstallerUploadClusterIngressCertHandler = installer.UploadClusterIngressCertHandlerFunc(func(params installer.UploadClusterIngressCertParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.UploadClusterIngressCert(ctx, params)
	})
	api.InstallerUploadHostLogsHandler = installer.UploadHostLogsHandlerFunc(func(params installer.UploadHostLogsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.UploadHostLogs(ctx, params)
	})
	api.InstallerUploadLogsHandler = installer.UploadLogsHandlerFunc(func(params installer.UploadLogsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.UploadLogs(ctx, params)
	})
	api.ServerShutdown = func() {}
	return api.Serve(c.InnerMiddleware), api, nil
}

// swaggerCopy copies the swagger json to prevent data races in runtime
func swaggerCopy(orig json.RawMessage) json.RawMessage {
	c := make(json.RawMessage, len(orig))
	copy(c, orig)
	return c
}

// authorizer is a helper function to implement the runtime.Authorizer interface.
type authorizer func(*http.Request) error

func (a authorizer) Authorize(req *http.Request, principal interface{}) error {
	if a == nil {
		return nil
	}
	ctx := storeAuth(req.Context(), principal)
	return a(req.WithContext(ctx))
}

func storeAuth(ctx context.Context, principal interface{}) context.Context {
	return context.WithValue(ctx, AuthKey, principal)
}
