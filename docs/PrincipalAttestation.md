# PrincipalAttestation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to [**PrincipalAgentBrief**](PrincipalAgentBrief.md) | Agent is the agent. | [optional] 
**Expires** | Pointer to **int64** | Expires is when the statement stops being valid, unix seconds. | [optional] 
**Kid** | Pointer to **string** | Kid is the key that signed it. | [optional] 
**Org** | Pointer to **string** | Org is the principal it answers to — the org it was defined in. | [optional] 
**Parents** | Pointer to [**[]PrincipalAgentBrief**](PrincipalAgentBrief.md) | Parents are the agents that spawned it, nearest first. | [optional] 
**Statement** | Pointer to **string** | Statement is the same, signed: a JWS a verifier checks offline against GET /v1/principal/keys. | [optional] 

## Methods

### NewPrincipalAttestation

`func NewPrincipalAttestation() *PrincipalAttestation`

NewPrincipalAttestation instantiates a new PrincipalAttestation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalAttestationWithDefaults

`func NewPrincipalAttestationWithDefaults() *PrincipalAttestation`

NewPrincipalAttestationWithDefaults instantiates a new PrincipalAttestation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *PrincipalAttestation) GetAgent() PrincipalAgentBrief`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *PrincipalAttestation) GetAgentOk() (*PrincipalAgentBrief, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *PrincipalAttestation) SetAgent(v PrincipalAgentBrief)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *PrincipalAttestation) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetExpires

`func (o *PrincipalAttestation) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PrincipalAttestation) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PrincipalAttestation) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PrincipalAttestation) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetKid

`func (o *PrincipalAttestation) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *PrincipalAttestation) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *PrincipalAttestation) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *PrincipalAttestation) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetOrg

`func (o *PrincipalAttestation) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PrincipalAttestation) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PrincipalAttestation) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PrincipalAttestation) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetParents

`func (o *PrincipalAttestation) GetParents() []PrincipalAgentBrief`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *PrincipalAttestation) GetParentsOk() (*[]PrincipalAgentBrief, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *PrincipalAttestation) SetParents(v []PrincipalAgentBrief)`

SetParents sets Parents field to given value.

### HasParents

`func (o *PrincipalAttestation) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetStatement

`func (o *PrincipalAttestation) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *PrincipalAttestation) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *PrincipalAttestation) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *PrincipalAttestation) HasStatement() bool`

HasStatement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


