# FlagPlatformOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flags** | Pointer to **map[string]bool** | Flags maps each principal switch&#39;s key to whether it is on for the caller: &#x60;beta&#x60; is whether the caller sees products and catalog items that are not yet generally available, and beta.&lt;id&gt; whether they see that one app in beta in the Hanzo app. | [optional] 

## Methods

### NewFlagPlatformOut

`func NewFlagPlatformOut() *FlagPlatformOut`

NewFlagPlatformOut instantiates a new FlagPlatformOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlagPlatformOutWithDefaults

`func NewFlagPlatformOutWithDefaults() *FlagPlatformOut`

NewFlagPlatformOutWithDefaults instantiates a new FlagPlatformOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlags

`func (o *FlagPlatformOut) GetFlags() map[string]bool`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *FlagPlatformOut) GetFlagsOk() (*map[string]bool, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *FlagPlatformOut) SetFlags(v map[string]bool)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *FlagPlatformOut) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


