# ExecCodeFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the file&#39;s path RELATIVE to its session&#39;s artifact directory, which is also how it is fetched: GET /v1/exec/download/{session}/{id}. | [optional] 
**Name** | Pointer to **string** | Name is the display name. On an ANSWER it carries the &#x60;{session}/{id}&#x60; identifier whole, because the client matches on that prefix. | [optional] 
**SessionId** | Pointer to **string** | SessionID is the other accepted spelling of the same fact on the way IN. Both are read; whichever is set wins. | [optional] 
**StorageSessionId** | Pointer to **string** | StorageSessionID names the session holding the bytes, and is the spelling the answer always uses. | [optional] 

## Methods

### NewExecCodeFile

`func NewExecCodeFile() *ExecCodeFile`

NewExecCodeFile instantiates a new ExecCodeFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecCodeFileWithDefaults

`func NewExecCodeFileWithDefaults() *ExecCodeFile`

NewExecCodeFileWithDefaults instantiates a new ExecCodeFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecCodeFile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecCodeFile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecCodeFile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExecCodeFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExecCodeFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExecCodeFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExecCodeFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExecCodeFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSessionId

`func (o *ExecCodeFile) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ExecCodeFile) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ExecCodeFile) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ExecCodeFile) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStorageSessionId

`func (o *ExecCodeFile) GetStorageSessionId() string`

GetStorageSessionId returns the StorageSessionId field if non-nil, zero value otherwise.

### GetStorageSessionIdOk

`func (o *ExecCodeFile) GetStorageSessionIdOk() (*string, bool)`

GetStorageSessionIdOk returns a tuple with the StorageSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSessionId

`func (o *ExecCodeFile) SetStorageSessionId(v string)`

SetStorageSessionId sets StorageSessionId field to given value.

### HasStorageSessionId

`func (o *ExecCodeFile) HasStorageSessionId() bool`

HasStorageSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


