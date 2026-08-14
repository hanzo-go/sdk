# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**External** | Pointer to **bool** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Range** | Pointer to [**Range**](Range.md) |  | [optional] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternal

`func (o *Location) GetExternal() bool`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *Location) GetExternalOk() (*bool, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *Location) SetExternal(v bool)`

SetExternal sets External field to given value.

### HasExternal

`func (o *Location) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetPath

`func (o *Location) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Location) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Location) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Location) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRange

`func (o *Location) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Location) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Location) SetRange(v Range)`

SetRange sets Range field to given value.

### HasRange

`func (o *Location) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


