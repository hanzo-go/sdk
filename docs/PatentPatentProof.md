# PatentPatentProof

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anchor** | Pointer to [**PatentPatentAnchor**](PatentPatentAnchor.md) | PatentAnchor is the external standing of this exact root. | [optional] 
**Coverage** | Pointer to [**PatentPatentCoverage**](PatentPatentCoverage.md) | Coverage counts what the fold held, per kind, so a short list reads as a short log and not as a complete one. | [optional] 
**Gaps** | Pointer to [**[]PatentPatentGap**](PatentPatentGap.md) | Gaps names every field of the covered records that states nothing, with how many records leave it unstated. | [optional] 
**Order** | Pointer to [**[]PatentPatentOrderCheck**](PatentPatentOrderCheck.md) | Order is one verdict per implementation-idea pair. | [optional] 
**Project** | Pointer to **string** | Project is the scope the fold covered. | [optional] 
**Records** | Pointer to [**[]PatentPatentProofEntry**](PatentPatentProofEntry.md) | Records are the entries folded, oldest clock first. | [optional] 
**Root** | Pointer to **string** | Root is the fold digest over every record in append order. | [optional] 
**Verdict** | Pointer to **string** | Verdict is the whole answer in one word: supported when every pair is &#x60;before&#x60;, contested when any is &#x60;after&#x60;, unproven when any is missing, and empty when there was nothing to fold. It is the WORST pair, because a chain with one inverted edge is not a chain. | [optional] 
**VerifiedAt** | Pointer to **string** | VerifiedAt is when this fold was computed. The records can have grown since; the root is the set as of here. | [optional] 

## Methods

### NewPatentPatentProof

`func NewPatentPatentProof() *PatentPatentProof`

NewPatentPatentProof instantiates a new PatentPatentProof object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentProofWithDefaults

`func NewPatentPatentProofWithDefaults() *PatentPatentProof`

NewPatentPatentProofWithDefaults instantiates a new PatentPatentProof object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnchor

`func (o *PatentPatentProof) GetAnchor() PatentPatentAnchor`

GetAnchor returns the Anchor field if non-nil, zero value otherwise.

### GetAnchorOk

`func (o *PatentPatentProof) GetAnchorOk() (*PatentPatentAnchor, bool)`

GetAnchorOk returns a tuple with the Anchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchor

`func (o *PatentPatentProof) SetAnchor(v PatentPatentAnchor)`

SetAnchor sets Anchor field to given value.

### HasAnchor

`func (o *PatentPatentProof) HasAnchor() bool`

HasAnchor returns a boolean if a field has been set.

### GetCoverage

`func (o *PatentPatentProof) GetCoverage() PatentPatentCoverage`

GetCoverage returns the Coverage field if non-nil, zero value otherwise.

### GetCoverageOk

`func (o *PatentPatentProof) GetCoverageOk() (*PatentPatentCoverage, bool)`

GetCoverageOk returns a tuple with the Coverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverage

`func (o *PatentPatentProof) SetCoverage(v PatentPatentCoverage)`

SetCoverage sets Coverage field to given value.

### HasCoverage

`func (o *PatentPatentProof) HasCoverage() bool`

HasCoverage returns a boolean if a field has been set.

### GetGaps

`func (o *PatentPatentProof) GetGaps() []PatentPatentGap`

GetGaps returns the Gaps field if non-nil, zero value otherwise.

### GetGapsOk

`func (o *PatentPatentProof) GetGapsOk() (*[]PatentPatentGap, bool)`

GetGapsOk returns a tuple with the Gaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaps

`func (o *PatentPatentProof) SetGaps(v []PatentPatentGap)`

SetGaps sets Gaps field to given value.

### HasGaps

`func (o *PatentPatentProof) HasGaps() bool`

HasGaps returns a boolean if a field has been set.

### GetOrder

`func (o *PatentPatentProof) GetOrder() []PatentPatentOrderCheck`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PatentPatentProof) GetOrderOk() (*[]PatentPatentOrderCheck, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PatentPatentProof) SetOrder(v []PatentPatentOrderCheck)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PatentPatentProof) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetProject

`func (o *PatentPatentProof) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatentPatentProof) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatentPatentProof) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatentPatentProof) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRecords

`func (o *PatentPatentProof) GetRecords() []PatentPatentProofEntry`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *PatentPatentProof) GetRecordsOk() (*[]PatentPatentProofEntry, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *PatentPatentProof) SetRecords(v []PatentPatentProofEntry)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *PatentPatentProof) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetRoot

`func (o *PatentPatentProof) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *PatentPatentProof) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *PatentPatentProof) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *PatentPatentProof) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetVerdict

`func (o *PatentPatentProof) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *PatentPatentProof) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *PatentPatentProof) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *PatentPatentProof) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### GetVerifiedAt

`func (o *PatentPatentProof) GetVerifiedAt() string`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *PatentPatentProof) GetVerifiedAtOk() (*string, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *PatentPatentProof) SetVerifiedAt(v string)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *PatentPatentProof) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


