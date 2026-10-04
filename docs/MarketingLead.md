# MarketingLead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lead** | Pointer to **string** | Lead is the lead&#39;s reference, \&quot;lead_\&quot; and 24 hex characters. It is random, not the lead&#39;s internal name, so it says nothing about how many leads the team has. | [optional] 
**Stage** | Pointer to **string** | Stage is where the lead stands. Always \&quot;new\&quot; on a fresh lead. | [optional] 

## Methods

### NewMarketingLead

`func NewMarketingLead() *MarketingLead`

NewMarketingLead instantiates a new MarketingLead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingLeadWithDefaults

`func NewMarketingLeadWithDefaults() *MarketingLead`

NewMarketingLeadWithDefaults instantiates a new MarketingLead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLead

`func (o *MarketingLead) GetLead() string`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *MarketingLead) GetLeadOk() (*string, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *MarketingLead) SetLead(v string)`

SetLead sets Lead field to given value.

### HasLead

`func (o *MarketingLead) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetStage

`func (o *MarketingLead) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *MarketingLead) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *MarketingLead) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *MarketingLead) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


