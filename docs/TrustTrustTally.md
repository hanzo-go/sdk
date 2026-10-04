# TrustTrustTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Absent** | Pointer to **int64** | Absent is how many the organization does not have. An absent control still names the clause it would satisfy — that is a roadmap — but it never moves a coverage number. | [optional] 
**Automated** | Pointer to **int64** | Automated is how many run with nobody in the loop. | [optional] 
**Partial** | Pointer to **int64** | Partial is how many run but do not cover their whole claim. Each says what is missing. | [optional] 
**Statement** | Pointer to **string** | Statement is the counts as one sentence, safe to quote. | [optional] 
**Total** | Pointer to **int64** | Total is how many controls this organization publishes. | [optional] 
**Unverified** | Pointer to **int64** | Unverified is how many rest on somebody having READ the source rather than on a test or an audit row. Only a check that can FAIL counts as verified, and coverage counts those one rung weaker than they claim to be. | [optional] 

## Methods

### NewTrustTrustTally

`func NewTrustTrustTally() *TrustTrustTally`

NewTrustTrustTally instantiates a new TrustTrustTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustTrustTallyWithDefaults

`func NewTrustTrustTallyWithDefaults() *TrustTrustTally`

NewTrustTrustTallyWithDefaults instantiates a new TrustTrustTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsent

`func (o *TrustTrustTally) GetAbsent() int64`

GetAbsent returns the Absent field if non-nil, zero value otherwise.

### GetAbsentOk

`func (o *TrustTrustTally) GetAbsentOk() (*int64, bool)`

GetAbsentOk returns a tuple with the Absent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsent

`func (o *TrustTrustTally) SetAbsent(v int64)`

SetAbsent sets Absent field to given value.

### HasAbsent

`func (o *TrustTrustTally) HasAbsent() bool`

HasAbsent returns a boolean if a field has been set.

### GetAutomated

`func (o *TrustTrustTally) GetAutomated() int64`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *TrustTrustTally) GetAutomatedOk() (*int64, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *TrustTrustTally) SetAutomated(v int64)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *TrustTrustTally) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetPartial

`func (o *TrustTrustTally) GetPartial() int64`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *TrustTrustTally) GetPartialOk() (*int64, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *TrustTrustTally) SetPartial(v int64)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *TrustTrustTally) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetStatement

`func (o *TrustTrustTally) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *TrustTrustTally) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *TrustTrustTally) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *TrustTrustTally) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetTotal

`func (o *TrustTrustTally) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TrustTrustTally) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TrustTrustTally) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TrustTrustTally) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnverified

`func (o *TrustTrustTally) GetUnverified() int64`

GetUnverified returns the Unverified field if non-nil, zero value otherwise.

### GetUnverifiedOk

`func (o *TrustTrustTally) GetUnverifiedOk() (*int64, bool)`

GetUnverifiedOk returns a tuple with the Unverified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverified

`func (o *TrustTrustTally) SetUnverified(v int64)`

SetUnverified sets Unverified field to given value.

### HasUnverified

`func (o *TrustTrustTally) HasUnverified() bool`

HasUnverified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


