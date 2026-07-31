# SecurityScanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** | Optional sub-scope, ^[A-Za-z0-9][A-Za-z0-9._-]{0,127}$ | [optional] 
**Files** | [**[]SecurityFileInput**](SecurityFileInput.md) | At least one file; max 500 files / 8 MiB total | 

## Methods

### NewSecurityScanRequest

`func NewSecurityScanRequest(files []SecurityFileInput, ) *SecurityScanRequest`

NewSecurityScanRequest instantiates a new SecurityScanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityScanRequestWithDefaults

`func NewSecurityScanRequestWithDefaults() *SecurityScanRequest`

NewSecurityScanRequestWithDefaults instantiates a new SecurityScanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *SecurityScanRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SecurityScanRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SecurityScanRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SecurityScanRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetFiles

`func (o *SecurityScanRequest) GetFiles() []SecurityFileInput`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SecurityScanRequest) GetFilesOk() (*[]SecurityFileInput, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SecurityScanRequest) SetFiles(v []SecurityFileInput)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


