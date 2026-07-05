# DidLinkIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** |  | 
**ExternalId** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 

## Methods

### NewDidLinkIdentityRequest

`func NewDidLinkIdentityRequest(provider string, externalId string, ) *DidLinkIdentityRequest`

NewDidLinkIdentityRequest instantiates a new DidLinkIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDidLinkIdentityRequestWithDefaults

`func NewDidLinkIdentityRequestWithDefaults() *DidLinkIdentityRequest`

NewDidLinkIdentityRequestWithDefaults instantiates a new DidLinkIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *DidLinkIdentityRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DidLinkIdentityRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DidLinkIdentityRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetExternalId

`func (o *DidLinkIdentityRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *DidLinkIdentityRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *DidLinkIdentityRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetEmail

`func (o *DidLinkIdentityRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DidLinkIdentityRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DidLinkIdentityRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DidLinkIdentityRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


