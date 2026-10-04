# WorldWorldWire

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to **string** | Auth states what the wire asks of the caller, including which parts of it answer without a token. | [optional] 
**Name** | Pointer to **string** | Name is the wire&#39;s short id — rest, mcp or zap. | [optional] 
**Path** | Pointer to **string** | Path is the address the wire answers on, under this same origin. | [optional] 
**Protocol** | Pointer to **string** | Protocol names what the wire speaks, so a caller knows which client to point at it. | [optional] 
**Spec** | Pointer to **string** | Spec is where this wire&#39;s operations are enumerated, when they are enumerated in a document at all. Empty for a wire that describes itself over its own protocol. | [optional] 

## Methods

### NewWorldWorldWire

`func NewWorldWorldWire() *WorldWorldWire`

NewWorldWorldWire instantiates a new WorldWorldWire object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorldWorldWireWithDefaults

`func NewWorldWorldWireWithDefaults() *WorldWorldWire`

NewWorldWorldWireWithDefaults instantiates a new WorldWorldWire object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *WorldWorldWire) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *WorldWorldWire) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *WorldWorldWire) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *WorldWorldWire) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetName

`func (o *WorldWorldWire) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorldWorldWire) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorldWorldWire) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorldWorldWire) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *WorldWorldWire) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WorldWorldWire) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WorldWorldWire) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WorldWorldWire) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProtocol

`func (o *WorldWorldWire) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *WorldWorldWire) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *WorldWorldWire) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *WorldWorldWire) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSpec

`func (o *WorldWorldWire) GetSpec() string`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *WorldWorldWire) GetSpecOk() (*string, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *WorldWorldWire) SetSpec(v string)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *WorldWorldWire) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


