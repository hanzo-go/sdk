# Activity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldCount** | Pointer to **map[string]int32** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 

## Methods

### NewActivity

`func NewActivity() *Activity`

NewActivity instantiates a new Activity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityWithDefaults

`func NewActivityWithDefaults() *Activity`

NewActivityWithDefaults instantiates a new Activity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldCount

`func (o *Activity) GetFieldCount() map[string]int32`

GetFieldCount returns the FieldCount field if non-nil, zero value otherwise.

### GetFieldCountOk

`func (o *Activity) GetFieldCountOk() (*map[string]int32, bool)`

GetFieldCountOk returns a tuple with the FieldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCount

`func (o *Activity) SetFieldCount(v map[string]int32)`

SetFieldCount sets FieldCount field to given value.

### HasFieldCount

`func (o *Activity) HasFieldCount() bool`

HasFieldCount returns a boolean if a field has been set.

### GetDate

`func (o *Activity) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Activity) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Activity) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Activity) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


