# CloudMemoryPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudMemoryEntry**](CloudMemoryEntry.md) | Data is the matching memory entries, newest first. | [optional] 

## Methods

### NewCloudMemoryPage

`func NewCloudMemoryPage() *CloudMemoryPage`

NewCloudMemoryPage instantiates a new CloudMemoryPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMemoryPageWithDefaults

`func NewCloudMemoryPageWithDefaults() *CloudMemoryPage`

NewCloudMemoryPageWithDefaults instantiates a new CloudMemoryPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudMemoryPage) GetData() []CloudMemoryEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudMemoryPage) GetDataOk() (*[]CloudMemoryEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudMemoryPage) SetData(v []CloudMemoryEntry)`

SetData sets Data field to given value.

### HasData

`func (o *CloudMemoryPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


