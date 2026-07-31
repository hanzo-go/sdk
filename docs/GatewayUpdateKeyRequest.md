# GatewayUpdateKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Models** | Pointer to **[]string** |  | [optional] 
**MaxBudget** | Pointer to **float32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGatewayUpdateKeyRequest

`func NewGatewayUpdateKeyRequest(key string, ) *GatewayUpdateKeyRequest`

NewGatewayUpdateKeyRequest instantiates a new GatewayUpdateKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateKeyRequestWithDefaults

`func NewGatewayUpdateKeyRequestWithDefaults() *GatewayUpdateKeyRequest`

NewGatewayUpdateKeyRequestWithDefaults instantiates a new GatewayUpdateKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GatewayUpdateKeyRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GatewayUpdateKeyRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GatewayUpdateKeyRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetModels

`func (o *GatewayUpdateKeyRequest) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *GatewayUpdateKeyRequest) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *GatewayUpdateKeyRequest) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *GatewayUpdateKeyRequest) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetMaxBudget

`func (o *GatewayUpdateKeyRequest) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *GatewayUpdateKeyRequest) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *GatewayUpdateKeyRequest) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.

### HasMaxBudget

`func (o *GatewayUpdateKeyRequest) HasMaxBudget() bool`

HasMaxBudget returns a boolean if a field has been set.

### GetMetadata

`func (o *GatewayUpdateKeyRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GatewayUpdateKeyRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GatewayUpdateKeyRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GatewayUpdateKeyRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


