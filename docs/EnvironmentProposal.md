# EnvironmentProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Install** | Pointer to **string** | Install is the install script the agent checked. | [optional] 
**Note** | Pointer to **string** | Note is what the agent found and checked, in its own words. | [optional] 
**Secrets** | Pointer to **[]string** | Secrets are the environment variables the agent found the codebase reads. Some may not be set yet; setting them is a person&#39;s act, not the agent&#39;s. | [optional] 
**Start** | Pointer to **string** | Start is the start command the agent checked, or empty. | [optional] 

## Methods

### NewEnvironmentProposal

`func NewEnvironmentProposal() *EnvironmentProposal`

NewEnvironmentProposal instantiates a new EnvironmentProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentProposalWithDefaults

`func NewEnvironmentProposalWithDefaults() *EnvironmentProposal`

NewEnvironmentProposalWithDefaults instantiates a new EnvironmentProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstall

`func (o *EnvironmentProposal) GetInstall() string`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *EnvironmentProposal) GetInstallOk() (*string, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *EnvironmentProposal) SetInstall(v string)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *EnvironmentProposal) HasInstall() bool`

HasInstall returns a boolean if a field has been set.

### GetNote

`func (o *EnvironmentProposal) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *EnvironmentProposal) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *EnvironmentProposal) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *EnvironmentProposal) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetSecrets

`func (o *EnvironmentProposal) GetSecrets() []string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *EnvironmentProposal) GetSecretsOk() (*[]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *EnvironmentProposal) SetSecrets(v []string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *EnvironmentProposal) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetStart

`func (o *EnvironmentProposal) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EnvironmentProposal) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EnvironmentProposal) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *EnvironmentProposal) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


