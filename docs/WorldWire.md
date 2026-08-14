# WorldWire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to **string** | Auth states what the wire asks of the caller, including which parts of it answer without a token. | [optional] 
**Name** | Pointer to **string** | Name is the wire&#39;s short id — rest, mcp or zap. | [optional] 
**Path** | Pointer to **string** | Path is the address the wire answers on, under this same origin. | [optional] 
**Protocol** | Pointer to **string** | Protocol names what the wire speaks, so a caller knows which client to point at it. | [optional] 
**Spec** | Pointer to **string** | Spec is where this wire&#39;s operations are enumerated, when they are enumerated in a document at all. Empty for a wire that describes itself over its own protocol. | [optional] 

## Methods

### NewWorldWire

`func NewWorldWire() *WorldWire`

NewWorldWire instantiates a new WorldWire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldWireWithDefaults

`func NewWorldWireWithDefaults() *WorldWire`

NewWorldWireWithDefaults instantiates a new WorldWire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *WorldWire) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *WorldWire) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *WorldWire) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *WorldWire) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetName

`func (o *WorldWire) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorldWire) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorldWire) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorldWire) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *WorldWire) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WorldWire) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WorldWire) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WorldWire) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProtocol

`func (o *WorldWire) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *WorldWire) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *WorldWire) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *WorldWire) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSpec

`func (o *WorldWire) GetSpec() string`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *WorldWire) GetSpecOk() (*string, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *WorldWire) SetSpec(v string)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *WorldWire) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


