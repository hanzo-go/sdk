# PrincipalMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Net** | Pointer to **string** | Net is what the payee receives, U.S. dollars. | [optional] 
**Rail** | Pointer to **string** | Rail is the rail: ledger or chain. | [optional] 
**Rule** | Pointer to [**PrincipalRule**](PrincipalRule.md) | Rule is why. | [optional] 
**Withheld** | Pointer to **string** | Withheld is what the payer keeps back and deposits itself, U.S. dollars. | [optional] 

## Methods

### NewPrincipalMethod

`func NewPrincipalMethod() *PrincipalMethod`

NewPrincipalMethod instantiates a new PrincipalMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalMethodWithDefaults

`func NewPrincipalMethodWithDefaults() *PrincipalMethod`

NewPrincipalMethodWithDefaults instantiates a new PrincipalMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNet

`func (o *PrincipalMethod) GetNet() string`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *PrincipalMethod) GetNetOk() (*string, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *PrincipalMethod) SetNet(v string)`

SetNet sets Net field to given value.

### HasNet

`func (o *PrincipalMethod) HasNet() bool`

HasNet returns a boolean if a field has been set.

### GetRail

`func (o *PrincipalMethod) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *PrincipalMethod) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *PrincipalMethod) SetRail(v string)`

SetRail sets Rail field to given value.

### HasRail

`func (o *PrincipalMethod) HasRail() bool`

HasRail returns a boolean if a field has been set.

### GetRule

`func (o *PrincipalMethod) GetRule() PrincipalRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *PrincipalMethod) GetRuleOk() (*PrincipalRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *PrincipalMethod) SetRule(v PrincipalRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *PrincipalMethod) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetWithheld

`func (o *PrincipalMethod) GetWithheld() string`

GetWithheld returns the Withheld field if non-nil, zero value otherwise.

### GetWithheldOk

`func (o *PrincipalMethod) GetWithheldOk() (*string, bool)`

GetWithheldOk returns a tuple with the Withheld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithheld

`func (o *PrincipalMethod) SetWithheld(v string)`

SetWithheld sets Withheld field to given value.

### HasWithheld

`func (o *PrincipalMethod) HasWithheld() bool`

HasWithheld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


