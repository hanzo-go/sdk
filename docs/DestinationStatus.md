# DestinationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the operator&#39;s own label for the connected account, as supplied on connect. Absent when unset. | [optional] 
**Category** | Pointer to **string** | groups the card: Analytics | Advertising | [optional] 
**Config** | Pointer to **map[string]string** | Config is the org&#39;s stored NON-SECRET configuration — the measurement/pixel ids keyed by DestinationField.Key. A secret is never in here; secrets live in KMS and only their names are published, in Secrets. | [optional] 
**Connected** | Pointer to **bool** | Connected is true when this org has a stored row for the platform — it has been configured here at least once. It says nothing about whether a credential still resolves; that is Live. | [optional] 
**Enabled** | Pointer to **bool** | Enabled is whether the fan-out forwards to this destination. False on a destination that is connected but paused, and on one never connected. | [optional] 
**Fields** | Pointer to [**[]DestinationField**](DestinationField.md) | Fields are the non-secret inputs this platform needs, which the console card renders and the connect body fills. | [optional] 
**Live** | Pointer to **bool** | Live is whether a credential resolves RIGHT NOW: a KMS-sealed secret for this org, else the integrations connection named by the platform&#39;s Fallback, else no credential needed at all (a public-ingest sink like Analytics). False on a connected destination whose secret has gone missing — Connected &amp;&amp; !Live is exactly the \&quot;reconnect me\&quot; state. | [optional] 
**Name** | Pointer to **string** | the platform&#39;s display name (\&quot;Google Analytics 4\&quot;) | [optional] 
**Pixel** | Pointer to **bool** | Pixel is whether the hosted tag can inject a browser pixel for this platform, so a console offers a per-SITE pixel input for exactly these. False means the platform receives conversions server-side only, and an input would promise an injection that never happens. Derived from the tag&#39;s own map (event.BrowserTags), never restated — a second list is how a console offers a pixel nothing fires. | [optional] 
**Platform** | Pointer to **string** | the platform slug, and the path segment every route addresses it by | [optional] 
**Secrets** | Pointer to **[]string** | Secrets are the KMS secret NAMES this platform custodies for the org — names only, never values. The connect body accepts each under its camelCase form. | [optional] 

## Methods

### NewDestinationStatus

`func NewDestinationStatus() *DestinationStatus`

NewDestinationStatus instantiates a new DestinationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationStatusWithDefaults

`func NewDestinationStatusWithDefaults() *DestinationStatus`

NewDestinationStatusWithDefaults instantiates a new DestinationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *DestinationStatus) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *DestinationStatus) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *DestinationStatus) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *DestinationStatus) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCategory

`func (o *DestinationStatus) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DestinationStatus) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DestinationStatus) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DestinationStatus) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *DestinationStatus) GetConfig() map[string]string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *DestinationStatus) GetConfigOk() (*map[string]string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *DestinationStatus) SetConfig(v map[string]string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *DestinationStatus) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnected

`func (o *DestinationStatus) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *DestinationStatus) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *DestinationStatus) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *DestinationStatus) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DestinationStatus) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFields

`func (o *DestinationStatus) GetFields() []DestinationField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DestinationStatus) GetFieldsOk() (*[]DestinationField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DestinationStatus) SetFields(v []DestinationField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *DestinationStatus) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLive

`func (o *DestinationStatus) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *DestinationStatus) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *DestinationStatus) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *DestinationStatus) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetName

`func (o *DestinationStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPixel

`func (o *DestinationStatus) GetPixel() bool`

GetPixel returns the Pixel field if non-nil, zero value otherwise.

### GetPixelOk

`func (o *DestinationStatus) GetPixelOk() (*bool, bool)`

GetPixelOk returns a tuple with the Pixel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixel

`func (o *DestinationStatus) SetPixel(v bool)`

SetPixel sets Pixel field to given value.

### HasPixel

`func (o *DestinationStatus) HasPixel() bool`

HasPixel returns a boolean if a field has been set.

### GetPlatform

`func (o *DestinationStatus) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DestinationStatus) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DestinationStatus) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DestinationStatus) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSecrets

`func (o *DestinationStatus) GetSecrets() []string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *DestinationStatus) GetSecretsOk() (*[]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *DestinationStatus) SetSecrets(v []string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *DestinationStatus) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


