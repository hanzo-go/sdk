# DidLinkedIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**LinkedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDidLinkedIdentity

`func NewDidLinkedIdentity() *DidLinkedIdentity`

NewDidLinkedIdentity instantiates a new DidLinkedIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDidLinkedIdentityWithDefaults

`func NewDidLinkedIdentityWithDefaults() *DidLinkedIdentity`

NewDidLinkedIdentityWithDefaults instantiates a new DidLinkedIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *DidLinkedIdentity) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DidLinkedIdentity) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DidLinkedIdentity) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DidLinkedIdentity) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetExternalId

`func (o *DidLinkedIdentity) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *DidLinkedIdentity) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *DidLinkedIdentity) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *DidLinkedIdentity) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetEmail

`func (o *DidLinkedIdentity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DidLinkedIdentity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DidLinkedIdentity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DidLinkedIdentity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLinkedAt

`func (o *DidLinkedIdentity) GetLinkedAt() time.Time`

GetLinkedAt returns the LinkedAt field if non-nil, zero value otherwise.

### GetLinkedAtOk

`func (o *DidLinkedIdentity) GetLinkedAtOk() (*time.Time, bool)`

GetLinkedAtOk returns a tuple with the LinkedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAt

`func (o *DidLinkedIdentity) SetLinkedAt(v time.Time)`

SetLinkedAt sets LinkedAt field to given value.

### HasLinkedAt

`func (o *DidLinkedIdentity) HasLinkedAt() bool`

HasLinkedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


