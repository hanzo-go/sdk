# O11yDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to **string** | Instance is the replica as the telemetry store labels it — the address the series was recorded against, which is what distinguishes two replicas of one service. | [optional] 
**Up** | Pointer to **bool** | Up is that replica&#39;s last reported state. Every target emits on every cycle, so a replica missing from the list is one the prober is not reporting at all, which is a different fact from down. | [optional] 

## Methods

### NewO11yDeployment

`func NewO11yDeployment() *O11yDeployment`

NewO11yDeployment instantiates a new O11yDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDeploymentWithDefaults

`func NewO11yDeploymentWithDefaults() *O11yDeployment`

NewO11yDeploymentWithDefaults instantiates a new O11yDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *O11yDeployment) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *O11yDeployment) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *O11yDeployment) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *O11yDeployment) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetUp

`func (o *O11yDeployment) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *O11yDeployment) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *O11yDeployment) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *O11yDeployment) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


