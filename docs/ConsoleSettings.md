# ConsoleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppsInAnyNamespaceEnabled** | Pointer to **bool** | AppsInAnyNamespaceEnabled is false: applications are projected from operator App CRs in the platform namespaces, never declared in an arbitrary one. | [optional] 
**DexConfig** | Pointer to [**ConsoleSettingsDexConfig**](ConsoleSettingsDexConfig.md) |  | [optional] 
**ExecEnabled** | Pointer to **bool** | ExecEnabled is false: this plane serves no container terminal. | [optional] 
**GoogleAnalytics** | Pointer to [**ConsoleSettingsGoogleAnalytics**](ConsoleSettingsGoogleAnalytics.md) |  | [optional] 
**Help** | Pointer to [**ConsoleSettingsHelp**](ConsoleSettingsHelp.md) |  | [optional] 
**HydratorEnabled** | Pointer to **bool** | HydratorEnabled is false: there is no manifest hydrator on this plane. | [optional] 
**KustomizeVersions** | Pointer to **[]string** | KustomizeVersions is always empty: an App CR is an image pin, not a kustomize build. | [optional] 
**OidcConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**Plugins** | Pointer to **[]map[string]interface{}** | Plugins is always empty: this plane loads no argocd config-management plugins. | [optional] 
**StatusBadgeEnabled** | Pointer to **bool** | StatusBadgeEnabled is false: no badge endpoint is served. | [optional] 
**StatusBadgeRootUrl** | Pointer to **string** | StatusBadgeRootUrl is always empty, for the same reason. | [optional] 
**SyncWithReplaceAllowed** | Pointer to **bool** | SyncWithReplaceAllowed is false: a sync here asks the operator to reconcile an App CR, and never replaces an object. | [optional] 
**UiBannerContent** | Pointer to **string** | UiBannerContent is always empty: this console shows no banner. | [optional] 
**UiCssURL** | Pointer to **string** | UiCssURL is always empty: no stylesheet is injected. | [optional] 
**Url** | Pointer to **string** | Url is the console&#39;s public origin, https://cd.hanzo.ai. | [optional] 
**UserLoginsDisabled** | Pointer to **bool** | UserLoginsDisabled is true: the SPA must not render its own username/password form. Signing in goes through IAM, at GET /v1/deploy/login. | [optional] 

## Methods

### NewConsoleSettings

`func NewConsoleSettings() *ConsoleSettings`

NewConsoleSettings instantiates a new ConsoleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleSettingsWithDefaults

`func NewConsoleSettingsWithDefaults() *ConsoleSettings`

NewConsoleSettingsWithDefaults instantiates a new ConsoleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppsInAnyNamespaceEnabled

`func (o *ConsoleSettings) GetAppsInAnyNamespaceEnabled() bool`

GetAppsInAnyNamespaceEnabled returns the AppsInAnyNamespaceEnabled field if non-nil, zero value otherwise.

### GetAppsInAnyNamespaceEnabledOk

`func (o *ConsoleSettings) GetAppsInAnyNamespaceEnabledOk() (*bool, bool)`

GetAppsInAnyNamespaceEnabledOk returns a tuple with the AppsInAnyNamespaceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsInAnyNamespaceEnabled

`func (o *ConsoleSettings) SetAppsInAnyNamespaceEnabled(v bool)`

SetAppsInAnyNamespaceEnabled sets AppsInAnyNamespaceEnabled field to given value.

### HasAppsInAnyNamespaceEnabled

`func (o *ConsoleSettings) HasAppsInAnyNamespaceEnabled() bool`

HasAppsInAnyNamespaceEnabled returns a boolean if a field has been set.

### GetDexConfig

`func (o *ConsoleSettings) GetDexConfig() ConsoleSettingsDexConfig`

GetDexConfig returns the DexConfig field if non-nil, zero value otherwise.

### GetDexConfigOk

`func (o *ConsoleSettings) GetDexConfigOk() (*ConsoleSettingsDexConfig, bool)`

GetDexConfigOk returns a tuple with the DexConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDexConfig

`func (o *ConsoleSettings) SetDexConfig(v ConsoleSettingsDexConfig)`

SetDexConfig sets DexConfig field to given value.

### HasDexConfig

`func (o *ConsoleSettings) HasDexConfig() bool`

HasDexConfig returns a boolean if a field has been set.

### GetExecEnabled

`func (o *ConsoleSettings) GetExecEnabled() bool`

GetExecEnabled returns the ExecEnabled field if non-nil, zero value otherwise.

### GetExecEnabledOk

`func (o *ConsoleSettings) GetExecEnabledOk() (*bool, bool)`

GetExecEnabledOk returns a tuple with the ExecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecEnabled

`func (o *ConsoleSettings) SetExecEnabled(v bool)`

SetExecEnabled sets ExecEnabled field to given value.

### HasExecEnabled

`func (o *ConsoleSettings) HasExecEnabled() bool`

HasExecEnabled returns a boolean if a field has been set.

### GetGoogleAnalytics

`func (o *ConsoleSettings) GetGoogleAnalytics() ConsoleSettingsGoogleAnalytics`

GetGoogleAnalytics returns the GoogleAnalytics field if non-nil, zero value otherwise.

### GetGoogleAnalyticsOk

`func (o *ConsoleSettings) GetGoogleAnalyticsOk() (*ConsoleSettingsGoogleAnalytics, bool)`

GetGoogleAnalyticsOk returns a tuple with the GoogleAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAnalytics

`func (o *ConsoleSettings) SetGoogleAnalytics(v ConsoleSettingsGoogleAnalytics)`

SetGoogleAnalytics sets GoogleAnalytics field to given value.

### HasGoogleAnalytics

`func (o *ConsoleSettings) HasGoogleAnalytics() bool`

HasGoogleAnalytics returns a boolean if a field has been set.

### GetHelp

`func (o *ConsoleSettings) GetHelp() ConsoleSettingsHelp`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *ConsoleSettings) GetHelpOk() (*ConsoleSettingsHelp, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *ConsoleSettings) SetHelp(v ConsoleSettingsHelp)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *ConsoleSettings) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetHydratorEnabled

`func (o *ConsoleSettings) GetHydratorEnabled() bool`

GetHydratorEnabled returns the HydratorEnabled field if non-nil, zero value otherwise.

### GetHydratorEnabledOk

`func (o *ConsoleSettings) GetHydratorEnabledOk() (*bool, bool)`

GetHydratorEnabledOk returns a tuple with the HydratorEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHydratorEnabled

`func (o *ConsoleSettings) SetHydratorEnabled(v bool)`

SetHydratorEnabled sets HydratorEnabled field to given value.

### HasHydratorEnabled

`func (o *ConsoleSettings) HasHydratorEnabled() bool`

HasHydratorEnabled returns a boolean if a field has been set.

### GetKustomizeVersions

`func (o *ConsoleSettings) GetKustomizeVersions() []string`

GetKustomizeVersions returns the KustomizeVersions field if non-nil, zero value otherwise.

### GetKustomizeVersionsOk

`func (o *ConsoleSettings) GetKustomizeVersionsOk() (*[]string, bool)`

GetKustomizeVersionsOk returns a tuple with the KustomizeVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKustomizeVersions

`func (o *ConsoleSettings) SetKustomizeVersions(v []string)`

SetKustomizeVersions sets KustomizeVersions field to given value.

### HasKustomizeVersions

`func (o *ConsoleSettings) HasKustomizeVersions() bool`

HasKustomizeVersions returns a boolean if a field has been set.

### GetOidcConfig

`func (o *ConsoleSettings) GetOidcConfig() map[string]interface{}`

GetOidcConfig returns the OidcConfig field if non-nil, zero value otherwise.

### GetOidcConfigOk

`func (o *ConsoleSettings) GetOidcConfigOk() (*map[string]interface{}, bool)`

GetOidcConfigOk returns a tuple with the OidcConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcConfig

`func (o *ConsoleSettings) SetOidcConfig(v map[string]interface{})`

SetOidcConfig sets OidcConfig field to given value.

### HasOidcConfig

`func (o *ConsoleSettings) HasOidcConfig() bool`

HasOidcConfig returns a boolean if a field has been set.

### GetPlugins

`func (o *ConsoleSettings) GetPlugins() []map[string]interface{}`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *ConsoleSettings) GetPluginsOk() (*[]map[string]interface{}, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *ConsoleSettings) SetPlugins(v []map[string]interface{})`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *ConsoleSettings) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetStatusBadgeEnabled

`func (o *ConsoleSettings) GetStatusBadgeEnabled() bool`

GetStatusBadgeEnabled returns the StatusBadgeEnabled field if non-nil, zero value otherwise.

### GetStatusBadgeEnabledOk

`func (o *ConsoleSettings) GetStatusBadgeEnabledOk() (*bool, bool)`

GetStatusBadgeEnabledOk returns a tuple with the StatusBadgeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusBadgeEnabled

`func (o *ConsoleSettings) SetStatusBadgeEnabled(v bool)`

SetStatusBadgeEnabled sets StatusBadgeEnabled field to given value.

### HasStatusBadgeEnabled

`func (o *ConsoleSettings) HasStatusBadgeEnabled() bool`

HasStatusBadgeEnabled returns a boolean if a field has been set.

### GetStatusBadgeRootUrl

`func (o *ConsoleSettings) GetStatusBadgeRootUrl() string`

GetStatusBadgeRootUrl returns the StatusBadgeRootUrl field if non-nil, zero value otherwise.

### GetStatusBadgeRootUrlOk

`func (o *ConsoleSettings) GetStatusBadgeRootUrlOk() (*string, bool)`

GetStatusBadgeRootUrlOk returns a tuple with the StatusBadgeRootUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusBadgeRootUrl

`func (o *ConsoleSettings) SetStatusBadgeRootUrl(v string)`

SetStatusBadgeRootUrl sets StatusBadgeRootUrl field to given value.

### HasStatusBadgeRootUrl

`func (o *ConsoleSettings) HasStatusBadgeRootUrl() bool`

HasStatusBadgeRootUrl returns a boolean if a field has been set.

### GetSyncWithReplaceAllowed

`func (o *ConsoleSettings) GetSyncWithReplaceAllowed() bool`

GetSyncWithReplaceAllowed returns the SyncWithReplaceAllowed field if non-nil, zero value otherwise.

### GetSyncWithReplaceAllowedOk

`func (o *ConsoleSettings) GetSyncWithReplaceAllowedOk() (*bool, bool)`

GetSyncWithReplaceAllowedOk returns a tuple with the SyncWithReplaceAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncWithReplaceAllowed

`func (o *ConsoleSettings) SetSyncWithReplaceAllowed(v bool)`

SetSyncWithReplaceAllowed sets SyncWithReplaceAllowed field to given value.

### HasSyncWithReplaceAllowed

`func (o *ConsoleSettings) HasSyncWithReplaceAllowed() bool`

HasSyncWithReplaceAllowed returns a boolean if a field has been set.

### GetUiBannerContent

`func (o *ConsoleSettings) GetUiBannerContent() string`

GetUiBannerContent returns the UiBannerContent field if non-nil, zero value otherwise.

### GetUiBannerContentOk

`func (o *ConsoleSettings) GetUiBannerContentOk() (*string, bool)`

GetUiBannerContentOk returns a tuple with the UiBannerContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiBannerContent

`func (o *ConsoleSettings) SetUiBannerContent(v string)`

SetUiBannerContent sets UiBannerContent field to given value.

### HasUiBannerContent

`func (o *ConsoleSettings) HasUiBannerContent() bool`

HasUiBannerContent returns a boolean if a field has been set.

### GetUiCssURL

`func (o *ConsoleSettings) GetUiCssURL() string`

GetUiCssURL returns the UiCssURL field if non-nil, zero value otherwise.

### GetUiCssURLOk

`func (o *ConsoleSettings) GetUiCssURLOk() (*string, bool)`

GetUiCssURLOk returns a tuple with the UiCssURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiCssURL

`func (o *ConsoleSettings) SetUiCssURL(v string)`

SetUiCssURL sets UiCssURL field to given value.

### HasUiCssURL

`func (o *ConsoleSettings) HasUiCssURL() bool`

HasUiCssURL returns a boolean if a field has been set.

### GetUrl

`func (o *ConsoleSettings) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConsoleSettings) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConsoleSettings) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ConsoleSettings) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserLoginsDisabled

`func (o *ConsoleSettings) GetUserLoginsDisabled() bool`

GetUserLoginsDisabled returns the UserLoginsDisabled field if non-nil, zero value otherwise.

### GetUserLoginsDisabledOk

`func (o *ConsoleSettings) GetUserLoginsDisabledOk() (*bool, bool)`

GetUserLoginsDisabledOk returns a tuple with the UserLoginsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLoginsDisabled

`func (o *ConsoleSettings) SetUserLoginsDisabled(v bool)`

SetUserLoginsDisabled sets UserLoginsDisabled field to given value.

### HasUserLoginsDisabled

`func (o *ConsoleSettings) HasUserLoginsDisabled() bool`

HasUserLoginsDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


