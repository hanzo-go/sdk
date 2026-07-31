# CloudSbomIngested

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentCount** | Pointer to **int32** | ComponentCount is how many components the CycloneDX document yielded and this call persisted. | [optional] 
**ImageDigest** | Pointer to **string** | ImageDigest is the content-addressed digest the components were keyed under. | [optional] 

## Methods

### NewCloudSbomIngested

`func NewCloudSbomIngested() *CloudSbomIngested`

NewCloudSbomIngested instantiates a new CloudSbomIngested object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSbomIngestedWithDefaults

`func NewCloudSbomIngestedWithDefaults() *CloudSbomIngested`

NewCloudSbomIngestedWithDefaults instantiates a new CloudSbomIngested object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentCount

`func (o *CloudSbomIngested) GetComponentCount() int32`

GetComponentCount returns the ComponentCount field if non-nil, zero value otherwise.

### GetComponentCountOk

`func (o *CloudSbomIngested) GetComponentCountOk() (*int32, bool)`

GetComponentCountOk returns a tuple with the ComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCount

`func (o *CloudSbomIngested) SetComponentCount(v int32)`

SetComponentCount sets ComponentCount field to given value.

### HasComponentCount

`func (o *CloudSbomIngested) HasComponentCount() bool`

HasComponentCount returns a boolean if a field has been set.

### GetImageDigest

`func (o *CloudSbomIngested) GetImageDigest() string`

GetImageDigest returns the ImageDigest field if non-nil, zero value otherwise.

### GetImageDigestOk

`func (o *CloudSbomIngested) GetImageDigestOk() (*string, bool)`

GetImageDigestOk returns a tuple with the ImageDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDigest

`func (o *CloudSbomIngested) SetImageDigest(v string)`

SetImageDigest sets ImageDigest field to given value.

### HasImageDigest

`func (o *CloudSbomIngested) HasImageDigest() bool`

HasImageDigest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


