# SandboxWrote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int64** | Bytes is how many bytes the file now holds. A write REPLACES the file, so this is its whole length and not an amount appended, and 0 is a legitimate answer: a WriteIn with no Data truncates the file to nothing. | [optional] 
**Path** | Pointer to **string** | Path is where the bytes actually landed: the caller&#39;s path resolved against the sandbox&#39;s working directory (Leased.Workdir), which is what a later read or a shell line inside the sandbox has to name. | [optional] 

## Methods

### NewSandboxWrote

`func NewSandboxWrote() *SandboxWrote`

NewSandboxWrote instantiates a new SandboxWrote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxWroteWithDefaults

`func NewSandboxWroteWithDefaults() *SandboxWrote`

NewSandboxWroteWithDefaults instantiates a new SandboxWrote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *SandboxWrote) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *SandboxWrote) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *SandboxWrote) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *SandboxWrote) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetPath

`func (o *SandboxWrote) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SandboxWrote) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SandboxWrote) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SandboxWrote) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


