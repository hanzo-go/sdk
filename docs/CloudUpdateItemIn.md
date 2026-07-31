# CloudUpdateItemIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | Assignee replaces the reviewer this item is for, up to 512 characters. | [optional] 
**Id** | Pointer to **string** | ID is the annotation queue the item belongs to, from the path. | [optional] 
**ItemId** | Pointer to **string** | ItemID is the item to update, from the path. | [optional] 
**Status** | Pointer to **string** | Status is the item&#39;s new review state: PENDING or COMPLETED. Required. | [optional] 

## Methods

### NewCloudUpdateItemIn

`func NewCloudUpdateItemIn() *CloudUpdateItemIn`

NewCloudUpdateItemIn instantiates a new CloudUpdateItemIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUpdateItemInWithDefaults

`func NewCloudUpdateItemInWithDefaults() *CloudUpdateItemIn`

NewCloudUpdateItemInWithDefaults instantiates a new CloudUpdateItemIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *CloudUpdateItemIn) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *CloudUpdateItemIn) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *CloudUpdateItemIn) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *CloudUpdateItemIn) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetId

`func (o *CloudUpdateItemIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudUpdateItemIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudUpdateItemIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudUpdateItemIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *CloudUpdateItemIn) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *CloudUpdateItemIn) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *CloudUpdateItemIn) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *CloudUpdateItemIn) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetStatus

`func (o *CloudUpdateItemIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudUpdateItemIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudUpdateItemIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudUpdateItemIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


