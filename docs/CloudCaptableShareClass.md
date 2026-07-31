# CloudCaptableShareClass

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

### NewCloudCaptableShareClass

`func NewCloudCaptableShareClass() *CloudCaptableShareClass`

NewCloudCaptableShareClass instantiates a new CloudCaptableShareClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableShareClassWithDefaults

`func NewCloudCaptableShareClassWithDefaults() *CloudCaptableShareClass`

NewCloudCaptableShareClassWithDefaults instantiates a new CloudCaptableShareClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassType

`func (o *CloudCaptableShareClass) GetClassType() string`

GetClassType returns the ClassType field if non-nil, zero value otherwise.

### GetClassTypeOk

`func (o *CloudCaptableShareClass) GetClassTypeOk() (*string, bool)`

GetClassTypeOk returns a tuple with the ClassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassType

`func (o *CloudCaptableShareClass) SetClassType(v string)`

SetClassType sets ClassType field to given value.

### HasClassType

`func (o *CloudCaptableShareClass) HasClassType() bool`

HasClassType returns a boolean if a field has been set.

### GetCompanyName

`func (o *CloudCaptableShareClass) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CloudCaptableShareClass) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CloudCaptableShareClass) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CloudCaptableShareClass) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetConversionRights

`func (o *CloudCaptableShareClass) GetConversionRights() string`

GetConversionRights returns the ConversionRights field if non-nil, zero value otherwise.

### GetConversionRightsOk

`func (o *CloudCaptableShareClass) GetConversionRightsOk() (*string, bool)`

GetConversionRightsOk returns a tuple with the ConversionRights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionRights

`func (o *CloudCaptableShareClass) SetConversionRights(v string)`

SetConversionRights sets ConversionRights field to given value.

### HasConversionRights

`func (o *CloudCaptableShareClass) HasConversionRights() bool`

HasConversionRights returns a boolean if a field has been set.

### GetId

`func (o *CloudCaptableShareClass) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCaptableShareClass) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCaptableShareClass) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCaptableShareClass) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdx

`func (o *CloudCaptableShareClass) GetIdx() int32`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *CloudCaptableShareClass) GetIdxOk() (*int32, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *CloudCaptableShareClass) SetIdx(v int32)`

SetIdx sets Idx field to given value.

### HasIdx

`func (o *CloudCaptableShareClass) HasIdx() bool`

HasIdx returns a boolean if a field has been set.

### GetInitialSharesAuthorized

`func (o *CloudCaptableShareClass) GetInitialSharesAuthorized() int32`

GetInitialSharesAuthorized returns the InitialSharesAuthorized field if non-nil, zero value otherwise.

### GetInitialSharesAuthorizedOk

`func (o *CloudCaptableShareClass) GetInitialSharesAuthorizedOk() (*int32, bool)`

GetInitialSharesAuthorizedOk returns a tuple with the InitialSharesAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSharesAuthorized

`func (o *CloudCaptableShareClass) SetInitialSharesAuthorized(v int32)`

SetInitialSharesAuthorized sets InitialSharesAuthorized field to given value.

### HasInitialSharesAuthorized

`func (o *CloudCaptableShareClass) HasInitialSharesAuthorized() bool`

HasInitialSharesAuthorized returns a boolean if a field has been set.

### GetLiquidationPreferenceMultiple

`func (o *CloudCaptableShareClass) GetLiquidationPreferenceMultiple() float32`

GetLiquidationPreferenceMultiple returns the LiquidationPreferenceMultiple field if non-nil, zero value otherwise.

### GetLiquidationPreferenceMultipleOk

`func (o *CloudCaptableShareClass) GetLiquidationPreferenceMultipleOk() (*float32, bool)`

GetLiquidationPreferenceMultipleOk returns a tuple with the LiquidationPreferenceMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationPreferenceMultiple

`func (o *CloudCaptableShareClass) SetLiquidationPreferenceMultiple(v float32)`

SetLiquidationPreferenceMultiple sets LiquidationPreferenceMultiple field to given value.

### HasLiquidationPreferenceMultiple

`func (o *CloudCaptableShareClass) HasLiquidationPreferenceMultiple() bool`

HasLiquidationPreferenceMultiple returns a boolean if a field has been set.

### GetName

`func (o *CloudCaptableShareClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCaptableShareClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCaptableShareClass) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCaptableShareClass) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParValue

`func (o *CloudCaptableShareClass) GetParValue() float32`

GetParValue returns the ParValue field if non-nil, zero value otherwise.

### GetParValueOk

`func (o *CloudCaptableShareClass) GetParValueOk() (*float32, bool)`

GetParValueOk returns a tuple with the ParValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParValue

`func (o *CloudCaptableShareClass) SetParValue(v float32)`

SetParValue sets ParValue field to given value.

### HasParValue

`func (o *CloudCaptableShareClass) HasParValue() bool`

HasParValue returns a boolean if a field has been set.

### GetParticipationCapMultiple

`func (o *CloudCaptableShareClass) GetParticipationCapMultiple() float32`

GetParticipationCapMultiple returns the ParticipationCapMultiple field if non-nil, zero value otherwise.

### GetParticipationCapMultipleOk

`func (o *CloudCaptableShareClass) GetParticipationCapMultipleOk() (*float32, bool)`

GetParticipationCapMultipleOk returns a tuple with the ParticipationCapMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipationCapMultiple

`func (o *CloudCaptableShareClass) SetParticipationCapMultiple(v float32)`

SetParticipationCapMultiple sets ParticipationCapMultiple field to given value.

### HasParticipationCapMultiple

`func (o *CloudCaptableShareClass) HasParticipationCapMultiple() bool`

HasParticipationCapMultiple returns a boolean if a field has been set.

### GetPrefix

`func (o *CloudCaptableShareClass) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CloudCaptableShareClass) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CloudCaptableShareClass) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *CloudCaptableShareClass) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPricePerShare

`func (o *CloudCaptableShareClass) GetPricePerShare() float32`

GetPricePerShare returns the PricePerShare field if non-nil, zero value otherwise.

### GetPricePerShareOk

`func (o *CloudCaptableShareClass) GetPricePerShareOk() (*float32, bool)`

GetPricePerShareOk returns a tuple with the PricePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerShare

`func (o *CloudCaptableShareClass) SetPricePerShare(v float32)`

SetPricePerShare sets PricePerShare field to given value.

### HasPricePerShare

`func (o *CloudCaptableShareClass) HasPricePerShare() bool`

HasPricePerShare returns a boolean if a field has been set.

### GetSeniority

`func (o *CloudCaptableShareClass) GetSeniority() int32`

GetSeniority returns the Seniority field if non-nil, zero value otherwise.

### GetSeniorityOk

`func (o *CloudCaptableShareClass) GetSeniorityOk() (*int32, bool)`

GetSeniorityOk returns a tuple with the Seniority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeniority

`func (o *CloudCaptableShareClass) SetSeniority(v int32)`

SetSeniority sets Seniority field to given value.

### HasSeniority

`func (o *CloudCaptableShareClass) HasSeniority() bool`

HasSeniority returns a boolean if a field has been set.

### GetVotesPerShare

`func (o *CloudCaptableShareClass) GetVotesPerShare() int32`

GetVotesPerShare returns the VotesPerShare field if non-nil, zero value otherwise.

### GetVotesPerShareOk

`func (o *CloudCaptableShareClass) GetVotesPerShareOk() (*int32, bool)`

GetVotesPerShareOk returns a tuple with the VotesPerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesPerShare

`func (o *CloudCaptableShareClass) SetVotesPerShare(v int32)`

SetVotesPerShare sets VotesPerShare field to given value.

### HasVotesPerShare

`func (o *CloudCaptableShareClass) HasVotesPerShare() bool`

HasVotesPerShare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


