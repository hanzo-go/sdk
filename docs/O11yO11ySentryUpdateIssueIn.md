# O11yO11ySentryUpdateIssueIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **string** | Assignee is who the issue is assigned to. | [optional] 
**Id** | **string** | ID is the issue id. | 
**Status** | Pointer to **string** | Status is the new lifecycle state: unresolved, resolved or ignored. | [optional] 

## Methods

### NewO11yO11ySentryUpdateIssueIn

`func NewO11yO11ySentryUpdateIssueIn(id string, ) *O11yO11ySentryUpdateIssueIn`

NewO11yO11ySentryUpdateIssueIn instantiates a new O11yO11ySentryUpdateIssueIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySentryUpdateIssueInWithDefaults

`func NewO11yO11ySentryUpdateIssueInWithDefaults() *O11yO11ySentryUpdateIssueIn`

NewO11yO11ySentryUpdateIssueInWithDefaults instantiates a new O11yO11ySentryUpdateIssueIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *O11yO11ySentryUpdateIssueIn) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *O11yO11ySentryUpdateIssueIn) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *O11yO11ySentryUpdateIssueIn) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *O11yO11ySentryUpdateIssueIn) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetId

`func (o *O11yO11ySentryUpdateIssueIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11ySentryUpdateIssueIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11ySentryUpdateIssueIn) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *O11yO11ySentryUpdateIssueIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11ySentryUpdateIssueIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11ySentryUpdateIssueIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11ySentryUpdateIssueIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


