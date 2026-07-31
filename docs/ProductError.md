# ProductError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Always \&quot;error\&quot; for error responses | 
**Msg** | **string** | Human-readable error message | 
**Code** | Pointer to **int32** | Application-specific error code | [optional] 
**Data** | Pointer to **map[string]interface{}** | Additional error context | [optional] 
**Data2** | Pointer to **map[string]interface{}** | Additional error details | [optional] 
**RequestId** | Pointer to **string** | Request ID for debugging | [optional] 

## Methods

### NewProductError

`func NewProductError(status string, msg string, ) *ProductError`

NewProductError instantiates a new ProductError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductErrorWithDefaults

`func NewProductErrorWithDefaults() *ProductError`

NewProductErrorWithDefaults instantiates a new ProductError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProductError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductError) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMsg

`func (o *ProductError) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ProductError) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ProductError) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetCode

`func (o *ProductError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProductError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProductError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProductError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *ProductError) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProductError) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProductError) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ProductError) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *ProductError) GetData2() map[string]interface{}`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *ProductError) GetData2Ok() (*map[string]interface{}, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *ProductError) SetData2(v map[string]interface{})`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *ProductError) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### GetRequestId

`func (o *ProductError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProductError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProductError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ProductError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


