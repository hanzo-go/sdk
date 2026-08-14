# ProviderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | DisplayName is the human label for the sign-in button; this deployment sends \&quot;Hanzo\&quot;. Omitted from the body when empty. | [optional] 
**Name** | Pointer to **string** | Name is the provider id, and it is the value that goes back in the URL to start a login: GET /v1/team/account/auth/{provider}. This deployment surfaces exactly one, \&quot;openid\&quot; — the hanzo.id door. | [optional] 

## Methods

### NewProviderInfo

`func NewProviderInfo() *ProviderInfo`

NewProviderInfo instantiates a new ProviderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderInfoWithDefaults

`func NewProviderInfoWithDefaults() *ProviderInfo`

NewProviderInfoWithDefaults instantiates a new ProviderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ProviderInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProviderInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProviderInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ProviderInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *ProviderInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProviderInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


