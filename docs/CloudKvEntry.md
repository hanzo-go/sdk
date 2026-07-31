# CloudKvEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Created is when this revision was written, RFC3339. | [optional] 
**Key** | Pointer to **string** | Key is the entry&#39;s key. | [optional] 
**Operation** | Pointer to **string** | Operation is what wrote the revision: put, del or purge. | [optional] 
**Revision** | Pointer to **int32** | Revision is the entry&#39;s revision in the bucket. | [optional] 
**Value** | Pointer to **string** | Value is the value as UTF-8 text; empty for delete and purge markers. | [optional] 

## Methods

### NewCloudKvEntry

`func NewCloudKvEntry() *CloudKvEntry`

NewCloudKvEntry instantiates a new CloudKvEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudKvEntryWithDefaults

`func NewCloudKvEntryWithDefaults() *CloudKvEntry`

NewCloudKvEntryWithDefaults instantiates a new CloudKvEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CloudKvEntry) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudKvEntry) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudKvEntry) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudKvEntry) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetKey

`func (o *CloudKvEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CloudKvEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CloudKvEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CloudKvEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOperation

`func (o *CloudKvEntry) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *CloudKvEntry) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *CloudKvEntry) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *CloudKvEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetRevision

`func (o *CloudKvEntry) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CloudKvEntry) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CloudKvEntry) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CloudKvEntry) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetValue

`func (o *CloudKvEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CloudKvEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CloudKvEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CloudKvEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


