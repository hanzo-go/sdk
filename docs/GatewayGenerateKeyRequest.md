# GatewayGenerateKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyAlias** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **string** | Duration like \&quot;30d\&quot;, \&quot;1h\&quot;, etc. | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**MaxBudget** | Pointer to **float32** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGatewayGenerateKeyRequest

`func NewGatewayGenerateKeyRequest() *GatewayGenerateKeyRequest`

NewGatewayGenerateKeyRequest instantiates a new GatewayGenerateKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayGenerateKeyRequestWithDefaults

`func NewGatewayGenerateKeyRequestWithDefaults() *GatewayGenerateKeyRequest`

NewGatewayGenerateKeyRequestWithDefaults instantiates a new GatewayGenerateKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyAlias

`func (o *GatewayGenerateKeyRequest) GetKeyAlias() string`

GetKeyAlias returns the KeyAlias field if non-nil, zero value otherwise.

### GetKeyAliasOk

`func (o *GatewayGenerateKeyRequest) GetKeyAliasOk() (*string, bool)`

GetKeyAliasOk returns a tuple with the KeyAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlias

`func (o *GatewayGenerateKeyRequest) SetKeyAlias(v string)`

SetKeyAlias sets KeyAlias field to given value.

### HasKeyAlias

`func (o *GatewayGenerateKeyRequest) HasKeyAlias() bool`

HasKeyAlias returns a boolean if a field has been set.

### GetDuration

`func (o *GatewayGenerateKeyRequest) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GatewayGenerateKeyRequest) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GatewayGenerateKeyRequest) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GatewayGenerateKeyRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetModels

`func (o *GatewayGenerateKeyRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *GatewayGenerateKeyRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *GatewayGenerateKeyRequest) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *GatewayGenerateKeyRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetMaxBudget

`func (o *GatewayGenerateKeyRequest) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *GatewayGenerateKeyRequest) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *GatewayGenerateKeyRequest) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.

### HasMaxBudget

`func (o *GatewayGenerateKeyRequest) HasMaxBudget() bool`

HasMaxBudget returns a boolean if a field has been set.

### GetUserId

`func (o *GatewayGenerateKeyRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GatewayGenerateKeyRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GatewayGenerateKeyRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GatewayGenerateKeyRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTeamId

`func (o *GatewayGenerateKeyRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *GatewayGenerateKeyRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *GatewayGenerateKeyRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *GatewayGenerateKeyRequest) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetMetadata

`func (o *GatewayGenerateKeyRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GatewayGenerateKeyRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GatewayGenerateKeyRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GatewayGenerateKeyRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


