# IndexIndexNew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryKey** | Pointer to **string** | PrimaryKey is the document field that identifies a row. Optional — the first write establishes one when it is omitted. | [optional] 
**Uid** | Pointer to **string** | UID is the index&#39;s name within the org. Required. | [optional] 

## Methods

### NewIndexIndexNew

`func NewIndexIndexNew() *IndexIndexNew`

NewIndexIndexNew instantiates a new IndexIndexNew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexIndexNewWithDefaults

`func NewIndexIndexNewWithDefaults() *IndexIndexNew`

NewIndexIndexNewWithDefaults instantiates a new IndexIndexNew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryKey

`func (o *IndexIndexNew) GetPrimaryKey() string`

GetPrimaryKey returns the PrimaryKey field if non-nil, zero value otherwise.

### GetPrimaryKeyOk

`func (o *IndexIndexNew) GetPrimaryKeyOk() (*string, bool)`

GetPrimaryKeyOk returns a tuple with the PrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKey

`func (o *IndexIndexNew) SetPrimaryKey(v string)`

SetPrimaryKey sets PrimaryKey field to given value.

### HasPrimaryKey

`func (o *IndexIndexNew) HasPrimaryKey() bool`

HasPrimaryKey returns a boolean if a field has been set.

### GetUid

`func (o *IndexIndexNew) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IndexIndexNew) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IndexIndexNew) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IndexIndexNew) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


