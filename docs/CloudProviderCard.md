# CloudProviderCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the provider slug used in the path: digitalocean, aws, gcp, azure. | [optional] 
**Keyless** | Pointer to **bool** | Keyless is whether the provider can be linked WITHOUT storing a long-lived secret — AWS by role assumption, GCP by workload identity federation, Azure by federated credential. DigitalOcean is not: it needs a stored token. | [optional] 
**Name** | Pointer to **string** | Name is the provider&#39;s display name. | [optional] 
**Requires** | Pointer to **[]string** | Requires names the credential fields a link body must carry for this provider. | [optional] 

## Methods

### NewCloudProviderCard

`func NewCloudProviderCard() *CloudProviderCard`

NewCloudProviderCard instantiates a new CloudProviderCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderCardWithDefaults

`func NewCloudProviderCardWithDefaults() *CloudProviderCard`

NewCloudProviderCardWithDefaults instantiates a new CloudProviderCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudProviderCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudProviderCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudProviderCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudProviderCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeyless

`func (o *CloudProviderCard) GetKeyless() bool`

GetKeyless returns the Keyless field if non-nil, zero value otherwise.

### GetKeylessOk

`func (o *CloudProviderCard) GetKeylessOk() (*bool, bool)`

GetKeylessOk returns a tuple with the Keyless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyless

`func (o *CloudProviderCard) SetKeyless(v bool)`

SetKeyless sets Keyless field to given value.

### HasKeyless

`func (o *CloudProviderCard) HasKeyless() bool`

HasKeyless returns a boolean if a field has been set.

### GetName

`func (o *CloudProviderCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProviderCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProviderCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudProviderCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequires

`func (o *CloudProviderCard) GetRequires() []string`

GetRequires returns the Requires field if non-nil, zero value otherwise.

### GetRequiresOk

`func (o *CloudProviderCard) GetRequiresOk() (*[]string, bool)`

GetRequiresOk returns a tuple with the Requires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequires

`func (o *CloudProviderCard) SetRequires(v []string)`

SetRequires sets Requires field to given value.

### HasRequires

`func (o *CloudProviderCard) HasRequires() bool`

HasRequires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


