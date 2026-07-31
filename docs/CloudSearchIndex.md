# CloudSearchIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is the index&#39;s creation time (RFC 3339); it falls back to now when the index list could not be read. | [optional] 
**DocCount** | Pointer to **int32** | DocCount is how many documents the index currently holds. | [optional] 
**LastIndexedAt** | Pointer to **string** | LastIndexedAt is the index&#39;s last update time (RFC 3339), null when the index list could not be read. | [optional] 
**Name** | Pointer to **string** | Name is the index uid. | [optional] 

## Methods

### NewCloudSearchIndex

`func NewCloudSearchIndex() *CloudSearchIndex`

NewCloudSearchIndex instantiates a new CloudSearchIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSearchIndexWithDefaults

`func NewCloudSearchIndexWithDefaults() *CloudSearchIndex`

NewCloudSearchIndexWithDefaults instantiates a new CloudSearchIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudSearchIndex) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSearchIndex) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSearchIndex) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSearchIndex) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDocCount

`func (o *CloudSearchIndex) GetDocCount() int32`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *CloudSearchIndex) GetDocCountOk() (*int32, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *CloudSearchIndex) SetDocCount(v int32)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *CloudSearchIndex) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetLastIndexedAt

`func (o *CloudSearchIndex) GetLastIndexedAt() string`

GetLastIndexedAt returns the LastIndexedAt field if non-nil, zero value otherwise.

### GetLastIndexedAtOk

`func (o *CloudSearchIndex) GetLastIndexedAtOk() (*string, bool)`

GetLastIndexedAtOk returns a tuple with the LastIndexedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexedAt

`func (o *CloudSearchIndex) SetLastIndexedAt(v string)`

SetLastIndexedAt sets LastIndexedAt field to given value.

### HasLastIndexedAt

`func (o *CloudSearchIndex) HasLastIndexedAt() bool`

HasLastIndexedAt returns a boolean if a field has been set.

### GetName

`func (o *CloudSearchIndex) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSearchIndex) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSearchIndex) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudSearchIndex) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


