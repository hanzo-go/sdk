# SpaceIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the space&#39;s name, matching ^[a-z0-9]([a-z0-9-]{0,38}[a-z0-9])?$ — the shape a drive name takes too, so a caller learns one rule. It is validated AS GIVEN and never lower-cased for you: a client that creates \&quot;Photos\&quot; and then lists \&quot;photos\&quot; would be reading a space it did not make, so mixed case is a clean 400. | [optional] 

## Methods

### NewSpaceIn

`func NewSpaceIn() *SpaceIn`

NewSpaceIn instantiates a new SpaceIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceInWithDefaults

`func NewSpaceInWithDefaults() *SpaceIn`

NewSpaceInWithDefaults instantiates a new SpaceIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SpaceIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SpaceIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SpaceIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SpaceIn) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


