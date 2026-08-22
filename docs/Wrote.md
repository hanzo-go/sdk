# Wrote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int32** | Bytes is how many bytes the file now holds. A write REPLACES the file, so this is its whole length and not an amount appended, and 0 is a legitimate answer: a WriteIn with no Data truncates the file to nothing. | [optional] 
**Path** | Pointer to **string** | Path is where the bytes actually landed: the caller&#39;s path resolved against the sandbox&#39;s working directory (Leased.Workdir), which is what a later read or a shell line inside the sandbox has to name. | [optional] 

## Methods

### NewWrote

`func NewWrote() *Wrote`

NewWrote instantiates a new Wrote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWroteWithDefaults

`func NewWroteWithDefaults() *Wrote`

NewWroteWithDefaults instantiates a new Wrote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *Wrote) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *Wrote) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *Wrote) SetBytes(v int32)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *Wrote) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetPath

`func (o *Wrote) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Wrote) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Wrote) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Wrote) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


