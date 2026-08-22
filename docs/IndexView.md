# IndexView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when this org first created the index, RFC 3339. | [optional] 
**PrimaryKey** | Pointer to **string** | PrimaryKey is the document field that identifies a row; an upsert is keyed on it. Empty until a document establishes one. | [optional] 
**Uid** | Pointer to **string** | UID is the index&#39;s name within the org. Two orgs may both hold &#x60;messages&#x60;. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is when the index or its settings last changed, RFC 3339. | [optional] 

## Methods

### NewIndexView

`func NewIndexView() *IndexView`

NewIndexView instantiates a new IndexView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexViewWithDefaults

`func NewIndexViewWithDefaults() *IndexView`

NewIndexViewWithDefaults instantiates a new IndexView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IndexView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IndexView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IndexView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IndexView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPrimaryKey

`func (o *IndexView) GetPrimaryKey() string`

GetPrimaryKey returns the PrimaryKey field if non-nil, zero value otherwise.

### GetPrimaryKeyOk

`func (o *IndexView) GetPrimaryKeyOk() (*string, bool)`

GetPrimaryKeyOk returns a tuple with the PrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKey

`func (o *IndexView) SetPrimaryKey(v string)`

SetPrimaryKey sets PrimaryKey field to given value.

### HasPrimaryKey

`func (o *IndexView) HasPrimaryKey() bool`

HasPrimaryKey returns a boolean if a field has been set.

### GetUid

`func (o *IndexView) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IndexView) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IndexView) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *IndexView) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IndexView) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IndexView) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IndexView) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IndexView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


