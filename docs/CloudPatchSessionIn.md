# CloudPatchSessionIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the session to update, from the path. | [optional] 
**Project** | Pointer to **string** | Project tags the product this session built; Published is the author&#39;s decision to let anyone read the story (provenance.go). Both are pointers so \&quot;absent\&quot; and \&quot;cleared\&quot; are different requests. | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** | Target re-dispatches a session to a run-target (the #48 association). \&quot;\&quot; detaches. | [optional] 
**Terminal** | Pointer to **string** | Terminal publishes (or, with \&quot;\&quot;, withdraws) the URL this session&#39;s live terminal can be watched at. A pointer so \&quot;absent\&quot; and \&quot;withdrawn\&quot; are different requests: a session that stops sharing must be able to say so. | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudPatchSessionIn

`func NewCloudPatchSessionIn() *CloudPatchSessionIn`

NewCloudPatchSessionIn instantiates a new CloudPatchSessionIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPatchSessionInWithDefaults

`func NewCloudPatchSessionInWithDefaults() *CloudPatchSessionIn`

NewCloudPatchSessionInWithDefaults instantiates a new CloudPatchSessionIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudPatchSessionIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudPatchSessionIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudPatchSessionIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudPatchSessionIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *CloudPatchSessionIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudPatchSessionIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudPatchSessionIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudPatchSessionIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPublished

`func (o *CloudPatchSessionIn) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CloudPatchSessionIn) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CloudPatchSessionIn) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CloudPatchSessionIn) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetStatus

`func (o *CloudPatchSessionIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudPatchSessionIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudPatchSessionIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudPatchSessionIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *CloudPatchSessionIn) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudPatchSessionIn) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudPatchSessionIn) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudPatchSessionIn) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTerminal

`func (o *CloudPatchSessionIn) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *CloudPatchSessionIn) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *CloudPatchSessionIn) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *CloudPatchSessionIn) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTitle

`func (o *CloudPatchSessionIn) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudPatchSessionIn) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudPatchSessionIn) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudPatchSessionIn) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


