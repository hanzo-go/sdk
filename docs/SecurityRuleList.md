# SecurityRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SecurityRuleView**](SecurityRuleView.md) | Data is every rule a scan can fire, each with the id, name and severity a finding cites. | [optional] 

## Methods

### NewSecurityRuleList

`func NewSecurityRuleList() *SecurityRuleList`

NewSecurityRuleList instantiates a new SecurityRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityRuleListWithDefaults

`func NewSecurityRuleListWithDefaults() *SecurityRuleList`

NewSecurityRuleListWithDefaults instantiates a new SecurityRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SecurityRuleList) GetData() []SecurityRuleView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SecurityRuleList) GetDataOk() (*[]SecurityRuleView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SecurityRuleList) SetData(v []SecurityRuleView)`

SetData sets Data field to given value.

### HasData

`func (o *SecurityRuleList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


