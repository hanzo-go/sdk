# CloudDestinationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the operator&#39;s own label for the connected account, as supplied on connect. Absent when unset. | [optional] 
**Category** | Pointer to **string** | groups the card: Analytics | Advertising | [optional] 
**Config** | Pointer to **map[string]string** | Config is the org&#39;s stored NON-SECRET configuration — the measurement/pixel ids keyed by DestinationField.Key. A secret is never in here; secrets live in KMS and only their names are published, in Secrets. | [optional] 
**Connected** | Pointer to **bool** | Connected is true when this org has a stored row for the platform — it has been configured here at least once. It says nothing about whether a credential still resolves; that is Live. | [optional] 
**Enabled** | Pointer to **bool** | Enabled is whether the fan-out forwards to this destination. False on a destination that is connected but paused, and on one never connected. | [optional] 
**Fields** | Pointer to [**[]CloudDestinationField**](CloudDestinationField.md) | Fields are the non-secret inputs this platform needs, which the console card renders and the connect body fills. | [optional] 
**Live** | Pointer to **bool** | Live is whether a credential resolves RIGHT NOW: a KMS-sealed secret for this org, else the integrations connection named by the platform&#39;s Fallback, else no credential needed at all (a public-ingest sink like Umami). False on a connected destination whose secret has gone missing — Connected &amp;&amp; !Live is exactly the \&quot;reconnect me\&quot; state. | [optional] 
**Name** | Pointer to **string** | the platform&#39;s display name (\&quot;Google Analytics 4\&quot;) | [optional] 
**Platform** | Pointer to **string** | the platform slug, and the path segment every route addresses it by | [optional] 
**Secrets** | Pointer to **[]string** | Secrets are the KMS secret NAMES this platform custodies for the org — names only, never values. The connect body accepts each under its camelCase form. | [optional] 

## Methods

### NewCloudDestinationStatus

`func NewCloudDestinationStatus() *CloudDestinationStatus`

NewCloudDestinationStatus instantiates a new CloudDestinationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDestinationStatusWithDefaults

`func NewCloudDestinationStatusWithDefaults() *CloudDestinationStatus`

NewCloudDestinationStatusWithDefaults instantiates a new CloudDestinationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudDestinationStatus) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudDestinationStatus) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudDestinationStatus) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudDestinationStatus) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCategory

`func (o *CloudDestinationStatus) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudDestinationStatus) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudDestinationStatus) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudDestinationStatus) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *CloudDestinationStatus) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CloudDestinationStatus) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CloudDestinationStatus) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CloudDestinationStatus) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnected

`func (o *CloudDestinationStatus) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *CloudDestinationStatus) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *CloudDestinationStatus) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *CloudDestinationStatus) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudDestinationStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudDestinationStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudDestinationStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudDestinationStatus) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFields

`func (o *CloudDestinationStatus) GetFields() []CloudDestinationField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CloudDestinationStatus) GetFieldsOk() (*[]CloudDestinationField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CloudDestinationStatus) SetFields(v []CloudDestinationField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CloudDestinationStatus) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLive

`func (o *CloudDestinationStatus) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *CloudDestinationStatus) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *CloudDestinationStatus) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *CloudDestinationStatus) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetName

`func (o *CloudDestinationStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudDestinationStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudDestinationStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudDestinationStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *CloudDestinationStatus) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CloudDestinationStatus) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CloudDestinationStatus) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CloudDestinationStatus) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSecrets

`func (o *CloudDestinationStatus) GetSecrets() []string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *CloudDestinationStatus) GetSecretsOk() (*[]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *CloudDestinationStatus) SetSecrets(v []string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *CloudDestinationStatus) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


