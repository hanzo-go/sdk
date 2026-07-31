# CloudToolCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to **map[string]map[string]interface{}** | Arguments is the tool&#39;s own input object, passed through verbatim to whichever source owns it. | [optional] 
**Name** | Pointer to **string** | Name is the tool to run, exactly as GET /v1/tools reports it. | [optional] 

## Methods

### NewCloudToolCall

`func NewCloudToolCall() *CloudToolCall`

NewCloudToolCall instantiates a new CloudToolCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudToolCallWithDefaults

`func NewCloudToolCallWithDefaults() *CloudToolCall`

NewCloudToolCallWithDefaults instantiates a new CloudToolCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *CloudToolCall) GetArguments() map[string]map[string]interface{}`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *CloudToolCall) GetArgumentsOk() (*map[string]map[string]interface{}, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *CloudToolCall) SetArguments(v map[string]map[string]interface{})`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *CloudToolCall) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetName

`func (o *CloudToolCall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudToolCall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudToolCall) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudToolCall) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


