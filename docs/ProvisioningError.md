# ProvisioningError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Msg** | **string** | Human-readable error message. | 
**Code** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 

## Methods

### NewProvisioningError

`func NewProvisioningError(status string, msg string, ) *ProvisioningError`

NewProvisioningError instantiates a new ProvisioningError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningErrorWithDefaults

`func NewProvisioningErrorWithDefaults() *ProvisioningError`

NewProvisioningErrorWithDefaults instantiates a new ProvisioningError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProvisioningError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisioningError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisioningError) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMsg

`func (o *ProvisioningError) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ProvisioningError) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ProvisioningError) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetCode

`func (o *ProvisioningError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProvisioningError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProvisioningError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProvisioningError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRequestId

`func (o *ProvisioningError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProvisioningError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProvisioningError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ProvisioningError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


