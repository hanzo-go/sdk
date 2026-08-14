# EndpointList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Endpoint**](Endpoint.md) | Data is the org&#39;s endpoints, newest first, each with its signing secret REDACTED — the secret leaves the server only on create and on rotate. | [optional] 

## Methods

### NewEndpointList

`func NewEndpointList() *EndpointList`

NewEndpointList instantiates a new EndpointList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointListWithDefaults

`func NewEndpointListWithDefaults() *EndpointList`

NewEndpointListWithDefaults instantiates a new EndpointList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EndpointList) GetData() []Endpoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EndpointList) GetDataOk() (*[]Endpoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EndpointList) SetData(v []Endpoint)`

SetData sets Data field to given value.

### HasData

`func (o *EndpointList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


