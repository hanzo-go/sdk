# TrustCentre

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controls** | Pointer to **[]interface{}** | Controls is the control inventory, each entry naming what it asserts, the mechanism, where it is enforced, how it is verified and the clauses it maps to. | [optional] 
**Coverage** | Pointer to [**[]TrustCoverRow**](TrustCoverRow.md) | Coverage is the per-framework counts, computed from Controls against each framework&#39;s whole published clause list. | [optional] 
**Documents** | Pointer to [**[]TrustDocRow**](TrustDocRow.md) | Documents are the artifacts, each saying whether this reader may read it. | [optional] 
**Faq** | Pointer to **[]interface{}** | Faq is the knowledge base — the questions a reviewer asks, answered. | [optional] 
**Frameworks** | Pointer to [**[]TrustFrameworkRow**](TrustFrameworkRow.md) | Frameworks are the clause universes the coverage is computed against. | [optional] 
**Generated** | Pointer to **int64** | Generated is when this answer was computed, unix milliseconds. | [optional] 
**Inventory** | Pointer to [**TrustTrustTally**](TrustTrustTally.md) | Inventory is how the controls themselves stand, independent of framework. | [optional] 
**Org** | Pointer to **string** | Org is whose centre this is. | [optional] 
**Policies** | Pointer to **[]interface{}** | Policies are the published policies. | [optional] 
**Profile** | Pointer to **interface{}** |  | [optional] 
**Risk** | Pointer to **interface{}** |  | [optional] 
**Subprocessors** | Pointer to **[]interface{}** | Subprocessors are the third parties this organization sends data to. | [optional] 
**Updates** | Pointer to **[]interface{}** | Updates is the changelog, newest as the organization ordered it. | [optional] 
**Version** | Pointer to **string** | Version is the embedded inventory&#39;s version. | [optional] 

## Methods

### NewTrustCentre

`func NewTrustCentre() *TrustCentre`

NewTrustCentre instantiates a new TrustCentre object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustCentreWithDefaults

`func NewTrustCentreWithDefaults() *TrustCentre`

NewTrustCentreWithDefaults instantiates a new TrustCentre object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControls

`func (o *TrustCentre) GetControls() []interface{}`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *TrustCentre) GetControlsOk() (*[]interface{}, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *TrustCentre) SetControls(v []interface{})`

SetControls sets Controls field to given value.

### HasControls

`func (o *TrustCentre) HasControls() bool`

HasControls returns a boolean if a field has been set.

### GetCoverage

`func (o *TrustCentre) GetCoverage() []TrustCoverRow`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *TrustCentre) GetCoverageOk() (*[]TrustCoverRow, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *TrustCentre) SetCoverage(v []TrustCoverRow)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *TrustCentre) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetDocuments

`func (o *TrustCentre) GetDocuments() []TrustDocRow`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *TrustCentre) GetDocumentsOk() (*[]TrustDocRow, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *TrustCentre) SetDocuments(v []TrustDocRow)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *TrustCentre) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetFaq

`func (o *TrustCentre) GetFaq() []interface{}`

GetFaq returns the Faq field if non-nil, zero value otherwise.

### GetFaqOk

`func (o *TrustCentre) GetFaqOk() (*[]interface{}, bool)`

GetFaqOk returns a tuple with the Faq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaq

`func (o *TrustCentre) SetFaq(v []interface{})`

SetFaq sets Faq field to given value.

### HasFaq

`func (o *TrustCentre) HasFaq() bool`

HasFaq returns a boolean if a field has been set.

### GetFrameworks

`func (o *TrustCentre) GetFrameworks() []TrustFrameworkRow`

GetFrameworks returns the Frameworks field if non-nil, zero value otherwise.

### GetFrameworksOk

`func (o *TrustCentre) GetFrameworksOk() (*[]TrustFrameworkRow, bool)`

GetFrameworksOk returns a tuple with the Frameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworks

`func (o *TrustCentre) SetFrameworks(v []TrustFrameworkRow)`

SetFrameworks sets Frameworks field to given value.

### HasFrameworks

`func (o *TrustCentre) HasFrameworks() bool`

HasFrameworks returns a boolean if a field has been set.

### GetGenerated

`func (o *TrustCentre) GetGenerated() int64`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *TrustCentre) GetGeneratedOk() (*int64, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *TrustCentre) SetGenerated(v int64)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *TrustCentre) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetInventory

`func (o *TrustCentre) GetInventory() TrustTrustTally`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *TrustCentre) GetInventoryOk() (*TrustTrustTally, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *TrustCentre) SetInventory(v TrustTrustTally)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *TrustCentre) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetOrg

`func (o *TrustCentre) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *TrustCentre) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *TrustCentre) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *TrustCentre) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPolicies

`func (o *TrustCentre) GetPolicies() []interface{}`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *TrustCentre) GetPoliciesOk() (*[]interface{}, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *TrustCentre) SetPolicies(v []interface{})`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *TrustCentre) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProfile

`func (o *TrustCentre) GetProfile() interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *TrustCentre) GetProfileOk() (*interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *TrustCentre) SetProfile(v interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *TrustCentre) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *TrustCentre) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *TrustCentre) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetRisk

`func (o *TrustCentre) GetRisk() interface{}`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *TrustCentre) GetRiskOk() (*interface{}, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *TrustCentre) SetRisk(v interface{})`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *TrustCentre) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### SetRiskNil

`func (o *TrustCentre) SetRiskNil(b bool)`

 SetRiskNil sets the value for Risk to be an explicit nil

### UnsetRisk
`func (o *TrustCentre) UnsetRisk()`

UnsetRisk ensures that no value is present for Risk, not even an explicit nil
### GetSubprocessors

`func (o *TrustCentre) GetSubprocessors() []interface{}`

GetSubprocessors returns the Subprocessors field if non-nil, zero value otherwise.

### GetSubprocessorsOk

`func (o *TrustCentre) GetSubprocessorsOk() (*[]interface{}, bool)`

GetSubprocessorsOk returns a tuple with the Subprocessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessors

`func (o *TrustCentre) SetSubprocessors(v []interface{})`

SetSubprocessors sets Subprocessors field to given value.

### HasSubprocessors

`func (o *TrustCentre) HasSubprocessors() bool`

HasSubprocessors returns a boolean if a field has been set.

### GetUpdates

`func (o *TrustCentre) GetUpdates() []interface{}`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *TrustCentre) GetUpdatesOk() (*[]interface{}, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *TrustCentre) SetUpdates(v []interface{})`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *TrustCentre) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetVersion

`func (o *TrustCentre) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrustCentre) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrustCentre) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TrustCentre) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


