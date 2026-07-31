# CloudRegistrationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the registration&#39;s handle. | [optional] 
**NodeID** | Pointer to **string** | NodeID is the luxd node the registration is for. | [optional] 
**Status** | Pointer to **string** | Status is the registration&#39;s lifecycle state; \&quot;pending_owner_approval\&quot; until the owner co-signs it out of band. | [optional] 

## Methods

### NewCloudRegistrationView

`func NewCloudRegistrationView() *CloudRegistrationView`

NewCloudRegistrationView instantiates a new CloudRegistrationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegistrationViewWithDefaults

`func NewCloudRegistrationViewWithDefaults() *CloudRegistrationView`

NewCloudRegistrationViewWithDefaults instantiates a new CloudRegistrationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudRegistrationView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudRegistrationView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudRegistrationView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudRegistrationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNodeID

`func (o *CloudRegistrationView) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *CloudRegistrationView) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *CloudRegistrationView) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *CloudRegistrationView) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetStatus

`func (o *CloudRegistrationView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudRegistrationView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudRegistrationView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudRegistrationView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


