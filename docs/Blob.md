# Blob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** |  | [optional] 
**Dir** | Pointer to **bool** |  | [optional] 
**Entries** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewBlob

`func NewBlob() *Blob`

NewBlob instantiates a new Blob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobWithDefaults

`func NewBlobWithDefaults() *Blob`

NewBlobWithDefaults instantiates a new Blob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Blob) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Blob) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Blob) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *Blob) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDir

`func (o *Blob) GetDir() bool`

GetDir returns the Dir field if non-nil, zero value otherwise.

### GetDirOk

`func (o *Blob) GetDirOk() (*bool, bool)`

GetDirOk returns a tuple with the Dir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDir

`func (o *Blob) SetDir(v bool)`

SetDir sets Dir field to given value.

### HasDir

`func (o *Blob) HasDir() bool`

HasDir returns a boolean if a field has been set.

### GetEntries

`func (o *Blob) GetEntries() []string`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *Blob) GetEntriesOk() (*[]string, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *Blob) SetEntries(v []string)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *Blob) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetPath

`func (o *Blob) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Blob) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Blob) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Blob) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


