# DbCreateEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | [**DbEndpointCreate**](DbEndpointCreate.md) |  | 

## Methods

### NewDbCreateEndpointRequest

`func NewDbCreateEndpointRequest(endpoint DbEndpointCreate, ) *DbCreateEndpointRequest`

NewDbCreateEndpointRequest instantiates a new DbCreateEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbCreateEndpointRequestWithDefaults

`func NewDbCreateEndpointRequestWithDefaults() *DbCreateEndpointRequest`

NewDbCreateEndpointRequestWithDefaults instantiates a new DbCreateEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *DbCreateEndpointRequest) GetEndpoint() DbEndpointCreate`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *DbCreateEndpointRequest) GetEndpointOk() (*DbEndpointCreate, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *DbCreateEndpointRequest) SetEndpoint(v DbEndpointCreate)`

SetEndpoint sets Endpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


