# AuthoredSkillList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skills** | Pointer to [**[]Skill**](Skill.md) | Skills is every skill this org authored, each with its SKILL.md content. | [optional] 

## Methods

### NewAuthoredSkillList

`func NewAuthoredSkillList() *AuthoredSkillList`

NewAuthoredSkillList instantiates a new AuthoredSkillList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthoredSkillListWithDefaults

`func NewAuthoredSkillListWithDefaults() *AuthoredSkillList`

NewAuthoredSkillListWithDefaults instantiates a new AuthoredSkillList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkills

`func (o *AuthoredSkillList) GetSkills() []Skill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *AuthoredSkillList) GetSkillsOk() (*[]Skill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *AuthoredSkillList) SetSkills(v []Skill)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *AuthoredSkillList) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


