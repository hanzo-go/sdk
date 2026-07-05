# ExecExecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lang** | **string** | Language / runtime (e.g. python, node) | 
**Code** | **string** |  | 
**Files** | Pointer to [**[]ExecExecFile**](ExecExecFile.md) | Optional input files | [optional] 

## Methods

### NewExecExecRequest

`func NewExecExecRequest(lang string, code string, ) *ExecExecRequest`

NewExecExecRequest instantiates a new ExecExecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecExecRequestWithDefaults

`func NewExecExecRequestWithDefaults() *ExecExecRequest`

NewExecExecRequestWithDefaults instantiates a new ExecExecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLang

`func (o *ExecExecRequest) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *ExecExecRequest) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *ExecExecRequest) SetLang(v string)`

SetLang sets Lang field to given value.


### GetCode

`func (o *ExecExecRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ExecExecRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ExecExecRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetFiles

`func (o *ExecExecRequest) GetFiles() []ExecExecFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ExecExecRequest) GetFilesOk() (*[]ExecExecFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ExecExecRequest) SetFiles(v []ExecExecFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ExecExecRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


