# MarketplaceOnboarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]MarketplaceStatement**](MarketplaceStatement.md) | Credentials are the TaxPrincipalCredentials signed for the org&#39;s agents, newest first. | [optional] 
**Earnings** | Pointer to [**MarketplaceEarnings**](MarketplaceEarnings.md) | Earnings is what the org was paid in Year, from the economic events the rails stated. | [optional] 
**Entity** | Pointer to **string** | Entity is the legal entity&#39;s formation stage (KYB) — \&quot;company\&quot; once formed — or empty when the org began none. | [optional] 
**Identity** | Pointer to **string** | Identity is the founders&#39; identity verification (KYC): verified, pending, failed or none. | [optional] 
**Missing** | Pointer to **[]string** | Missing names what the org still owes before it is paid without a hitch: identity, sanctions, tax_form, tax_invalid or payout. | [optional] 
**Org** | Pointer to **string** | Org is the seller org. | [optional] 
**Payout** | Pointer to [**MarketplacePayout**](MarketplacePayout.md) | Payout is the wallet the org proved it is paid into. | [optional] 
**Ready** | Pointer to **bool** | Ready is true when nothing is missing. | [optional] 
**Received** | Pointer to [**[]MarketplaceReceived**](MarketplaceReceived.md) | Received are the 1099s payers furnished the org for Year. | [optional] 
**Sanctions** | Pointer to **string** | Sanctions is the org&#39;s own screening: clear, review, blocked, unscreened or unavailable, and why in words. | [optional] 
**SanctionsReason** | Pointer to **string** | SanctionsReason says why Sanctions is what it is. | [optional] 
**Sources** | Pointer to [**[]MarketplaceSource**](MarketplaceSource.md) | Sources names each owning app asked, and whether it answered: ok, absent (not run by this deployment) or failed. | [optional] 
**Tax** | Pointer to [**MarketplaceTaxStatus**](MarketplaceTaxStatus.md) | Tax is the org&#39;s own tax form, when it has one. | [optional] 

## Methods

### NewMarketplaceOnboarding

`func NewMarketplaceOnboarding() *MarketplaceOnboarding`

NewMarketplaceOnboarding instantiates a new MarketplaceOnboarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceOnboardingWithDefaults

`func NewMarketplaceOnboardingWithDefaults() *MarketplaceOnboarding`

NewMarketplaceOnboardingWithDefaults instantiates a new MarketplaceOnboarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *MarketplaceOnboarding) GetCredentials() []MarketplaceStatement`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *MarketplaceOnboarding) GetCredentialsOk() (*[]MarketplaceStatement, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *MarketplaceOnboarding) SetCredentials(v []MarketplaceStatement)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *MarketplaceOnboarding) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEarnings

`func (o *MarketplaceOnboarding) GetEarnings() MarketplaceEarnings`

GetEarnings returns the Earnings field if non-nil, zero value otherwise.

### GetEarningsOk

`func (o *MarketplaceOnboarding) GetEarningsOk() (*MarketplaceEarnings, bool)`

GetEarningsOk returns a tuple with the Earnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnings

`func (o *MarketplaceOnboarding) SetEarnings(v MarketplaceEarnings)`

SetEarnings sets Earnings field to given value.

### HasEarnings

`func (o *MarketplaceOnboarding) HasEarnings() bool`

HasEarnings returns a boolean if a field has been set.

### GetEntity

`func (o *MarketplaceOnboarding) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *MarketplaceOnboarding) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *MarketplaceOnboarding) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *MarketplaceOnboarding) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetIdentity

`func (o *MarketplaceOnboarding) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *MarketplaceOnboarding) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *MarketplaceOnboarding) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *MarketplaceOnboarding) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetMissing

`func (o *MarketplaceOnboarding) GetMissing() []string`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *MarketplaceOnboarding) GetMissingOk() (*[]string, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *MarketplaceOnboarding) SetMissing(v []string)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *MarketplaceOnboarding) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetOrg

`func (o *MarketplaceOnboarding) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *MarketplaceOnboarding) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *MarketplaceOnboarding) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *MarketplaceOnboarding) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPayout

`func (o *MarketplaceOnboarding) GetPayout() MarketplacePayout`

GetPayout returns the Payout field if non-nil, zero value otherwise.

### GetPayoutOk

`func (o *MarketplaceOnboarding) GetPayoutOk() (*MarketplacePayout, bool)`

GetPayoutOk returns a tuple with the Payout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayout

`func (o *MarketplaceOnboarding) SetPayout(v MarketplacePayout)`

SetPayout sets Payout field to given value.

### HasPayout

`func (o *MarketplaceOnboarding) HasPayout() bool`

HasPayout returns a boolean if a field has been set.

### GetReady

`func (o *MarketplaceOnboarding) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *MarketplaceOnboarding) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *MarketplaceOnboarding) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *MarketplaceOnboarding) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetReceived

`func (o *MarketplaceOnboarding) GetReceived() []MarketplaceReceived`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *MarketplaceOnboarding) GetReceivedOk() (*[]MarketplaceReceived, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *MarketplaceOnboarding) SetReceived(v []MarketplaceReceived)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *MarketplaceOnboarding) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetSanctions

`func (o *MarketplaceOnboarding) GetSanctions() string`

GetSanctions returns the Sanctions field if non-nil, zero value otherwise.

### GetSanctionsOk

`func (o *MarketplaceOnboarding) GetSanctionsOk() (*string, bool)`

GetSanctionsOk returns a tuple with the Sanctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctions

`func (o *MarketplaceOnboarding) SetSanctions(v string)`

SetSanctions sets Sanctions field to given value.

### HasSanctions

`func (o *MarketplaceOnboarding) HasSanctions() bool`

HasSanctions returns a boolean if a field has been set.

### GetSanctionsReason

`func (o *MarketplaceOnboarding) GetSanctionsReason() string`

GetSanctionsReason returns the SanctionsReason field if non-nil, zero value otherwise.

### GetSanctionsReasonOk

`func (o *MarketplaceOnboarding) GetSanctionsReasonOk() (*string, bool)`

GetSanctionsReasonOk returns a tuple with the SanctionsReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanctionsReason

`func (o *MarketplaceOnboarding) SetSanctionsReason(v string)`

SetSanctionsReason sets SanctionsReason field to given value.

### HasSanctionsReason

`func (o *MarketplaceOnboarding) HasSanctionsReason() bool`

HasSanctionsReason returns a boolean if a field has been set.

### GetSources

`func (o *MarketplaceOnboarding) GetSources() []MarketplaceSource`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *MarketplaceOnboarding) GetSourcesOk() (*[]MarketplaceSource, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *MarketplaceOnboarding) SetSources(v []MarketplaceSource)`

SetSources sets Sources field to given value.

### HasSources

`func (o *MarketplaceOnboarding) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTax

`func (o *MarketplaceOnboarding) GetTax() MarketplaceTaxStatus`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *MarketplaceOnboarding) GetTaxOk() (*MarketplaceTaxStatus, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *MarketplaceOnboarding) SetTax(v MarketplaceTaxStatus)`

SetTax sets Tax field to given value.

### HasTax

`func (o *MarketplaceOnboarding) HasTax() bool`

HasTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


