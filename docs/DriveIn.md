# DriveIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the drive&#39;s name, matching the same shape a space name does. It becomes the FIRST SEGMENT of every key the drive holds, which is why it may carry no \&quot;/\&quot;. | [optional] 
**Space** | Pointer to **string** | Space is the space to create the drive in, from the path. It carries NO &#x60;url:\&quot;-\&quot;&#x60;, unlike the field below it, and the difference is the whole reason both tags are written out: zip&#39;s binder skips a field tagged \&quot;-\&quot; for EVERY URL source, path params included, so a path-borne value that carried it would arrive empty and the create would refuse a perfectly good address. | [optional] 

## Methods

### NewDriveIn

`func NewDriveIn() *DriveIn`

NewDriveIn instantiates a new DriveIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveInWithDefaults

`func NewDriveInWithDefaults() *DriveIn`

NewDriveInWithDefaults instantiates a new DriveIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DriveIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DriveIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DriveIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DriveIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpace

`func (o *DriveIn) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *DriveIn) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *DriveIn) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *DriveIn) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


