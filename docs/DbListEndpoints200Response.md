# DbListEndpoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]DbEndpoint**](DbEndpoint.md) |  | [optional] 

## Methods

### NewDbListEndpoints200Response

`func NewDbListEndpoints200Response() *DbListEndpoints200Response`

NewDbListEndpoints200Response instantiates a new DbListEndpoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbListEndpoints200ResponseWithDefaults

`func NewDbListEndpoints200ResponseWithDefaults() *DbListEndpoints200Response`

NewDbListEndpoints200ResponseWithDefaults instantiates a new DbListEndpoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *DbListEndpoints200Response) GetEndpoints() []DbEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *DbListEndpoints200Response) GetEndpointsOk() (*[]DbEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *DbListEndpoints200Response) SetEndpoints(v []DbEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *DbListEndpoints200Response) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


