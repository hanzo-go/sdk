# PlatformComposeDeleteRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComposeId** | **string** |  | 
**DeleteVolumes** | Pointer to **bool** |  | [optional] 

## Methods

### NewPlatformComposeDeleteRequestJson

`func NewPlatformComposeDeleteRequestJson(composeId string, ) *PlatformComposeDeleteRequestJson`

NewPlatformComposeDeleteRequestJson instantiates a new PlatformComposeDeleteRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformComposeDeleteRequestJsonWithDefaults

`func NewPlatformComposeDeleteRequestJsonWithDefaults() *PlatformComposeDeleteRequestJson`

NewPlatformComposeDeleteRequestJsonWithDefaults instantiates a new PlatformComposeDeleteRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComposeId

`func (o *PlatformComposeDeleteRequestJson) GetComposeId() string`

GetComposeId returns the ComposeId field if non-nil, zero value otherwise.

### GetComposeIdOk

`func (o *PlatformComposeDeleteRequestJson) GetComposeIdOk() (*string, bool)`

GetComposeIdOk returns a tuple with the ComposeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeId

`func (o *PlatformComposeDeleteRequestJson) SetComposeId(v string)`

SetComposeId sets ComposeId field to given value.


### GetDeleteVolumes

`func (o *PlatformComposeDeleteRequestJson) GetDeleteVolumes() bool`

GetDeleteVolumes returns the DeleteVolumes field if non-nil, zero value otherwise.

### GetDeleteVolumesOk

`func (o *PlatformComposeDeleteRequestJson) GetDeleteVolumesOk() (*bool, bool)`

GetDeleteVolumesOk returns a tuple with the DeleteVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVolumes

`func (o *PlatformComposeDeleteRequestJson) SetDeleteVolumes(v bool)`

SetDeleteVolumes sets DeleteVolumes field to given value.

### HasDeleteVolumes

`func (o *PlatformComposeDeleteRequestJson) HasDeleteVolumes() bool`

HasDeleteVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


