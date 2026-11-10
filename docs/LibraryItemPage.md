# LibraryItemPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]LibraryLibraryItem**](LibraryLibraryItem.md) | Data are the items, newest first. | [optional] 
**Next** | Pointer to **string** | Next continues the list as &#x60;cursor&#x60;; empty at its end. | [optional] 

## Methods

### NewLibraryItemPage

`func NewLibraryItemPage() *LibraryItemPage`

NewLibraryItemPage instantiates a new LibraryItemPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryItemPageWithDefaults

`func NewLibraryItemPageWithDefaults() *LibraryItemPage`

NewLibraryItemPageWithDefaults instantiates a new LibraryItemPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LibraryItemPage) GetData() []LibraryLibraryItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LibraryItemPage) GetDataOk() (*[]LibraryLibraryItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LibraryItemPage) SetData(v []LibraryLibraryItem)`

SetData sets Data field to given value.

### HasData

`func (o *LibraryItemPage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNext

`func (o *LibraryItemPage) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *LibraryItemPage) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *LibraryItemPage) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *LibraryItemPage) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


