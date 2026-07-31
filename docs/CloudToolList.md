# CloudToolList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | Pointer to [**[]CloudTool**](CloudTool.md) | Tools is every tool the caller may see, deduplicated by name with source precedence applied. | [optional] 

## Methods

### NewCloudToolList

`func NewCloudToolList() *CloudToolList`

NewCloudToolList instantiates a new CloudToolList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudToolListWithDefaults

`func NewCloudToolListWithDefaults() *CloudToolList`

NewCloudToolListWithDefaults instantiates a new CloudToolList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *CloudToolList) GetTools() []CloudTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CloudToolList) GetToolsOk() (*[]CloudTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CloudToolList) SetTools(v []CloudTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CloudToolList) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


