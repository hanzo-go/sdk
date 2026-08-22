# MCPError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewMCPError

`func NewMCPError() *MCPError`

NewMCPError instantiates a new MCPError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPErrorWithDefaults

`func NewMCPErrorWithDefaults() *MCPError`

NewMCPErrorWithDefaults instantiates a new MCPError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MCPError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MCPError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MCPError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *MCPError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *MCPError) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MCPError) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MCPError) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *MCPError) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *MCPError) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *MCPError) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessage

`func (o *MCPError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MCPError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MCPError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MCPError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


