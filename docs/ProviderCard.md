# ProviderCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the provider slug used in the path: digitalocean, aws, gcp, azure. | [optional] 
**Keyless** | Pointer to **bool** | Keyless is whether the provider can be linked WITHOUT storing a long-lived secret — AWS by role assumption, GCP by workload identity federation, Azure by federated credential. DigitalOcean is not: it needs a stored token. | [optional] 
**Name** | Pointer to **string** | Name is the provider&#39;s display name. | [optional] 
**Requires** | Pointer to **[]string** | Requires names the credential fields a link body must carry for this provider. | [optional] 

## Methods

### NewProviderCard

`func NewProviderCard() *ProviderCard`

NewProviderCard instantiates a new ProviderCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderCardWithDefaults

`func NewProviderCardWithDefaults() *ProviderCard`

NewProviderCardWithDefaults instantiates a new ProviderCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderCard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderCard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderCard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderCard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeyless

`func (o *ProviderCard) GetKeyless() bool`

GetKeyless returns the Keyless field if non-nil, zero value otherwise.

### GetKeylessOk

`func (o *ProviderCard) GetKeylessOk() (*bool, bool)`

GetKeylessOk returns a tuple with the Keyless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyless

`func (o *ProviderCard) SetKeyless(v bool)`

SetKeyless sets Keyless field to given value.

### HasKeyless

`func (o *ProviderCard) HasKeyless() bool`

HasKeyless returns a boolean if a field has been set.

### GetName

`func (o *ProviderCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequires

`func (o *ProviderCard) GetRequires() []string`

GetRequires returns the Requires field if non-nil, zero value otherwise.

### GetRequiresOk

`func (o *ProviderCard) GetRequiresOk() (*[]string, bool)`

GetRequiresOk returns a tuple with the Requires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequires

`func (o *ProviderCard) SetRequires(v []string)`

SetRequires sets Requires field to given value.

### HasRequires

`func (o *ProviderCard) HasRequires() bool`

HasRequires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


