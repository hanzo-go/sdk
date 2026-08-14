# MemoryPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MemoryEntry**](MemoryEntry.md) | Data is the matching memory entries, newest first. | [optional] 

## Methods

### NewMemoryPage

`func NewMemoryPage() *MemoryPage`

NewMemoryPage instantiates a new MemoryPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemoryPageWithDefaults

`func NewMemoryPageWithDefaults() *MemoryPage`

NewMemoryPageWithDefaults instantiates a new MemoryPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MemoryPage) GetData() []MemoryEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MemoryPage) GetDataOk() (*[]MemoryEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MemoryPage) SetData(v []MemoryEntry)`

SetData sets Data field to given value.

### HasData

`func (o *MemoryPage) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


