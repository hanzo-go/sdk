# ProvisioningProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to **string** | Instance binds a DEDICATED add-on to the app instance whose &lt;instance&gt;-addons Secret receives the &lt;KIND&gt;_URL (e.g. \&quot;commerce\&quot;). Optional: empty means \&quot;not instance-bound\&quot; — the DSN is returned once and wired by the caller. | [optional] 
**Name** | Pointer to **string** | Name is the org-unique slug for the new resource, matching ^[a-z0-9]([a-z0-9-]{0,38}[a-z0-9])?$. Every physical name derives from it. | [optional] 

## Methods

### NewProvisioningProvisionRequest

`func NewProvisioningProvisionRequest() *ProvisioningProvisionRequest`

NewProvisioningProvisionRequest instantiates a new ProvisioningProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningProvisionRequestWithDefaults

`func NewProvisioningProvisionRequestWithDefaults() *ProvisioningProvisionRequest`

NewProvisioningProvisionRequestWithDefaults instantiates a new ProvisioningProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *ProvisioningProvisionRequest) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ProvisioningProvisionRequest) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ProvisioningProvisionRequest) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ProvisioningProvisionRequest) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *ProvisioningProvisionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningProvisionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningProvisionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisioningProvisionRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


