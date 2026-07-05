# KvStreamEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Stream entry ID (timestamp-sequence) | [optional] 
**Fields** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewKvStreamEntry

`func NewKvStreamEntry() *KvStreamEntry`

NewKvStreamEntry instantiates a new KvStreamEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvStreamEntryWithDefaults

`func NewKvStreamEntryWithDefaults() *KvStreamEntry`

NewKvStreamEntryWithDefaults instantiates a new KvStreamEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KvStreamEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KvStreamEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KvStreamEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KvStreamEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFields

`func (o *KvStreamEntry) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *KvStreamEntry) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *KvStreamEntry) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *KvStreamEntry) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


