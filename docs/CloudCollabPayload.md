# CloudCollabPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **map[string]string** | Content maps a document field to its ProseMirror markup JSON. | [optional] 
**Source** | Pointer to **string** | Source is the blob ref a getContent reads the snapshot from. Absent means there is no snapshot to read, which answers empty content. | [optional] 
**Updates** | Pointer to **map[string]string** | Updates carries, per field, a base64 Y.js state update encoding the SAME markup — the front computes it (markupToYDoc → encodeStateAsUpdate) so a createContent seeds the live-editing lane&#39;s update log, not just the snapshot blob. Without it a dialog-created description is invisible in the collaborative editor, which replays the ydoc log, never the snapshot. | [optional] 

## Methods

### NewCloudCollabPayload

`func NewCloudCollabPayload() *CloudCollabPayload`

NewCloudCollabPayload instantiates a new CloudCollabPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCollabPayloadWithDefaults

`func NewCloudCollabPayloadWithDefaults() *CloudCollabPayload`

NewCloudCollabPayloadWithDefaults instantiates a new CloudCollabPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CloudCollabPayload) GetContent() map[string]string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CloudCollabPayload) GetContentOk() (*map[string]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CloudCollabPayload) SetContent(v map[string]string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CloudCollabPayload) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetSource

`func (o *CloudCollabPayload) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudCollabPayload) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudCollabPayload) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudCollabPayload) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUpdates

`func (o *CloudCollabPayload) GetUpdates() map[string]string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *CloudCollabPayload) GetUpdatesOk() (*map[string]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *CloudCollabPayload) SetUpdates(v map[string]string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *CloudCollabPayload) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


