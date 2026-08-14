# DataroomAddDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** | ID is the room to add to. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | [optional] 
**OrderIndex** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewDataroomAddDocument

`func NewDataroomAddDocument() *DataroomAddDocument`

NewDataroomAddDocument instantiates a new DataroomAddDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomAddDocumentWithDefaults

`func NewDataroomAddDocumentWithDefaults() *DataroomAddDocument`

NewDataroomAddDocumentWithDefaults instantiates a new DataroomAddDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *DataroomAddDocument) GetDocumentId() interface{}`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *DataroomAddDocument) GetDocumentIdOk() (*interface{}, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *DataroomAddDocument) SetDocumentId(v interface{})`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *DataroomAddDocument) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### SetDocumentIdNil

`func (o *DataroomAddDocument) SetDocumentIdNil(b bool)`

 SetDocumentIdNil sets the value for DocumentId to be an explicit nil

### UnsetDocumentId
`func (o *DataroomAddDocument) UnsetDocumentId()`

UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil
### GetId

`func (o *DataroomAddDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataroomAddDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataroomAddDocument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataroomAddDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderIndex

`func (o *DataroomAddDocument) GetOrderIndex() interface{}`

GetOrderIndex returns the OrderIndex field if non-nil, zero value otherwise.

### GetOrderIndexOk

`func (o *DataroomAddDocument) GetOrderIndexOk() (*interface{}, bool)`

GetOrderIndexOk returns a tuple with the OrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderIndex

`func (o *DataroomAddDocument) SetOrderIndex(v interface{})`

SetOrderIndex sets OrderIndex field to given value.

### HasOrderIndex

`func (o *DataroomAddDocument) HasOrderIndex() bool`

HasOrderIndex returns a boolean if a field has been set.

### SetOrderIndexNil

`func (o *DataroomAddDocument) SetOrderIndexNil(b bool)`

 SetOrderIndexNil sets the value for OrderIndex to be an explicit nil

### UnsetOrderIndex
`func (o *DataroomAddDocument) UnsetOrderIndex()`

UnsetOrderIndex ensures that no value is present for OrderIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


