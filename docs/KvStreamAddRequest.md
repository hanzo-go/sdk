# KvStreamAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Entry ID (auto-generated if omitted) | [optional] 
**Fields** | **map[string]string** |  | 
**Maxlen** | Pointer to **int32** | Trim stream to this max length | [optional] 

## Methods

### NewKvStreamAddRequest

`func NewKvStreamAddRequest(fields map[string]string, ) *KvStreamAddRequest`

NewKvStreamAddRequest instantiates a new KvStreamAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvStreamAddRequestWithDefaults

`func NewKvStreamAddRequestWithDefaults() *KvStreamAddRequest`

NewKvStreamAddRequestWithDefaults instantiates a new KvStreamAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KvStreamAddRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KvStreamAddRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KvStreamAddRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KvStreamAddRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFields

`func (o *KvStreamAddRequest) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *KvStreamAddRequest) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *KvStreamAddRequest) SetFields(v map[string]string)`

SetFields sets Fields field to given value.


### GetMaxlen

`func (o *KvStreamAddRequest) GetMaxlen() int32`

GetMaxlen returns the Maxlen field if non-nil, zero value otherwise.

### GetMaxlenOk

`func (o *KvStreamAddRequest) GetMaxlenOk() (*int32, bool)`

GetMaxlenOk returns a tuple with the Maxlen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxlen

`func (o *KvStreamAddRequest) SetMaxlen(v int32)`

SetMaxlen sets Maxlen field to given value.

### HasMaxlen

`func (o *KvStreamAddRequest) HasMaxlen() bool`

HasMaxlen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


