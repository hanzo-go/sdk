# CloudProvidersView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to [**[]CloudProviderCard**](CloudProviderCard.md) | Providers is every cloud this deployment can link, with what each needs. | [optional] 

## Methods

### NewCloudProvidersView

`func NewCloudProvidersView() *CloudProvidersView`

NewCloudProvidersView instantiates a new CloudProvidersView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProvidersViewWithDefaults

`func NewCloudProvidersViewWithDefaults() *CloudProvidersView`

NewCloudProvidersViewWithDefaults instantiates a new CloudProvidersView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *CloudProvidersView) GetProviders() []CloudProviderCard`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *CloudProvidersView) GetProvidersOk() (*[]CloudProviderCard, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *CloudProvidersView) SetProviders(v []CloudProviderCard)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *CloudProvidersView) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


