# SandboxPathIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the sandbox to read from, from an earlier lease. | [optional] 
**Path** | Pointer to **string** | Path is read relative to the sandbox&#39;s working directory unless it is absolute, and a path that climbs out of it is refused rather than rewritten. Empty names the working directory itself, which lists it. | [optional] 

## Methods

### NewSandboxPathIn

`func NewSandboxPathIn() *SandboxPathIn`

NewSandboxPathIn instantiates a new SandboxPathIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxPathInWithDefaults

`func NewSandboxPathInWithDefaults() *SandboxPathIn`

NewSandboxPathInWithDefaults instantiates a new SandboxPathIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SandboxPathIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SandboxPathIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SandboxPathIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SandboxPathIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *SandboxPathIn) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SandboxPathIn) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SandboxPathIn) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SandboxPathIn) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


