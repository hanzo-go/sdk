# BoardItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctype** | Pointer to **string** | DocType is which content type the row came from: Campaign, SocialPost or Asset. The board spans all three at once, so this is what tells them apart. | [optional] 
**Name** | Pointer to **string** | Name is the document within that type. (doctype, name) is the pair every /v1/content write addresses an item by. | [optional] 
**Project** | Pointer to **string** | Project is the brand/site sub-scope within the org. Absent for an item held at org level rather than under one brand. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle state: draft, in_review, approved, queued, published or archived. It decides what a reader may see — the public site pulls exactly \&quot;published\&quot; and nothing else — so it is a visibility fact, not a workflow label. | [optional] 
**Title** | Pointer to **string** | Title is the item&#39;s headline, read from its type&#39;s own title field. Empty for a document that has none. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is unix seconds of the document&#39;s last write, and the key the board sorts on, newest first. | [optional] 

## Methods

### NewBoardItem

`func NewBoardItem() *BoardItem`

NewBoardItem instantiates a new BoardItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardItemWithDefaults

`func NewBoardItemWithDefaults() *BoardItem`

NewBoardItemWithDefaults instantiates a new BoardItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctype

`func (o *BoardItem) GetDoctype() string`

GetDoctype returns the Doctype field if non-nil, zero value otherwise.

### GetDoctypeOk

`func (o *BoardItem) GetDoctypeOk() (*string, bool)`

GetDoctypeOk returns a tuple with the Doctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctype

`func (o *BoardItem) SetDoctype(v string)`

SetDoctype sets Doctype field to given value.

### HasDoctype

`func (o *BoardItem) HasDoctype() bool`

HasDoctype returns a boolean if a field has been set.

### GetName

`func (o *BoardItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BoardItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *BoardItem) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BoardItem) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BoardItem) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BoardItem) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetStatus

`func (o *BoardItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BoardItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BoardItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BoardItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *BoardItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BoardItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BoardItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BoardItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BoardItem) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BoardItem) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BoardItem) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BoardItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


