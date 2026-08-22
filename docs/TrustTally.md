# TrustTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Absent** | Pointer to **int32** | Absent is how many the organization does not have. An absent control still names the clause it would satisfy — that is a roadmap — but it never moves a coverage number. | [optional] 
**Automated** | Pointer to **int32** | Automated is how many run with nobody in the loop. | [optional] 
**Partial** | Pointer to **int32** | Partial is how many run but do not cover their whole claim. Each says what is missing. | [optional] 
**Statement** | Pointer to **string** | Statement is the counts as one sentence, safe to quote. | [optional] 
**Total** | Pointer to **int32** | Total is how many controls this organization publishes. | [optional] 
**Unverified** | Pointer to **int32** | Unverified is how many rest on somebody having READ the source rather than on a test or an audit row. Only a check that can FAIL counts as verified, and coverage counts those one rung weaker than they claim to be. | [optional] 

## Methods

### NewTrustTally

`func NewTrustTally() *TrustTally`

NewTrustTally instantiates a new TrustTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustTallyWithDefaults

`func NewTrustTallyWithDefaults() *TrustTally`

NewTrustTallyWithDefaults instantiates a new TrustTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsent

`func (o *TrustTally) GetAbsent() int32`

GetAbsent returns the Absent field if non-nil, zero value otherwise.

### GetAbsentOk

`func (o *TrustTally) GetAbsentOk() (*int32, bool)`

GetAbsentOk returns a tuple with the Absent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsent

`func (o *TrustTally) SetAbsent(v int32)`

SetAbsent sets Absent field to given value.

### HasAbsent

`func (o *TrustTally) HasAbsent() bool`

HasAbsent returns a boolean if a field has been set.

### GetAutomated

`func (o *TrustTally) GetAutomated() int32`

GetAutomated returns the Automated field if non-nil, zero value otherwise.

### GetAutomatedOk

`func (o *TrustTally) GetAutomatedOk() (*int32, bool)`

GetAutomatedOk returns a tuple with the Automated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomated

`func (o *TrustTally) SetAutomated(v int32)`

SetAutomated sets Automated field to given value.

### HasAutomated

`func (o *TrustTally) HasAutomated() bool`

HasAutomated returns a boolean if a field has been set.

### GetPartial

`func (o *TrustTally) GetPartial() int32`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *TrustTally) GetPartialOk() (*int32, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *TrustTally) SetPartial(v int32)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *TrustTally) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetStatement

`func (o *TrustTally) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *TrustTally) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *TrustTally) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *TrustTally) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetTotal

`func (o *TrustTally) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TrustTally) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TrustTally) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TrustTally) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUnverified

`func (o *TrustTally) GetUnverified() int32`

GetUnverified returns the Unverified field if non-nil, zero value otherwise.

### GetUnverifiedOk

`func (o *TrustTally) GetUnverifiedOk() (*int32, bool)`

GetUnverifiedOk returns a tuple with the Unverified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnverified

`func (o *TrustTally) SetUnverified(v int32)`

SetUnverified sets Unverified field to given value.

### HasUnverified

`func (o *TrustTally) HasUnverified() bool`

HasUnverified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


