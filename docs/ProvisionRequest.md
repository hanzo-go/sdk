# ProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to **string** | Instance binds a DEDICATED add-on to the app instance whose &lt;instance&gt;-addons Secret receives the &lt;KIND&gt;_URL (e.g. \&quot;commerce\&quot;). Optional: empty means \&quot;not instance-bound\&quot; — the DSN is returned once and wired by the caller. | [optional] 
**Name** | Pointer to **string** | Name is the org-unique slug for the new resource, matching ^[a-z0-9]([a-z0-9-]{0,38}[a-z0-9])?$. Every physical name derives from it. | [optional] 

## Methods

### NewProvisionRequest

`func NewProvisionRequest() *ProvisionRequest`

NewProvisionRequest instantiates a new ProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionRequestWithDefaults

`func NewProvisionRequestWithDefaults() *ProvisionRequest`

NewProvisionRequestWithDefaults instantiates a new ProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *ProvisionRequest) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ProvisionRequest) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ProvisionRequest) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ProvisionRequest) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetName

`func (o *ProvisionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProvisionRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


