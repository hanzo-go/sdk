# O11yO11yTopLevelOpsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s end, epoch nanoseconds as a string; empty means unbounded. | [optional] 
**Service** | Pointer to **string** | Service narrows the map to one service when set. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s start, epoch nanoseconds as a string; empty means unbounded. | [optional] 

## Methods

### NewO11yO11yTopLevelOpsIn

`func NewO11yO11yTopLevelOpsIn() *O11yO11yTopLevelOpsIn`

NewO11yO11yTopLevelOpsIn instantiates a new O11yO11yTopLevelOpsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTopLevelOpsInWithDefaults

`func NewO11yO11yTopLevelOpsInWithDefaults() *O11yO11yTopLevelOpsIn`

NewO11yO11yTopLevelOpsInWithDefaults instantiates a new O11yO11yTopLevelOpsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yTopLevelOpsIn) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yTopLevelOpsIn) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yTopLevelOpsIn) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yO11yTopLevelOpsIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetService

`func (o *O11yO11yTopLevelOpsIn) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *O11yO11yTopLevelOpsIn) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *O11yO11yTopLevelOpsIn) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *O11yO11yTopLevelOpsIn) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11yTopLevelOpsIn) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yTopLevelOpsIn) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yTopLevelOpsIn) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yO11yTopLevelOpsIn) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


