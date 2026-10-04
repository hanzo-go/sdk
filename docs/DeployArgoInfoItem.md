# DeployArgoInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the chip&#39;s label. The only one this projection produces is \&quot;Image Tag\&quot;. | [optional] 
**Value** | Pointer to **string** | Value is the chip&#39;s value — for \&quot;Image Tag\&quot;, the tag the node runs. | [optional] 

## Methods

### NewDeployArgoInfoItem

`func NewDeployArgoInfoItem() *DeployArgoInfoItem`

NewDeployArgoInfoItem instantiates a new DeployArgoInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoInfoItemWithDefaults

`func NewDeployArgoInfoItemWithDefaults() *DeployArgoInfoItem`

NewDeployArgoInfoItemWithDefaults instantiates a new DeployArgoInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeployArgoInfoItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeployArgoInfoItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeployArgoInfoItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeployArgoInfoItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *DeployArgoInfoItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeployArgoInfoItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeployArgoInfoItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeployArgoInfoItem) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


