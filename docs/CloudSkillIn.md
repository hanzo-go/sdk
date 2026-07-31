# CloudSkillIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the SKILL.md body. Required, at most 256 KiB. | [optional] 
**Description** | Pointer to **string** | Description is the one-line summary discovery shows for the skill. | [optional] 
**Name** | Pointer to **string** | Name is the skill&#39;s id within the org: one lowercase path segment (a-z0-9, _ or -). Writing an existing name REVISES that skill. | [optional] 

## Methods

### NewCloudSkillIn

`func NewCloudSkillIn() *CloudSkillIn`

NewCloudSkillIn instantiates a new CloudSkillIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSkillInWithDefaults

`func NewCloudSkillInWithDefaults() *CloudSkillIn`

NewCloudSkillInWithDefaults instantiates a new CloudSkillIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CloudSkillIn) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CloudSkillIn) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CloudSkillIn) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CloudSkillIn) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDescription

`func (o *CloudSkillIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudSkillIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudSkillIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudSkillIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CloudSkillIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSkillIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSkillIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudSkillIn) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


