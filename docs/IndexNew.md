# IndexNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryKey** | Pointer to **string** | PrimaryKey is the document field that identifies a row. Optional — the first write establishes one when it is omitted. | [optional] 
**Uid** | Pointer to **string** | UID is the index&#39;s name within the org. Required. | [optional] 

## Methods

### NewIndexNew

`func NewIndexNew() *IndexNew`

NewIndexNew instantiates a new IndexNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexNewWithDefaults

`func NewIndexNewWithDefaults() *IndexNew`

NewIndexNewWithDefaults instantiates a new IndexNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryKey

`func (o *IndexNew) GetPrimaryKey() string`

GetPrimaryKey returns the PrimaryKey field if non-nil, zero value otherwise.

### GetPrimaryKeyOk

`func (o *IndexNew) GetPrimaryKeyOk() (*string, bool)`

GetPrimaryKeyOk returns a tuple with the PrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKey

`func (o *IndexNew) SetPrimaryKey(v string)`

SetPrimaryKey sets PrimaryKey field to given value.

### HasPrimaryKey

`func (o *IndexNew) HasPrimaryKey() bool`

HasPrimaryKey returns a boolean if a field has been set.

### GetUid

`func (o *IndexNew) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IndexNew) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IndexNew) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IndexNew) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


