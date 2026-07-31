# ProductSearchIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Index UID. | 
**DocCount** | **int64** | Number of documents in the index. | 
**LastIndexedAt** | **string** | Last-indexed timestamp (RFC 3339); null if unknown. | 
**CreatedAt** | **string** | Created timestamp (RFC 3339); falls back to now() if the upstream omits it. | 

## Methods

### NewProductSearchIndex

`func NewProductSearchIndex(name string, docCount int64, lastIndexedAt string, createdAt string, ) *ProductSearchIndex`

NewProductSearchIndex instantiates a new ProductSearchIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSearchIndexWithDefaults

`func NewProductSearchIndexWithDefaults() *ProductSearchIndex`

NewProductSearchIndexWithDefaults instantiates a new ProductSearchIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProductSearchIndex) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductSearchIndex) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductSearchIndex) SetName(v string)`

SetName sets Name field to given value.


### GetDocCount

`func (o *ProductSearchIndex) GetDocCount() int64`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *ProductSearchIndex) GetDocCountOk() (*int64, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *ProductSearchIndex) SetDocCount(v int64)`

SetDocCount sets DocCount field to given value.


### GetLastIndexedAt

`func (o *ProductSearchIndex) GetLastIndexedAt() string`

GetLastIndexedAt returns the LastIndexedAt field if non-nil, zero value otherwise.

### GetLastIndexedAtOk

`func (o *ProductSearchIndex) GetLastIndexedAtOk() (*string, bool)`

GetLastIndexedAtOk returns a tuple with the LastIndexedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexedAt

`func (o *ProductSearchIndex) SetLastIndexedAt(v string)`

SetLastIndexedAt sets LastIndexedAt field to given value.


### GetCreatedAt

`func (o *ProductSearchIndex) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProductSearchIndex) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProductSearchIndex) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


