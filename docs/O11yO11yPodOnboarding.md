# O11yO11yPodOnboarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | ClusterName is the pod&#39;s cluster. | [optional] 
**ClusterNamePresent** | Pointer to **bool** | HasClusterName says whether the cluster label is present. | [optional] 
**CronjobNamePresent** | Pointer to **bool** | HasCronjobName says whether the cronjob label is present. | [optional] 
**DaemonsetNamePresent** | Pointer to **bool** | HasDaemonsetName says whether the daemonset label is present. | [optional] 
**DeploymentNamePresent** | Pointer to **bool** | HasDeploymentName says whether the deployment label is present. | [optional] 
**JobNamePresent** | Pointer to **bool** | HasJobName says whether the job label is present. | [optional] 
**NamespaceNamePresent** | Pointer to **bool** | HasNamespaceName says whether the namespace label is present. | [optional] 
**NodeNamePresent** | Pointer to **bool** | HasNodeName says whether the node label is present. | [optional] 
**StatefulsetNamePresent** | Pointer to **bool** | HasStatefulsetName says whether the statefulset label is present. | [optional] 
**NamespaceName** | Pointer to **string** | NamespaceName is the pod&#39;s namespace. | [optional] 
**NodeName** | Pointer to **string** | NodeName is the pod&#39;s node. | [optional] 
**PodName** | Pointer to **string** | PodName is the pod. | [optional] 

## Methods

### NewO11yO11yPodOnboarding

`func NewO11yO11yPodOnboarding() *O11yO11yPodOnboarding`

NewO11yO11yPodOnboarding instantiates a new O11yO11yPodOnboarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yPodOnboardingWithDefaults

`func NewO11yO11yPodOnboardingWithDefaults() *O11yO11yPodOnboarding`

NewO11yO11yPodOnboardingWithDefaults instantiates a new O11yO11yPodOnboarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *O11yO11yPodOnboarding) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *O11yO11yPodOnboarding) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *O11yO11yPodOnboarding) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *O11yO11yPodOnboarding) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetClusterNamePresent

`func (o *O11yO11yPodOnboarding) GetClusterNamePresent() bool`

GetClusterNamePresent returns the ClusterNamePresent field if non-nil, zero value otherwise.

### GetClusterNamePresentOk

`func (o *O11yO11yPodOnboarding) GetClusterNamePresentOk() (*bool, bool)`

GetClusterNamePresentOk returns a tuple with the ClusterNamePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNamePresent

`func (o *O11yO11yPodOnboarding) SetClusterNamePresent(v bool)`

SetClusterNamePresent sets ClusterNamePresent field to given value.

### HasClusterNamePresent

`func (o *O11yO11yPodOnboarding) HasClusterNamePresent() bool`

HasClusterNamePresent returns a boolean if a field has been set.

### GetCronjobNamePresent

`func (o *O11yO11yPodOnboarding) GetCronjobNamePresent() bool`

GetCronjobNamePresent returns the CronjobNamePresent field if non-nil, zero value otherwise.

### GetCronjobNamePresentOk

`func (o *O11yO11yPodOnboarding) GetCronjobNamePresentOk() (*bool, bool)`

GetCronjobNamePresentOk returns a tuple with the CronjobNamePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronjobNamePresent

`func (o *O11yO11yPodOnboarding) SetCronjobNamePresent(v bool)`

SetCronjobNamePresent sets CronjobNamePresent field to given value.

### HasCronjobNamePresent

`func (o *O11yO11yPodOnboarding) HasCronjobNamePresent() bool`

HasCronjobNamePresent returns a boolean if a field has been set.

### GetDaemonsetNamePresent

`func (o *O11yO11yPodOnboarding) GetDaemonsetNamePresent() bool`

GetDaemonsetNamePresent returns the DaemonsetNamePresent field if non-nil, zero value otherwise.

### GetDaemonsetNamePresentOk

`func (o *O11yO11yPodOnboarding) GetDaemonsetNamePresentOk() (*bool, bool)`

GetDaemonsetNamePresentOk returns a tuple with the DaemonsetNamePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonsetNamePresent

`func (o *O11yO11yPodOnboarding) SetDaemonsetNamePresent(v bool)`

SetDaemonsetNamePresent sets DaemonsetNamePresent field to given value.

### HasDaemonsetNamePresent

`func (o *O11yO11yPodOnboarding) HasDaemonsetNamePresent() bool`

HasDaemonsetNamePresent returns a boolean if a field has been set.

### GetDeploymentNamePresent

`func (o *O11yO11yPodOnboarding) GetDeploymentNamePresent() bool`

GetDeploymentNamePresent returns the DeploymentNamePresent field if non-nil, zero value otherwise.

### GetDeploymentNamePresentOk

`func (o *O11yO11yPodOnboarding) GetDeploymentNamePresentOk() (*bool, bool)`

GetDeploymentNamePresentOk returns a tuple with the DeploymentNamePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentNamePresent

`func (o *O11yO11yPodOnboarding) SetDeploymentNamePresent(v bool)`

SetDeploymentNamePresent sets DeploymentNamePresent field to given value.

### HasDeploymentNamePresent

`func (o *O11yO11yPodOnboarding) HasDeploymentNamePresent() bool`

HasDeploymentNamePresent returns a boolean if a field has been set.

### GetJobNamePresent

`func (o *O11yO11yPodOnboarding) GetJobNamePresent() bool`

GetJobNamePresent returns the JobNamePresent field if non-nil, zero value otherwise.

### GetJobNamePresentOk

`func (o *O11yO11yPodOnboarding) GetJobNamePresentOk() (*bool, bool)`

GetJobNamePresentOk returns a tuple with the JobNamePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobNamePresent

`func (o *O11yO11yPodOnboarding) SetJobNamePresent(v bool)`

SetJobNamePresent sets JobNamePresent field to given value.

### HasJobNamePresent

`func (o *O11yO11yPodOnboarding) HasJobNamePresent() bool`

HasJobNamePresent returns a boolean if a field has been set.

### GetNamespaceNamePresent

`func (o *O11yO11yPodOnboarding) GetNamespaceNamePresent() bool`

GetNamespaceNamePresent returns the NamespaceNamePresent field if non-nil, zero value otherwise.

### GetNamespaceNamePresentOk

`func (o *O11yO11yPodOnboarding) GetNamespaceNamePresentOk() (*bool, bool)`

GetNamespaceNamePresentOk returns a tuple with the NamespaceNamePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceNamePresent

`func (o *O11yO11yPodOnboarding) SetNamespaceNamePresent(v bool)`

SetNamespaceNamePresent sets NamespaceNamePresent field to given value.

### HasNamespaceNamePresent

`func (o *O11yO11yPodOnboarding) HasNamespaceNamePresent() bool`

HasNamespaceNamePresent returns a boolean if a field has been set.

### GetNodeNamePresent

`func (o *O11yO11yPodOnboarding) GetNodeNamePresent() bool`

GetNodeNamePresent returns the NodeNamePresent field if non-nil, zero value otherwise.

### GetNodeNamePresentOk

`func (o *O11yO11yPodOnboarding) GetNodeNamePresentOk() (*bool, bool)`

GetNodeNamePresentOk returns a tuple with the NodeNamePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNamePresent

`func (o *O11yO11yPodOnboarding) SetNodeNamePresent(v bool)`

SetNodeNamePresent sets NodeNamePresent field to given value.

### HasNodeNamePresent

`func (o *O11yO11yPodOnboarding) HasNodeNamePresent() bool`

HasNodeNamePresent returns a boolean if a field has been set.

### GetStatefulsetNamePresent

`func (o *O11yO11yPodOnboarding) GetStatefulsetNamePresent() bool`

GetStatefulsetNamePresent returns the StatefulsetNamePresent field if non-nil, zero value otherwise.

### GetStatefulsetNamePresentOk

`func (o *O11yO11yPodOnboarding) GetStatefulsetNamePresentOk() (*bool, bool)`

GetStatefulsetNamePresentOk returns a tuple with the StatefulsetNamePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulsetNamePresent

`func (o *O11yO11yPodOnboarding) SetStatefulsetNamePresent(v bool)`

SetStatefulsetNamePresent sets StatefulsetNamePresent field to given value.

### HasStatefulsetNamePresent

`func (o *O11yO11yPodOnboarding) HasStatefulsetNamePresent() bool`

HasStatefulsetNamePresent returns a boolean if a field has been set.

### GetNamespaceName

`func (o *O11yO11yPodOnboarding) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *O11yO11yPodOnboarding) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *O11yO11yPodOnboarding) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *O11yO11yPodOnboarding) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### GetNodeName

`func (o *O11yO11yPodOnboarding) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *O11yO11yPodOnboarding) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *O11yO11yPodOnboarding) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *O11yO11yPodOnboarding) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPodName

`func (o *O11yO11yPodOnboarding) GetPodName() string`

GetPodName returns the PodName field if non-nil, zero value otherwise.

### GetPodNameOk

`func (o *O11yO11yPodOnboarding) GetPodNameOk() (*string, bool)`

GetPodNameOk returns a tuple with the PodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodName

`func (o *O11yO11yPodOnboarding) SetPodName(v string)`

SetPodName sets PodName field to given value.

### HasPodName

`func (o *O11yO11yPodOnboarding) HasPodName() bool`

HasPodName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


