# KvListPushRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | **[]string** |  | 
**Direction** | Pointer to **string** |  | [optional] [default to "right"]

## Methods

### NewKvListPushRequest

`func NewKvListPushRequest(values []string, ) *KvListPushRequest`

NewKvListPushRequest instantiates a new KvListPushRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvListPushRequestWithDefaults

`func NewKvListPushRequestWithDefaults() *KvListPushRequest`

NewKvListPushRequestWithDefaults instantiates a new KvListPushRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *KvListPushRequest) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *KvListPushRequest) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *KvListPushRequest) SetValues(v []string)`

SetValues sets Values field to given value.


### GetDirection

`func (o *KvListPushRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *KvListPushRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *KvListPushRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *KvListPushRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


