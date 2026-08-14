# SubmitReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]Scan**](Scan.md) | Files is the batch to scan, at most 500 files and 8 MiB of content in total. | 
**Project** | Pointer to **string** | Project names the sub-scope the scan is filed under. It must be a slug; omit it and the caller&#39;s project header is used instead. | [optional] 

## Methods

### NewSubmitReq

`func NewSubmitReq(files []Scan, ) *SubmitReq`

NewSubmitReq instantiates a new SubmitReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitReqWithDefaults

`func NewSubmitReqWithDefaults() *SubmitReq`

NewSubmitReqWithDefaults instantiates a new SubmitReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *SubmitReq) GetFiles() []Scan`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *SubmitReq) GetFilesOk() (*[]Scan, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *SubmitReq) SetFiles(v []Scan)`

SetFiles sets Files field to given value.


### GetProject

`func (o *SubmitReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SubmitReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SubmitReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SubmitReq) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


