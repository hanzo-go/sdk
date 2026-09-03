# SbomIngested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentCount** | Pointer to **int64** | ComponentCount is how many components the CycloneDX document yielded and this call persisted. | [optional] 
**ImageDigest** | Pointer to **string** | ImageDigest is the content-addressed digest the components were keyed under. | [optional] 

## Methods

### NewSbomIngested

`func NewSbomIngested() *SbomIngested`

NewSbomIngested instantiates a new SbomIngested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSbomIngestedWithDefaults

`func NewSbomIngestedWithDefaults() *SbomIngested`

NewSbomIngestedWithDefaults instantiates a new SbomIngested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentCount

`func (o *SbomIngested) GetComponentCount() int64`

GetComponentCount returns the ComponentCount field if non-nil, zero value otherwise.

### GetComponentCountOk

`func (o *SbomIngested) GetComponentCountOk() (*int64, bool)`

GetComponentCountOk returns a tuple with the ComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCount

`func (o *SbomIngested) SetComponentCount(v int64)`

SetComponentCount sets ComponentCount field to given value.

### HasComponentCount

`func (o *SbomIngested) HasComponentCount() bool`

HasComponentCount returns a boolean if a field has been set.

### GetImageDigest

`func (o *SbomIngested) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *SbomIngested) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *SbomIngested) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *SbomIngested) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


