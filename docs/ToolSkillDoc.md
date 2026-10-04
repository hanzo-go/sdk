# ToolSkillDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the skill&#39;s SKILL.md, verbatim. | [optional] 
**Name** | Pointer to **string** | Name is the skill&#39;s bare name, which is also the directory its SKILL.md is written into. | [optional] 

## Methods

### NewToolSkillDoc

`func NewToolSkillDoc() *ToolSkillDoc`

NewToolSkillDoc instantiates a new ToolSkillDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolSkillDocWithDefaults

`func NewToolSkillDocWithDefaults() *ToolSkillDoc`

NewToolSkillDocWithDefaults instantiates a new ToolSkillDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ToolSkillDoc) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ToolSkillDoc) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ToolSkillDoc) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ToolSkillDoc) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetName

`func (o *ToolSkillDoc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolSkillDoc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolSkillDoc) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ToolSkillDoc) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


