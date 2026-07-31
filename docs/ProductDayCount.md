# ProductDayCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Day (date string). | 
**Count** | **int64** | Count for that day. | 

## Methods

### NewProductDayCount

`func NewProductDayCount(date string, count int64, ) *ProductDayCount`

NewProductDayCount instantiates a new ProductDayCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDayCountWithDefaults

`func NewProductDayCountWithDefaults() *ProductDayCount`

NewProductDayCountWithDefaults instantiates a new ProductDayCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ProductDayCount) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ProductDayCount) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ProductDayCount) SetDate(v string)`

SetDate sets Date field to given value.


### GetCount

`func (o *ProductDayCount) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ProductDayCount) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ProductDayCount) SetCount(v int64)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


