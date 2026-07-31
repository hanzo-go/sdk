# CloudCatalogList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudCatalogEntry**](CloudCatalogEntry.md) | Data is every starter prompt, each importable as-is with POST /v1/prompts. | [optional] 

## Methods

### NewCloudCatalogList

`func NewCloudCatalogList() *CloudCatalogList`

NewCloudCatalogList instantiates a new CloudCatalogList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCatalogListWithDefaults

`func NewCloudCatalogListWithDefaults() *CloudCatalogList`

NewCloudCatalogListWithDefaults instantiates a new CloudCatalogList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudCatalogList) GetData() []CloudCatalogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudCatalogList) GetDataOk() (*[]CloudCatalogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudCatalogList) SetData(v []CloudCatalogEntry)`

SetData sets Data field to given value.

### HasData

`func (o *CloudCatalogList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


