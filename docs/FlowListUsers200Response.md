# FlowListUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]FlowUser**](FlowUser.md) |  | [optional] 
**Next** | Pointer to **string** |  | [optional] 
**Previous** | Pointer to **string** |  | [optional] 

## Methods

### NewFlowListUsers200Response

`func NewFlowListUsers200Response() *FlowListUsers200Response`

NewFlowListUsers200Response instantiates a new FlowListUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowListUsers200ResponseWithDefaults

`func NewFlowListUsers200ResponseWithDefaults() *FlowListUsers200Response`

NewFlowListUsers200ResponseWithDefaults instantiates a new FlowListUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FlowListUsers200Response) GetData() []FlowUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FlowListUsers200Response) GetDataOk() (*[]FlowUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FlowListUsers200Response) SetData(v []FlowUser)`

SetData sets Data field to given value.

### HasData

`func (o *FlowListUsers200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNext

`func (o *FlowListUsers200Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FlowListUsers200Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FlowListUsers200Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *FlowListUsers200Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *FlowListUsers200Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *FlowListUsers200Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *FlowListUsers200Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *FlowListUsers200Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


