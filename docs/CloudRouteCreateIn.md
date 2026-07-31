# CloudRouteCreateIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | Pointer to **string** | Pattern is the URL pattern to bind, e.g. \&quot;acme.com/api/_*\&quot;. | [optional] 
**Script** | Pointer to **string** | Script is the Worker script to dispatch to. Omit it to leave the pattern bound to no script, which is how Cloudflare expresses \&quot;bypass the Worker here\&quot;. | [optional] 
**Zone** | Pointer to **string** | Zone is the 32-hex Cloudflare zone id, from the path. | [optional] 

## Methods

### NewCloudRouteCreateIn

`func NewCloudRouteCreateIn() *CloudRouteCreateIn`

NewCloudRouteCreateIn instantiates a new CloudRouteCreateIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouteCreateInWithDefaults

`func NewCloudRouteCreateInWithDefaults() *CloudRouteCreateIn`

NewCloudRouteCreateInWithDefaults instantiates a new CloudRouteCreateIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *CloudRouteCreateIn) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CloudRouteCreateIn) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CloudRouteCreateIn) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CloudRouteCreateIn) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetScript

`func (o *CloudRouteCreateIn) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *CloudRouteCreateIn) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *CloudRouteCreateIn) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *CloudRouteCreateIn) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetZone

`func (o *CloudRouteCreateIn) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CloudRouteCreateIn) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CloudRouteCreateIn) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CloudRouteCreateIn) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


