# PrincipalHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** | Kind is name (matched fuzzily), address (matched exactly) or jurisdiction. | [optional] 
**Role** | Pointer to **string** | Role is where it came from: entity, founder, form, wallet, or — for a country — citizenship, organization or residence. | [optional] 
**Subject** | Pointer to **string** | Subject is the name, address or country as the org stated it. | [optional] 

## Methods

### NewPrincipalHit

`func NewPrincipalHit() *PrincipalHit`

NewPrincipalHit instantiates a new PrincipalHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalHitWithDefaults

`func NewPrincipalHitWithDefaults() *PrincipalHit`

NewPrincipalHitWithDefaults instantiates a new PrincipalHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PrincipalHit) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PrincipalHit) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PrincipalHit) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PrincipalHit) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRole

`func (o *PrincipalHit) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PrincipalHit) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PrincipalHit) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PrincipalHit) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubject

`func (o *PrincipalHit) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PrincipalHit) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PrincipalHit) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *PrincipalHit) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


