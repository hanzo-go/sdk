# SandboxWriteIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | Data is the file&#39;s bytes, and replaces whatever was there. | [optional] 
**Id** | Pointer to **string** | ID is the sandbox to write into, from an earlier lease. | [optional] 
**Path** | Pointer to **string** | Path is confined the same way PathIn.Path is. Missing parent directories are created. | [optional] 

## Methods

### NewSandboxWriteIn

`func NewSandboxWriteIn() *SandboxWriteIn`

NewSandboxWriteIn instantiates a new SandboxWriteIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxWriteInWithDefaults

`func NewSandboxWriteInWithDefaults() *SandboxWriteIn`

NewSandboxWriteInWithDefaults instantiates a new SandboxWriteIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SandboxWriteIn) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SandboxWriteIn) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SandboxWriteIn) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SandboxWriteIn) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *SandboxWriteIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxWriteIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxWriteIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxWriteIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *SandboxWriteIn) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SandboxWriteIn) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SandboxWriteIn) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SandboxWriteIn) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


