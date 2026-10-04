# DeployArgoSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**DeployArgoDestination**](DeployArgoDestination.md) | Destination is which cluster and namespace it lands in. Zero-valued on a CD row: this projection reports CD&#39;s source, not its destination. | [optional] 
**Project** | Pointer to **string** | Project is the AppProject this application is grouped and filtered under. For an App CR it is the app.kubernetes.io/part-of label — the IAM project name — falling back to \&quot;default\&quot; when the CR carries no such label. | [optional] 
**Source** | Pointer to [**DeployArgoSource**](DeployArgoSource.md) | Source is where the desired state is declared. | [optional] 

## Methods

### NewDeployArgoSpec

`func NewDeployArgoSpec() *DeployArgoSpec`

NewDeployArgoSpec instantiates a new DeployArgoSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoSpecWithDefaults

`func NewDeployArgoSpecWithDefaults() *DeployArgoSpec`

NewDeployArgoSpecWithDefaults instantiates a new DeployArgoSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *DeployArgoSpec) GetDestination() DeployArgoDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *DeployArgoSpec) GetDestinationOk() (*DeployArgoDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *DeployArgoSpec) SetDestination(v DeployArgoDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *DeployArgoSpec) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetProject

`func (o *DeployArgoSpec) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *DeployArgoSpec) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *DeployArgoSpec) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *DeployArgoSpec) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSource

`func (o *DeployArgoSpec) GetSource() DeployArgoSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DeployArgoSpec) GetSourceOk() (*DeployArgoSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DeployArgoSpec) SetSource(v DeployArgoSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *DeployArgoSpec) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


