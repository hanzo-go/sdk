# Centre

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controls** | Pointer to **[]interface{}** | Controls is the control inventory, each entry naming what it asserts, the mechanism, where it is enforced, how it is verified and the clauses it maps to. | [optional] 
**Coverage** | Pointer to [**[]CoverRow**](CoverRow.md) | Coverage is the per-framework counts, computed from Controls against each framework&#39;s whole published clause list. | [optional] 
**Documents** | Pointer to [**[]DocRow**](DocRow.md) | Documents are the artifacts, each saying whether this reader may read it. | [optional] 
**Faq** | Pointer to **[]interface{}** | Faq is the knowledge base — the questions a reviewer asks, answered. | [optional] 
**Frameworks** | Pointer to [**[]FrameworkRow**](FrameworkRow.md) | Frameworks are the clause universes the coverage is computed against. | [optional] 
**Generated** | Pointer to **int32** | Generated is when this answer was computed, unix milliseconds. | [optional] 
**Inventory** | Pointer to [**TrustTally**](TrustTally.md) | Inventory is how the controls themselves stand, independent of framework. | [optional] 
**Org** | Pointer to **string** | Org is whose centre this is. | [optional] 
**Policies** | Pointer to **[]interface{}** | Policies are the published policies. | [optional] 
**Profile** | Pointer to **interface{}** |  | [optional] 
**Risk** | Pointer to **interface{}** |  | [optional] 
**Subprocessors** | Pointer to **[]interface{}** | Subprocessors are the third parties this organization sends data to. | [optional] 
**Updates** | Pointer to **[]interface{}** | Updates is the changelog, newest as the organization ordered it. | [optional] 
**Version** | Pointer to **string** | Version is the embedded inventory&#39;s version. | [optional] 

## Methods

### NewCentre

`func NewCentre() *Centre`

NewCentre instantiates a new Centre object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentreWithDefaults

`func NewCentreWithDefaults() *Centre`

NewCentreWithDefaults instantiates a new Centre object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControls

`func (o *Centre) GetControls() []interface{}`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *Centre) GetControlsOk() (*[]interface{}, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *Centre) SetControls(v []interface{})`

SetControls sets Controls field to given value.

### HasControls

`func (o *Centre) HasControls() bool`

HasControls returns a boolean if a field has been set.

### GetCoverage

`func (o *Centre) GetCoverage() []CoverRow`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *Centre) GetCoverageOk() (*[]CoverRow, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *Centre) SetCoverage(v []CoverRow)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *Centre) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetDocuments

`func (o *Centre) GetDocuments() []DocRow`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Centre) GetDocumentsOk() (*[]DocRow, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Centre) SetDocuments(v []DocRow)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Centre) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetFaq

`func (o *Centre) GetFaq() []interface{}`

GetFaq returns the Faq field if non-nil, zero value otherwise.

### GetFaqOk

`func (o *Centre) GetFaqOk() (*[]interface{}, bool)`

GetFaqOk returns a tuple with the Faq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaq

`func (o *Centre) SetFaq(v []interface{})`

SetFaq sets Faq field to given value.

### HasFaq

`func (o *Centre) HasFaq() bool`

HasFaq returns a boolean if a field has been set.

### GetFrameworks

`func (o *Centre) GetFrameworks() []FrameworkRow`

GetFrameworks returns the Frameworks field if non-nil, zero value otherwise.

### GetFrameworksOk

`func (o *Centre) GetFrameworksOk() (*[]FrameworkRow, bool)`

GetFrameworksOk returns a tuple with the Frameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworks

`func (o *Centre) SetFrameworks(v []FrameworkRow)`

SetFrameworks sets Frameworks field to given value.

### HasFrameworks

`func (o *Centre) HasFrameworks() bool`

HasFrameworks returns a boolean if a field has been set.

### GetGenerated

`func (o *Centre) GetGenerated() int32`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *Centre) GetGeneratedOk() (*int32, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *Centre) SetGenerated(v int32)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *Centre) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetInventory

`func (o *Centre) GetInventory() TrustTally`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *Centre) GetInventoryOk() (*TrustTally, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *Centre) SetInventory(v TrustTally)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *Centre) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetOrg

`func (o *Centre) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Centre) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Centre) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Centre) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPolicies

`func (o *Centre) GetPolicies() []interface{}`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *Centre) GetPoliciesOk() (*[]interface{}, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *Centre) SetPolicies(v []interface{})`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *Centre) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProfile

`func (o *Centre) GetProfile() interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Centre) GetProfileOk() (*interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Centre) SetProfile(v interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Centre) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *Centre) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *Centre) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetRisk

`func (o *Centre) GetRisk() interface{}`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *Centre) GetRiskOk() (*interface{}, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *Centre) SetRisk(v interface{})`

SetRisk sets Risk field to given value.

### HasRisk

`func (o *Centre) HasRisk() bool`

HasRisk returns a boolean if a field has been set.

### SetRiskNil

`func (o *Centre) SetRiskNil(b bool)`

 SetRiskNil sets the value for Risk to be an explicit nil

### UnsetRisk
`func (o *Centre) UnsetRisk()`

UnsetRisk ensures that no value is present for Risk, not even an explicit nil
### GetSubprocessors

`func (o *Centre) GetSubprocessors() []interface{}`

GetSubprocessors returns the Subprocessors field if non-nil, zero value otherwise.

### GetSubprocessorsOk

`func (o *Centre) GetSubprocessorsOk() (*[]interface{}, bool)`

GetSubprocessorsOk returns a tuple with the Subprocessors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubprocessors

`func (o *Centre) SetSubprocessors(v []interface{})`

SetSubprocessors sets Subprocessors field to given value.

### HasSubprocessors

`func (o *Centre) HasSubprocessors() bool`

HasSubprocessors returns a boolean if a field has been set.

### GetUpdates

`func (o *Centre) GetUpdates() []interface{}`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *Centre) GetUpdatesOk() (*[]interface{}, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *Centre) SetUpdates(v []interface{})`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *Centre) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetVersion

`func (o *Centre) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Centre) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Centre) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Centre) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


