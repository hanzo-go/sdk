# DocumentList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]map[string]map[string]interface{}** | Data is the matching documents, newest-updated first unless order_by said otherwise, each projected to the requested fields plus the envelope keys. | [optional] 

## Methods

### NewDocumentList

`func NewDocumentList() *DocumentList`

NewDocumentList instantiates a new DocumentList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentListWithDefaults

`func NewDocumentListWithDefaults() *DocumentList`

NewDocumentListWithDefaults instantiates a new DocumentList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DocumentList) GetData() []map[string]map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DocumentList) GetDataOk() (*[]map[string]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DocumentList) SetData(v []map[string]map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DocumentList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


