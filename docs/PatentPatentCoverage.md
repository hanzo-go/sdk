# PatentPatentCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | Pointer to **int64** | Approvals is how many decisions it covered. | [optional] 
**Disclosures** | Pointer to **int64** | Disclosures is how many documents it covered. | [optional] 
**Ideas** | Pointer to **int64** | Ideas is how many conception entries the fold covered. | [optional] 
**Implementations** | Pointer to **int64** | Implementations is how many implementation records it covered. | [optional] 
**Versions** | Pointer to **int64** | Versions is how many RETAINED versions were folded, which is at least the sum above and is the number that says corrections happened. | [optional] 

## Methods

### NewPatentPatentCoverage

`func NewPatentPatentCoverage() *PatentPatentCoverage`

NewPatentPatentCoverage instantiates a new PatentPatentCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentCoverageWithDefaults

`func NewPatentPatentCoverageWithDefaults() *PatentPatentCoverage`

NewPatentPatentCoverageWithDefaults instantiates a new PatentPatentCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *PatentPatentCoverage) GetApprovals() int64`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *PatentPatentCoverage) GetApprovalsOk() (*int64, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *PatentPatentCoverage) SetApprovals(v int64)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *PatentPatentCoverage) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetDisclosures

`func (o *PatentPatentCoverage) GetDisclosures() int64`

GetDisclosures returns the Disclosures field if non-nil, zero value otherwise.

### GetDisclosuresOk

`func (o *PatentPatentCoverage) GetDisclosuresOk() (*int64, bool)`

GetDisclosuresOk returns a tuple with the Disclosures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosures

`func (o *PatentPatentCoverage) SetDisclosures(v int64)`

SetDisclosures sets Disclosures field to given value.

### HasDisclosures

`func (o *PatentPatentCoverage) HasDisclosures() bool`

HasDisclosures returns a boolean if a field has been set.

### GetIdeas

`func (o *PatentPatentCoverage) GetIdeas() int64`

GetIdeas returns the Ideas field if non-nil, zero value otherwise.

### GetIdeasOk

`func (o *PatentPatentCoverage) GetIdeasOk() (*int64, bool)`

GetIdeasOk returns a tuple with the Ideas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeas

`func (o *PatentPatentCoverage) SetIdeas(v int64)`

SetIdeas sets Ideas field to given value.

### HasIdeas

`func (o *PatentPatentCoverage) HasIdeas() bool`

HasIdeas returns a boolean if a field has been set.

### GetImplementations

`func (o *PatentPatentCoverage) GetImplementations() int64`

GetImplementations returns the Implementations field if non-nil, zero value otherwise.

### GetImplementationsOk

`func (o *PatentPatentCoverage) GetImplementationsOk() (*int64, bool)`

GetImplementationsOk returns a tuple with the Implementations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementations

`func (o *PatentPatentCoverage) SetImplementations(v int64)`

SetImplementations sets Implementations field to given value.

### HasImplementations

`func (o *PatentPatentCoverage) HasImplementations() bool`

HasImplementations returns a boolean if a field has been set.

### GetVersions

`func (o *PatentPatentCoverage) GetVersions() int64`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PatentPatentCoverage) GetVersionsOk() (*int64, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PatentPatentCoverage) SetVersions(v int64)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PatentPatentCoverage) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


