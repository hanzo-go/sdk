# PathIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the sandbox to read from, from an earlier lease. | [optional] 
**Path** | Pointer to **string** | Path is read relative to the sandbox&#39;s working directory unless it is absolute, and a path that climbs out of it is refused rather than rewritten. Empty names the working directory itself, which lists it. | [optional] 

## Methods

### NewPathIn

`func NewPathIn() *PathIn`

NewPathIn instantiates a new PathIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathInWithDefaults

`func NewPathInWithDefaults() *PathIn`

NewPathInWithDefaults instantiates a new PathIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PathIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PathIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PathIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PathIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *PathIn) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PathIn) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PathIn) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PathIn) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


