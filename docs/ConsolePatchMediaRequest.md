# ConsolePatchMediaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UploadedAt** | **time.Time** |  | 
**UploadHttpStatus** | **int32** |  | 
**UploadHttpError** | Pointer to **string** |  | [optional] 
**UploadTimeMs** | Pointer to **int32** |  | [optional] 

## Methods

### NewConsolePatchMediaRequest

`func NewConsolePatchMediaRequest(uploadedAt time.Time, uploadHttpStatus int32, ) *ConsolePatchMediaRequest`

NewConsolePatchMediaRequest instantiates a new ConsolePatchMediaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsolePatchMediaRequestWithDefaults

`func NewConsolePatchMediaRequestWithDefaults() *ConsolePatchMediaRequest`

NewConsolePatchMediaRequestWithDefaults instantiates a new ConsolePatchMediaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploadedAt

`func (o *ConsolePatchMediaRequest) GetUploadedAt() time.Time`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *ConsolePatchMediaRequest) GetUploadedAtOk() (*time.Time, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *ConsolePatchMediaRequest) SetUploadedAt(v time.Time)`

SetUploadedAt sets UploadedAt field to given value.


### GetUploadHttpStatus

`func (o *ConsolePatchMediaRequest) GetUploadHttpStatus() int32`

GetUploadHttpStatus returns the UploadHttpStatus field if non-nil, zero value otherwise.

### GetUploadHttpStatusOk

`func (o *ConsolePatchMediaRequest) GetUploadHttpStatusOk() (*int32, bool)`

GetUploadHttpStatusOk returns a tuple with the UploadHttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadHttpStatus

`func (o *ConsolePatchMediaRequest) SetUploadHttpStatus(v int32)`

SetUploadHttpStatus sets UploadHttpStatus field to given value.


### GetUploadHttpError

`func (o *ConsolePatchMediaRequest) GetUploadHttpError() string`

GetUploadHttpError returns the UploadHttpError field if non-nil, zero value otherwise.

### GetUploadHttpErrorOk

`func (o *ConsolePatchMediaRequest) GetUploadHttpErrorOk() (*string, bool)`

GetUploadHttpErrorOk returns a tuple with the UploadHttpError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadHttpError

`func (o *ConsolePatchMediaRequest) SetUploadHttpError(v string)`

SetUploadHttpError sets UploadHttpError field to given value.

### HasUploadHttpError

`func (o *ConsolePatchMediaRequest) HasUploadHttpError() bool`

HasUploadHttpError returns a boolean if a field has been set.

### GetUploadTimeMs

`func (o *ConsolePatchMediaRequest) GetUploadTimeMs() int32`

GetUploadTimeMs returns the UploadTimeMs field if non-nil, zero value otherwise.

### GetUploadTimeMsOk

`func (o *ConsolePatchMediaRequest) GetUploadTimeMsOk() (*int32, bool)`

GetUploadTimeMsOk returns a tuple with the UploadTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTimeMs

`func (o *ConsolePatchMediaRequest) SetUploadTimeMs(v int32)`

SetUploadTimeMs sets UploadTimeMs field to given value.

### HasUploadTimeMs

`func (o *ConsolePatchMediaRequest) HasUploadTimeMs() bool`

HasUploadTimeMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


