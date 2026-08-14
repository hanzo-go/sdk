# CaptableSafes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CaptableSafe**](CaptableSafe.md) | Data is every SAFE, newest first. | [optional] 

## Methods

### NewCaptableSafes

`func NewCaptableSafes() *CaptableSafes`

NewCaptableSafes instantiates a new CaptableSafes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableSafesWithDefaults

`func NewCaptableSafesWithDefaults() *CaptableSafes`

NewCaptableSafesWithDefaults instantiates a new CaptableSafes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CaptableSafes) GetData() []CaptableSafe`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CaptableSafes) GetDataOk() (*[]CaptableSafe, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CaptableSafes) SetData(v []CaptableSafe)`

SetData sets Data field to given value.

### HasData

`func (o *CaptableSafes) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


