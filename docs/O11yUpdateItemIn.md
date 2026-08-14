# O11yUpdateItemIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | Assignee replaces the reviewer this item is for, up to 512 characters. | [optional] 
**Id** | Pointer to **string** | ID is the annotation queue the item belongs to, from the path. | [optional] 
**ItemId** | Pointer to **string** | ItemID is the item to update, from the path. | [optional] 
**Status** | Pointer to **string** | Status is the item&#39;s new review state: PENDING or COMPLETED. Required. | [optional] 

## Methods

### NewO11yUpdateItemIn

`func NewO11yUpdateItemIn() *O11yUpdateItemIn`

NewO11yUpdateItemIn instantiates a new O11yUpdateItemIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yUpdateItemInWithDefaults

`func NewO11yUpdateItemInWithDefaults() *O11yUpdateItemIn`

NewO11yUpdateItemInWithDefaults instantiates a new O11yUpdateItemIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *O11yUpdateItemIn) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *O11yUpdateItemIn) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *O11yUpdateItemIn) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *O11yUpdateItemIn) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetId

`func (o *O11yUpdateItemIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yUpdateItemIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yUpdateItemIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yUpdateItemIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *O11yUpdateItemIn) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *O11yUpdateItemIn) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *O11yUpdateItemIn) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *O11yUpdateItemIn) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetStatus

`func (o *O11yUpdateItemIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yUpdateItemIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yUpdateItemIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yUpdateItemIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


