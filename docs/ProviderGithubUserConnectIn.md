# ProviderGithubUserConnectIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Return** | Pointer to **string** | Return is the page the callback sends the person back to, with complete&#x3D;github&amp;grant&#x3D;&lt;id&gt; (or error&#x3D;github&amp;reason&#x3D;&lt;why&gt;) added to its query. It must name one of this deployment&#39;s return pages (on api.hanzo.ai: https://hanzo.ai/?at&#x3D;-/settings/integrations, https://console.hanzo.ai/connectors, https://platform.hanzo.ai/connectors); anything else, or nothing, returns to the console&#39;s /connectors. | [optional] 

## Methods

### NewProviderGithubUserConnectIn

`func NewProviderGithubUserConnectIn() *ProviderGithubUserConnectIn`

NewProviderGithubUserConnectIn instantiates a new ProviderGithubUserConnectIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubUserConnectInWithDefaults

`func NewProviderGithubUserConnectInWithDefaults() *ProviderGithubUserConnectIn`

NewProviderGithubUserConnectInWithDefaults instantiates a new ProviderGithubUserConnectIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturn

`func (o *ProviderGithubUserConnectIn) GetReturn() string`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *ProviderGithubUserConnectIn) GetReturnOk() (*string, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturn

`func (o *ProviderGithubUserConnectIn) SetReturn(v string)`

SetReturn sets Return field to given value.

### HasReturn

`func (o *ProviderGithubUserConnectIn) HasReturn() bool`

HasReturn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


