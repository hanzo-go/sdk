# CloudAuthoredSkillList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skills** | Pointer to [**[]CloudSkill**](CloudSkill.md) | Skills is every skill this org authored, each with its SKILL.md content. | [optional] 

## Methods

### NewCloudAuthoredSkillList

`func NewCloudAuthoredSkillList() *CloudAuthoredSkillList`

NewCloudAuthoredSkillList instantiates a new CloudAuthoredSkillList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAuthoredSkillListWithDefaults

`func NewCloudAuthoredSkillListWithDefaults() *CloudAuthoredSkillList`

NewCloudAuthoredSkillListWithDefaults instantiates a new CloudAuthoredSkillList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkills

`func (o *CloudAuthoredSkillList) GetSkills() []CloudSkill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *CloudAuthoredSkillList) GetSkillsOk() (*[]CloudSkill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *CloudAuthoredSkillList) SetSkills(v []CloudSkill)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *CloudAuthoredSkillList) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


