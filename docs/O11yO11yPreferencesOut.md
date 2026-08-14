# O11yO11yPreferencesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]O11yO11yPreference**](O11yO11yPreference.md) | Data holds the preferences. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;success\&quot;. | [optional] 

## Methods

### NewO11yO11yPreferencesOut

`func NewO11yO11yPreferencesOut() *O11yO11yPreferencesOut`

NewO11yO11yPreferencesOut instantiates a new O11yO11yPreferencesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yPreferencesOutWithDefaults

`func NewO11yO11yPreferencesOutWithDefaults() *O11yO11yPreferencesOut`

NewO11yO11yPreferencesOutWithDefaults instantiates a new O11yO11yPreferencesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yO11yPreferencesOut) GetData() []O11yO11yPreference`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yO11yPreferencesOut) GetDataOk() (*[]O11yO11yPreference, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yO11yPreferencesOut) SetData(v []O11yO11yPreference)`

SetData sets Data field to given value.

### HasData

`func (o *O11yO11yPreferencesOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yPreferencesOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yPreferencesOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yPreferencesOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yPreferencesOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


