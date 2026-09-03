# Skill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the SKILL.md body, markdown. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the skill was last written, Unix seconds. | [optional] 
**Description** | Pointer to **string** | Description is the one-line summary discovery shows for the skill. | [optional] 
**Id** | Pointer to **string** | ID is the skill&#39;s id within the org. It is DERIVED from Name, so writing the same name again revises that skill rather than adding another. | [optional] 
**Name** | Pointer to **string** | Name is the skill&#39;s name: one lowercase path segment (a-z0-9, _ or -). | [optional] 
**Org** | Pointer to **string** | Org is the org that authored the skill — the validated caller&#39;s, never a value the body supplied. | [optional] 
**Source** | Pointer to **string** | Source is the repository the skill was read from, \&quot;&lt;project&gt;/&lt;name&gt;\&quot; or \&quot;&lt;name&gt;\&quot;; empty for a skill written through the API. A push replaces every skill of its source at once, so a skill leaves when its file does. | [optional] 

## Methods

### NewSkill

`func NewSkill() *Skill`

NewSkill instantiates a new Skill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillWithDefaults

`func NewSkillWithDefaults() *Skill`

NewSkillWithDefaults instantiates a new Skill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Skill) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Skill) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Skill) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Skill) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Skill) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Skill) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Skill) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Skill) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Skill) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Skill) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Skill) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Skill) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Skill) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Skill) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Skill) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Skill) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Skill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Skill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Skill) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Skill) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *Skill) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Skill) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Skill) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Skill) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSource

`func (o *Skill) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Skill) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Skill) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Skill) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


