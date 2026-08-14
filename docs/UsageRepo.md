# UsageRepo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the repo&#39;s org-unique handle. | [optional] 
**Project** | Pointer to **string** | Project is the sub-scope the repo lives in; absent for the default scope. | [optional] 
**SizeBytes** | Pointer to **int32** | SizeBytes is the repo&#39;s on-disk size at its last measurement. | [optional] 

## Methods

### NewUsageRepo

`func NewUsageRepo() *UsageRepo`

NewUsageRepo instantiates a new UsageRepo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageRepoWithDefaults

`func NewUsageRepoWithDefaults() *UsageRepo`

NewUsageRepoWithDefaults instantiates a new UsageRepo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UsageRepo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsageRepo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsageRepo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsageRepo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *UsageRepo) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *UsageRepo) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *UsageRepo) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *UsageRepo) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSizeBytes

`func (o *UsageRepo) GetSizeBytes() int32`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *UsageRepo) GetSizeBytesOk() (*int32, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *UsageRepo) SetSizeBytes(v int32)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *UsageRepo) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


