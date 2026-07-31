# CloudBotMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Active is whether the agent projects as a LIVE workspace member, derived from its registry status: empty, \&quot;active\&quot; and \&quot;ready\&quot; are live, anything else (archived/retired) is not. An inactive bot drops out of the Team list while its past authorship survives. | [optional] 
**Id** | Pointer to **string** | the agent id | [optional] 
**Name** | Pointer to **string** | display name | [optional] 
**PersonRef** | Pointer to **string** | the projected Person _id | [optional] 
**UserId** | Pointer to **string** | derived member account uuid (personUuid) | [optional] 

## Methods

### NewCloudBotMember

`func NewCloudBotMember() *CloudBotMember`

NewCloudBotMember instantiates a new CloudBotMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBotMemberWithDefaults

`func NewCloudBotMemberWithDefaults() *CloudBotMember`

NewCloudBotMemberWithDefaults instantiates a new CloudBotMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CloudBotMember) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CloudBotMember) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CloudBotMember) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CloudBotMember) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetId

`func (o *CloudBotMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudBotMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudBotMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudBotMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudBotMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudBotMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudBotMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudBotMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersonRef

`func (o *CloudBotMember) GetPersonRef() string`

GetPersonRef returns the PersonRef field if non-nil, zero value otherwise.

### GetPersonRefOk

`func (o *CloudBotMember) GetPersonRefOk() (*string, bool)`

GetPersonRefOk returns a tuple with the PersonRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonRef

`func (o *CloudBotMember) SetPersonRef(v string)`

SetPersonRef sets PersonRef field to given value.

### HasPersonRef

`func (o *CloudBotMember) HasPersonRef() bool`

HasPersonRef returns a boolean if a field has been set.

### GetUserId

`func (o *CloudBotMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CloudBotMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CloudBotMember) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CloudBotMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


