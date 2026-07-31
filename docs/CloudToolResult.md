# CloudToolResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the tool that ran. | [optional] 
**Result** | Pointer to **map[string]interface{}** | Result is the tool&#39;s own output, verbatim — its shape is the tool&#39;s, not this plane&#39;s. | [optional] 

## Methods

### NewCloudToolResult

`func NewCloudToolResult() *CloudToolResult`

NewCloudToolResult instantiates a new CloudToolResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudToolResultWithDefaults

`func NewCloudToolResultWithDefaults() *CloudToolResult`

NewCloudToolResultWithDefaults instantiates a new CloudToolResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudToolResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudToolResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudToolResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudToolResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *CloudToolResult) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CloudToolResult) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CloudToolResult) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *CloudToolResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


