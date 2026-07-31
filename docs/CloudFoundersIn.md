# CloudFoundersIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Founders** | Pointer to [**[]CloudFounder**](CloudFounder.md) | Founders is every founding stakeholder. Each needs a name and an email, and equityBps between 0 and 10000 (1% &#x3D;&#x3D; 100 bps). | [optional] 

## Methods

### NewCloudFoundersIn

`func NewCloudFoundersIn() *CloudFoundersIn`

NewCloudFoundersIn instantiates a new CloudFoundersIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFoundersInWithDefaults

`func NewCloudFoundersInWithDefaults() *CloudFoundersIn`

NewCloudFoundersInWithDefaults instantiates a new CloudFoundersIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFounders

`func (o *CloudFoundersIn) GetFounders() []CloudFounder`

GetFounders returns the Founders field if non-nil, zero value otherwise.

### GetFoundersOk

`func (o *CloudFoundersIn) GetFoundersOk() (*[]CloudFounder, bool)`

GetFoundersOk returns a tuple with the Founders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFounders

`func (o *CloudFoundersIn) SetFounders(v []CloudFounder)`

SetFounders sets Founders field to given value.

### HasFounders

`func (o *CloudFoundersIn) HasFounders() bool`

HasFounders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


