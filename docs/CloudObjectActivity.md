# CloudObjectActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldCount** | Pointer to  |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudObjectActivity

`func NewCloudObjectActivity() *CloudObjectActivity`

NewCloudObjectActivity instantiates a new CloudObjectActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectActivityWithDefaults

`func NewCloudObjectActivityWithDefaults() *CloudObjectActivity`

NewCloudObjectActivityWithDefaults instantiates a new CloudObjectActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldCount

`func (o *CloudObjectActivity) GetFieldCount() map[string]int64`

GetFieldCount returns the FieldCount field if non-nil, zero value otherwise.

### GetFieldCountOk

`func (o *CloudObjectActivity) GetFieldCountOk() (*map[string]int64, bool)`

GetFieldCountOk returns a tuple with the FieldCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCount

`func (o *CloudObjectActivity) SetFieldCount(v map[string]int64)`

SetFieldCount sets FieldCount field to given value.

### HasFieldCount

`func (o *CloudObjectActivity) HasFieldCount() bool`

HasFieldCount returns a boolean if a field has been set.

### SetFieldCountNil

`func (o *CloudObjectActivity) SetFieldCountNil(b bool)`

 SetFieldCountNil sets the value for FieldCount to be an explicit nil

### UnsetFieldCount
`func (o *CloudObjectActivity) UnsetFieldCount()`

UnsetFieldCount ensures that no value is present for FieldCount, not even an explicit nil
### GetDate

`func (o *CloudObjectActivity) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CloudObjectActivity) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CloudObjectActivity) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CloudObjectActivity) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


