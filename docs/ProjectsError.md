# ProjectsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Always \&quot;error\&quot; for error responses. | 
**Msg** | **string** | Human-readable error message. | 
**Code** | Pointer to **int32** | Application-specific error code. | [optional] 
**RequestId** | Pointer to **string** | Request ID for debugging. | [optional] 

## Methods

### NewProjectsError

`func NewProjectsError(status string, msg string, ) *ProjectsError`

NewProjectsError instantiates a new ProjectsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsErrorWithDefaults

`func NewProjectsErrorWithDefaults() *ProjectsError`

NewProjectsErrorWithDefaults instantiates a new ProjectsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProjectsError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProjectsError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProjectsError) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMsg

`func (o *ProjectsError) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ProjectsError) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ProjectsError) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetCode

`func (o *ProjectsError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProjectsError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProjectsError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProjectsError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRequestId

`func (o *ProjectsError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProjectsError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProjectsError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ProjectsError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


