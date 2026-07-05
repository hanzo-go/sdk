# ExecUploadFile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]ExecExecFile**](ExecExecFile.md) |  | [optional] 

## Methods

### NewExecUploadFile200Response

`func NewExecUploadFile200Response() *ExecUploadFile200Response`

NewExecUploadFile200Response instantiates a new ExecUploadFile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecUploadFile200ResponseWithDefaults

`func NewExecUploadFile200ResponseWithDefaults() *ExecUploadFile200Response`

NewExecUploadFile200ResponseWithDefaults instantiates a new ExecUploadFile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *ExecUploadFile200Response) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ExecUploadFile200Response) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ExecUploadFile200Response) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ExecUploadFile200Response) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetFiles

`func (o *ExecUploadFile200Response) GetFiles() []ExecExecFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ExecUploadFile200Response) GetFilesOk() (*[]ExecExecFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ExecUploadFile200Response) SetFiles(v []ExecExecFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ExecUploadFile200Response) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


