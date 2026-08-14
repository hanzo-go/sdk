# CaptableShareClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassType** | Pointer to **string** | ClassType is COMMON or PREFERRED. | [optional] 
**CompanyName** | Pointer to **string** | CompanyName is the name of the company whose cap table this is. | [optional] 
**ConversionRights** | Pointer to **string** | ConversionRights describes what the class converts into, e.g. CONVERTS_TO_FUTURE_ROUND. | [optional] 
**Id** | Pointer to **string** | ID is the share class id. | [optional] 
**Idx** | Pointer to **int32** | Idx is the class&#39;s 1-based position within the company, in creation order. | [optional] 
**InitialSharesAuthorized** | Pointer to **int32** | InitialSharesAuthorized is how many shares of this class are authorized. | [optional] 
**LiquidationPreferenceMultiple** | Pointer to **float32** | LiquidationPreferenceMultiple is the preference multiple on liquidation. | [optional] 
**Name** | Pointer to **string** | Name is the class name, e.g. \&quot;Common\&quot; or \&quot;Series A Preferred\&quot;. | [optional] 
**ParValue** | Pointer to **float32** | ParValue is the par value per share. | [optional] 
**ParticipationCapMultiple** | Pointer to **float32** | ParticipationCapMultiple caps participation on liquidation; 0 is uncapped. | [optional] 
**Prefix** | Pointer to **string** | Prefix is the certificate prefix, CS for common and PS for preferred. | [optional] 
**PricePerShare** | Pointer to **float32** | PricePerShare is the issue price per share. | [optional] 
**Seniority** | Pointer to **int32** | Seniority orders classes in a liquidation waterfall; higher is more senior. | [optional] 
**VotesPerShare** | Pointer to **int32** | VotesPerShare is how many votes one share of this class carries. | [optional] 

## Methods

### NewCaptableShareClass

`func NewCaptableShareClass() *CaptableShareClass`

NewCaptableShareClass instantiates a new CaptableShareClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableShareClassWithDefaults

`func NewCaptableShareClassWithDefaults() *CaptableShareClass`

NewCaptableShareClassWithDefaults instantiates a new CaptableShareClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassType

`func (o *CaptableShareClass) GetClassType() string`

GetClassType returns the ClassType field if non-nil, zero value otherwise.

### GetClassTypeOk

`func (o *CaptableShareClass) GetClassTypeOk() (*string, bool)`

GetClassTypeOk returns a tuple with the ClassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassType

`func (o *CaptableShareClass) SetClassType(v string)`

SetClassType sets ClassType field to given value.

### HasClassType

`func (o *CaptableShareClass) HasClassType() bool`

HasClassType returns a boolean if a field has been set.

### GetCompanyName

`func (o *CaptableShareClass) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CaptableShareClass) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CaptableShareClass) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CaptableShareClass) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetConversionRights

`func (o *CaptableShareClass) GetConversionRights() string`

GetConversionRights returns the ConversionRights field if non-nil, zero value otherwise.

### GetConversionRightsOk

`func (o *CaptableShareClass) GetConversionRightsOk() (*string, bool)`

GetConversionRightsOk returns a tuple with the ConversionRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRights

`func (o *CaptableShareClass) SetConversionRights(v string)`

SetConversionRights sets ConversionRights field to given value.

### HasConversionRights

`func (o *CaptableShareClass) HasConversionRights() bool`

HasConversionRights returns a boolean if a field has been set.

### GetId

`func (o *CaptableShareClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaptableShareClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaptableShareClass) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaptableShareClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdx

`func (o *CaptableShareClass) GetIdx() int32`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *CaptableShareClass) GetIdxOk() (*int32, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *CaptableShareClass) SetIdx(v int32)`

SetIdx sets Idx field to given value.

### HasIdx

`func (o *CaptableShareClass) HasIdx() bool`

HasIdx returns a boolean if a field has been set.

### GetInitialSharesAuthorized

`func (o *CaptableShareClass) GetInitialSharesAuthorized() int32`

GetInitialSharesAuthorized returns the InitialSharesAuthorized field if non-nil, zero value otherwise.

### GetInitialSharesAuthorizedOk

`func (o *CaptableShareClass) GetInitialSharesAuthorizedOk() (*int32, bool)`

GetInitialSharesAuthorizedOk returns a tuple with the InitialSharesAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSharesAuthorized

`func (o *CaptableShareClass) SetInitialSharesAuthorized(v int32)`

SetInitialSharesAuthorized sets InitialSharesAuthorized field to given value.

### HasInitialSharesAuthorized

`func (o *CaptableShareClass) HasInitialSharesAuthorized() bool`

HasInitialSharesAuthorized returns a boolean if a field has been set.

### GetLiquidationPreferenceMultiple

`func (o *CaptableShareClass) GetLiquidationPreferenceMultiple() float32`

GetLiquidationPreferenceMultiple returns the LiquidationPreferenceMultiple field if non-nil, zero value otherwise.

### GetLiquidationPreferenceMultipleOk

`func (o *CaptableShareClass) GetLiquidationPreferenceMultipleOk() (*float32, bool)`

GetLiquidationPreferenceMultipleOk returns a tuple with the LiquidationPreferenceMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationPreferenceMultiple

`func (o *CaptableShareClass) SetLiquidationPreferenceMultiple(v float32)`

SetLiquidationPreferenceMultiple sets LiquidationPreferenceMultiple field to given value.

### HasLiquidationPreferenceMultiple

`func (o *CaptableShareClass) HasLiquidationPreferenceMultiple() bool`

HasLiquidationPreferenceMultiple returns a boolean if a field has been set.

### GetName

`func (o *CaptableShareClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptableShareClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptableShareClass) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaptableShareClass) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParValue

`func (o *CaptableShareClass) GetParValue() float32`

GetParValue returns the ParValue field if non-nil, zero value otherwise.

### GetParValueOk

`func (o *CaptableShareClass) GetParValueOk() (*float32, bool)`

GetParValueOk returns a tuple with the ParValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParValue

`func (o *CaptableShareClass) SetParValue(v float32)`

SetParValue sets ParValue field to given value.

### HasParValue

`func (o *CaptableShareClass) HasParValue() bool`

HasParValue returns a boolean if a field has been set.

### GetParticipationCapMultiple

`func (o *CaptableShareClass) GetParticipationCapMultiple() float32`

GetParticipationCapMultiple returns the ParticipationCapMultiple field if non-nil, zero value otherwise.

### GetParticipationCapMultipleOk

`func (o *CaptableShareClass) GetParticipationCapMultipleOk() (*float32, bool)`

GetParticipationCapMultipleOk returns a tuple with the ParticipationCapMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationCapMultiple

`func (o *CaptableShareClass) SetParticipationCapMultiple(v float32)`

SetParticipationCapMultiple sets ParticipationCapMultiple field to given value.

### HasParticipationCapMultiple

`func (o *CaptableShareClass) HasParticipationCapMultiple() bool`

HasParticipationCapMultiple returns a boolean if a field has been set.

### GetPrefix

`func (o *CaptableShareClass) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CaptableShareClass) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CaptableShareClass) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *CaptableShareClass) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPricePerShare

`func (o *CaptableShareClass) GetPricePerShare() float32`

GetPricePerShare returns the PricePerShare field if non-nil, zero value otherwise.

### GetPricePerShareOk

`func (o *CaptableShareClass) GetPricePerShareOk() (*float32, bool)`

GetPricePerShareOk returns a tuple with the PricePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerShare

`func (o *CaptableShareClass) SetPricePerShare(v float32)`

SetPricePerShare sets PricePerShare field to given value.

### HasPricePerShare

`func (o *CaptableShareClass) HasPricePerShare() bool`

HasPricePerShare returns a boolean if a field has been set.

### GetSeniority

`func (o *CaptableShareClass) GetSeniority() int32`

GetSeniority returns the Seniority field if non-nil, zero value otherwise.

### GetSeniorityOk

`func (o *CaptableShareClass) GetSeniorityOk() (*int32, bool)`

GetSeniorityOk returns a tuple with the Seniority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeniority

`func (o *CaptableShareClass) SetSeniority(v int32)`

SetSeniority sets Seniority field to given value.

### HasSeniority

`func (o *CaptableShareClass) HasSeniority() bool`

HasSeniority returns a boolean if a field has been set.

### GetVotesPerShare

`func (o *CaptableShareClass) GetVotesPerShare() int32`

GetVotesPerShare returns the VotesPerShare field if non-nil, zero value otherwise.

### GetVotesPerShareOk

`func (o *CaptableShareClass) GetVotesPerShareOk() (*int32, bool)`

GetVotesPerShareOk returns a tuple with the VotesPerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesPerShare

`func (o *CaptableShareClass) SetVotesPerShare(v int32)`

SetVotesPerShare sets VotesPerShare field to given value.

### HasVotesPerShare

`func (o *CaptableShareClass) HasVotesPerShare() bool`

HasVotesPerShare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


