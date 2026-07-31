# CloudSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the SKILL.md body, markdown. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the skill was last written, Unix seconds. | [optional] 
**Description** | Pointer to **string** | Description is the one-line summary discovery shows for the skill. | [optional] 
**Id** | Pointer to **string** | ID is the skill&#39;s id within the org. It is DERIVED from Name, so writing the same name again revises that skill rather than adding another. | [optional] 
**Name** | Pointer to **string** | Name is the skill&#39;s name: one lowercase path segment (a-z0-9, _ or -). | [optional] 
**Org** | Pointer to **string** | Org is the org that authored the skill — the validated caller&#39;s, never a value the body supplied. | [optional] 

## Methods

### NewCloudSkill

`func NewCloudSkill() *CloudSkill`

NewCloudSkill instantiates a new CloudSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSkillWithDefaults

`func NewCloudSkillWithDefaults() *CloudSkill`

NewCloudSkillWithDefaults instantiates a new CloudSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CloudSkill) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CloudSkill) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CloudSkill) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CloudSkill) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudSkill) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudSkill) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudSkill) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudSkill) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *CloudSkill) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudSkill) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudSkill) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudSkill) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CloudSkill) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudSkill) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudSkill) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudSkill) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudSkill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSkill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSkill) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudSkill) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *CloudSkill) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *CloudSkill) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *CloudSkill) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *CloudSkill) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


