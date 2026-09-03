# O11yStatusResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckedAt** | Pointer to **string** | CheckedAt is when this answer was measured, RFC3339 UTC. | [optional] 
**Deployments** | Pointer to [**[]O11yDeployment**](O11yDeployment.md) | Deployments is the per-replica inventory behind the verdict. Empty means the telemetry store reported none, not that the service runs on none. | [optional] 
**LatencyMs** | Pointer to **int64** | LatencyMs is the health probe&#39;s round trip in MILLISECONDS, time-boxed at two seconds. It is 0 when no probe answered, which is not a fast service. | [optional] 
**Product** | Pointer to **string** | Product is the service this answer is about, echoed back. | [optional] 
**Source** | Pointer to **string** | Source is where the verdict came from: \&quot;probe\&quot; (we asked and it answered), \&quot;datastore\&quot; (the probe did not answer and the replica inventory decided it), \&quot;unreachable\&quot; (neither), or \&quot;unknown-service\&quot; for a well-formed product name nothing backs — which is answered without probing, since dialling an arbitrary host on a caller&#39;s say-so is the request forgery this refuses. | [optional] 
**Up** | Pointer to **bool** | Up is true when the health probe succeeded OR any replica reports up, so a service reachable by either route reads up. Read Source to know which. | [optional] 

## Methods

### NewO11yStatusResult

`func NewO11yStatusResult() *O11yStatusResult`

NewO11yStatusResult instantiates a new O11yStatusResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yStatusResultWithDefaults

`func NewO11yStatusResultWithDefaults() *O11yStatusResult`

NewO11yStatusResultWithDefaults instantiates a new O11yStatusResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckedAt

`func (o *O11yStatusResult) GetCheckedAt() string`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *O11yStatusResult) GetCheckedAtOk() (*string, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *O11yStatusResult) SetCheckedAt(v string)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *O11yStatusResult) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### GetDeployments

`func (o *O11yStatusResult) GetDeployments() []O11yDeployment`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *O11yStatusResult) GetDeploymentsOk() (*[]O11yDeployment, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *O11yStatusResult) SetDeployments(v []O11yDeployment)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *O11yStatusResult) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetLatencyMs

`func (o *O11yStatusResult) GetLatencyMs() int64`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *O11yStatusResult) GetLatencyMsOk() (*int64, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *O11yStatusResult) SetLatencyMs(v int64)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *O11yStatusResult) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetProduct

`func (o *O11yStatusResult) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *O11yStatusResult) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *O11yStatusResult) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *O11yStatusResult) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSource

`func (o *O11yStatusResult) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *O11yStatusResult) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *O11yStatusResult) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *O11yStatusResult) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUp

`func (o *O11yStatusResult) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *O11yStatusResult) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *O11yStatusResult) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *O11yStatusResult) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


