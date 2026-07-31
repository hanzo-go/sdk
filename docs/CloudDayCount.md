# CloudDayCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count is that day&#39;s total. | [optional] 
**Date** | Pointer to **string** | Date is the day, YYYY-MM-DD. | [optional] 

## Methods

### NewCloudDayCount

`func NewCloudDayCount() *CloudDayCount`

NewCloudDayCount instantiates a new CloudDayCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDayCountWithDefaults

`func NewCloudDayCountWithDefaults() *CloudDayCount`

NewCloudDayCountWithDefaults instantiates a new CloudDayCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CloudDayCount) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloudDayCount) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloudDayCount) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CloudDayCount) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDate

`func (o *CloudDayCount) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CloudDayCount) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CloudDayCount) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CloudDayCount) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


