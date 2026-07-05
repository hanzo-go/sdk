# EngineListServingEndpoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]EngineServingEndpoint**](EngineServingEndpoint.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineListServingEndpoints200Response

`func NewEngineListServingEndpoints200Response() *EngineListServingEndpoints200Response`

NewEngineListServingEndpoints200Response instantiates a new EngineListServingEndpoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineListServingEndpoints200ResponseWithDefaults

`func NewEngineListServingEndpoints200ResponseWithDefaults() *EngineListServingEndpoints200Response`

NewEngineListServingEndpoints200ResponseWithDefaults instantiates a new EngineListServingEndpoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *EngineListServingEndpoints200Response) GetEndpoints() []EngineServingEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *EngineListServingEndpoints200Response) GetEndpointsOk() (*[]EngineServingEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *EngineListServingEndpoints200Response) SetEndpoints(v []EngineServingEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *EngineListServingEndpoints200Response) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetTotal

`func (o *EngineListServingEndpoints200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *EngineListServingEndpoints200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *EngineListServingEndpoints200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *EngineListServingEndpoints200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


