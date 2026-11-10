# ComplianceAct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the operator&#39;s IAM username. | [optional] 
**Owner** | Pointer to **string** | Owner is the org of the operator&#39;s IAM row: \&quot;admin\&quot; for a SuperAdmin. | [optional] 
**Sub** | Pointer to **string** | Sub is the operator&#39;s IAM subject. | [optional] 

## Methods

### NewComplianceAct

`func NewComplianceAct() *ComplianceAct`

NewComplianceAct instantiates a new ComplianceAct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceActWithDefaults

`func NewComplianceActWithDefaults() *ComplianceAct`

NewComplianceActWithDefaults instantiates a new ComplianceAct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ComplianceAct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComplianceAct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComplianceAct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComplianceAct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *ComplianceAct) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ComplianceAct) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ComplianceAct) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ComplianceAct) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSub

`func (o *ComplianceAct) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *ComplianceAct) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *ComplianceAct) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *ComplianceAct) HasSub() bool`

HasSub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


