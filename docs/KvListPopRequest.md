# KvListPopRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **string** |  | [optional] [default to "left"]
**Count** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewKvListPopRequest

`func NewKvListPopRequest() *KvListPopRequest`

NewKvListPopRequest instantiates a new KvListPopRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvListPopRequestWithDefaults

`func NewKvListPopRequestWithDefaults() *KvListPopRequest`

NewKvListPopRequestWithDefaults instantiates a new KvListPopRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *KvListPopRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *KvListPopRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *KvListPopRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *KvListPopRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetCount

`func (o *KvListPopRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KvListPopRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KvListPopRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *KvListPopRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


