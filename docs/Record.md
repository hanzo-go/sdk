# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | the record name the customer creates | [optional] 
**Type** | Pointer to **string** | TXT | CNAME | [optional] 
**Value** | Pointer to **string** | the record value | [optional] 

## Methods

### NewRecord

`func NewRecord() *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Record) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Record) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Record) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Record) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Record) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Record) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Record) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Record) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *Record) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Record) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Record) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Record) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


