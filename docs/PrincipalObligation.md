# PrincipalObligation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authority** | Pointer to **string** | Authority is who it is filed with. | [optional] 
**Code** | Pointer to **string** | Code is the obligation&#39;s stable name. | [optional] 
**Due** | Pointer to **string** | Due is when. | [optional] 
**Filer** | Pointer to **string** | Filer is the org that files it. It is never Hanzo: Hanzo prepares what an org files, and files nothing itself. | [optional] 
**Form** | Pointer to **string** | Form is what is filed. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is where, ISO 3166-1 alpha-2 or \&quot;EU\&quot;. | [optional] 
**Rule** | Pointer to [**PrincipalRule**](PrincipalRule.md) | Rule is why. | [optional] 
**When** | Pointer to **string** | When is the condition it is owed on, when it is owed on one. | [optional] 

## Methods

### NewPrincipalObligation

`func NewPrincipalObligation() *PrincipalObligation`

NewPrincipalObligation instantiates a new PrincipalObligation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalObligationWithDefaults

`func NewPrincipalObligationWithDefaults() *PrincipalObligation`

NewPrincipalObligationWithDefaults instantiates a new PrincipalObligation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthority

`func (o *PrincipalObligation) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *PrincipalObligation) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *PrincipalObligation) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *PrincipalObligation) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetCode

`func (o *PrincipalObligation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PrincipalObligation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PrincipalObligation) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PrincipalObligation) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDue

`func (o *PrincipalObligation) GetDue() string`

GetDue returns the Due field if non-nil, zero value otherwise.

### GetDueOk

`func (o *PrincipalObligation) GetDueOk() (*string, bool)`

GetDueOk returns a tuple with the Due field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDue

`func (o *PrincipalObligation) SetDue(v string)`

SetDue sets Due field to given value.

### HasDue

`func (o *PrincipalObligation) HasDue() bool`

HasDue returns a boolean if a field has been set.

### GetFiler

`func (o *PrincipalObligation) GetFiler() string`

GetFiler returns the Filer field if non-nil, zero value otherwise.

### GetFilerOk

`func (o *PrincipalObligation) GetFilerOk() (*string, bool)`

GetFilerOk returns a tuple with the Filer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiler

`func (o *PrincipalObligation) SetFiler(v string)`

SetFiler sets Filer field to given value.

### HasFiler

`func (o *PrincipalObligation) HasFiler() bool`

HasFiler returns a boolean if a field has been set.

### GetForm

`func (o *PrincipalObligation) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *PrincipalObligation) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *PrincipalObligation) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *PrincipalObligation) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetJurisdiction

`func (o *PrincipalObligation) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *PrincipalObligation) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *PrincipalObligation) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *PrincipalObligation) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetRule

`func (o *PrincipalObligation) GetRule() PrincipalRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *PrincipalObligation) GetRuleOk() (*PrincipalRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *PrincipalObligation) SetRule(v PrincipalRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *PrincipalObligation) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetWhen

`func (o *PrincipalObligation) GetWhen() string`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *PrincipalObligation) GetWhenOk() (*string, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *PrincipalObligation) SetWhen(v string)`

SetWhen sets When field to given value.

### HasWhen

`func (o *PrincipalObligation) HasWhen() bool`

HasWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


