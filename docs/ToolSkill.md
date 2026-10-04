# ToolSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admitted** | Pointer to **bool** | Admitted is whether an admin of the org, or a SuperAdmin, wrote this content. Only admitted content is carried into the org&#39;s agent runs. A push writes content nobody administering the org attested, and a skill written before writing one took an admin says no until an admin writes it again. | [optional] 
**Content** | Pointer to **string** | Content is the SKILL.md body, markdown. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the skill was last written, Unix seconds. | [optional] 
**Description** | Pointer to **string** | Description is the one-line summary discovery shows for the skill. | [optional] 
**Id** | Pointer to **string** | ID is the skill&#39;s id within the org. It is DERIVED from Name, so writing the same name again revises that skill rather than adding another. | [optional] 
**Name** | Pointer to **string** | Name is the skill&#39;s name: one lowercase path segment (a-z0-9, _ or -). | [optional] 
**Org** | Pointer to **string** | Org is the org that authored the skill — the validated caller&#39;s, never a value the body supplied. | [optional] 
**Source** | Pointer to **string** | Source is the repository the skill was read from, \&quot;&lt;project&gt;/&lt;name&gt;\&quot; or \&quot;&lt;name&gt;\&quot;; empty for a skill written through the API. A push replaces every skill of its source at once, so a skill leaves when its file does. | [optional] 

## Methods

### NewToolSkill

`func NewToolSkill() *ToolSkill`

NewToolSkill instantiates a new ToolSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolSkillWithDefaults

`func NewToolSkillWithDefaults() *ToolSkill`

NewToolSkillWithDefaults instantiates a new ToolSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmitted

`func (o *ToolSkill) GetAdmitted() bool`

GetAdmitted returns the Admitted field if non-nil, zero value otherwise.

### GetAdmittedOk

`func (o *ToolSkill) GetAdmittedOk() (*bool, bool)`

GetAdmittedOk returns a tuple with the Admitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmitted

`func (o *ToolSkill) SetAdmitted(v bool)`

SetAdmitted sets Admitted field to given value.

### HasAdmitted

`func (o *ToolSkill) HasAdmitted() bool`

HasAdmitted returns a boolean if a field has been set.

### GetContent

`func (o *ToolSkill) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ToolSkill) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ToolSkill) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ToolSkill) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ToolSkill) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ToolSkill) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ToolSkill) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ToolSkill) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ToolSkill) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolSkill) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolSkill) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ToolSkill) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ToolSkill) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolSkill) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolSkill) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ToolSkill) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ToolSkill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolSkill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolSkill) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ToolSkill) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *ToolSkill) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ToolSkill) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ToolSkill) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ToolSkill) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetSource

`func (o *ToolSkill) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ToolSkill) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ToolSkill) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ToolSkill) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


