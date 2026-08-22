# CatalogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** | Configured is whether THIS DEPLOYMENT holds the OAuth client credentials for the provider. False means Connect would dead-end, so the console can offer it disabled instead of broken. It is deployment-wide and says nothing about whether the caller&#39;s org has connected the source — that is the connector list&#39;s &#x60;status&#x60;. | [optional] 
**Description** | Pointer to **string** | Description is one line of shop copy: what connecting this source pulls in. Native connectors carry written prose; a piece-backed one reads \&quot;activepieces connector (&lt;piece&gt;)\&quot;. | [optional] 
**DisplayName** | Pointer to **string** | DisplayName is the label to show a person. First-party connectors carry a written name (\&quot;GitHub\&quot;, \&quot;Google Drive\&quot;); a piece-backed one falls back to the provider capitalized, because the rich activepieces metadata lives behind a cross-service call this read will not make. | [optional] 
**Kind** | Pointer to **string** | \&quot;native\&quot; | \&quot;piece\&quot; | [optional] 
**Provider** | Pointer to **string** | Provider is the source&#39;s id and the address every connector op takes it by (/v1/knowledge/connectors/:provider). One of github, slack, google, notion. | [optional] 

## Methods

### NewCatalogEntry

`func NewCatalogEntry() *CatalogEntry`

NewCatalogEntry instantiates a new CatalogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogEntryWithDefaults

`func NewCatalogEntryWithDefaults() *CatalogEntry`

NewCatalogEntryWithDefaults instantiates a new CatalogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *CatalogEntry) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *CatalogEntry) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *CatalogEntry) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *CatalogEntry) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CatalogEntry) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CatalogEntry) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CatalogEntry) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CatalogEntry) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetKind

`func (o *CatalogEntry) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CatalogEntry) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CatalogEntry) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CatalogEntry) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetProvider

`func (o *CatalogEntry) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CatalogEntry) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CatalogEntry) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CatalogEntry) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


