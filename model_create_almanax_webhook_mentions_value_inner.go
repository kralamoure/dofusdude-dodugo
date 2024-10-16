/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [PHP](https://github.com/dofusdude/dofusdude-php) - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 0.9.3
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
)

// checks if the CreateAlmanaxWebhookMentionsValueInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAlmanaxWebhookMentionsValueInner{}

// CreateAlmanaxWebhookMentionsValueInner Mention
type CreateAlmanaxWebhookMentionsValueInner struct {
	// User or role ID directly from Discord. Activate developer mode and right click a user or role to get them.
	DiscordId *int32 `json:"discord_id,omitempty"`
	// Whether an ID describes a role (true) or user (false). This is needed for formatting the mention command so Discord understands it.
	IsRole *bool `json:"is_role,omitempty"`
	// Get a mention days before the bonus comes up. It will post on the specified time. Also works when the interval is not daily.
	PingDaysBefore NullableInt32 `json:"ping_days_before,omitempty"`
}

// NewCreateAlmanaxWebhookMentionsValueInner instantiates a new CreateAlmanaxWebhookMentionsValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAlmanaxWebhookMentionsValueInner() *CreateAlmanaxWebhookMentionsValueInner {
	this := CreateAlmanaxWebhookMentionsValueInner{}
	return &this
}

// NewCreateAlmanaxWebhookMentionsValueInnerWithDefaults instantiates a new CreateAlmanaxWebhookMentionsValueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAlmanaxWebhookMentionsValueInnerWithDefaults() *CreateAlmanaxWebhookMentionsValueInner {
	this := CreateAlmanaxWebhookMentionsValueInner{}
	return &this
}

// GetDiscordId returns the DiscordId field value if set, zero value otherwise.
func (o *CreateAlmanaxWebhookMentionsValueInner) GetDiscordId() int32 {
	if o == nil || IsNil(o.DiscordId) {
		var ret int32
		return ret
	}
	return *o.DiscordId
}

// GetDiscordIdOk returns a tuple with the DiscordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlmanaxWebhookMentionsValueInner) GetDiscordIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscordId) {
		return nil, false
	}
	return o.DiscordId, true
}

// HasDiscordId returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhookMentionsValueInner) HasDiscordId() bool {
	if o != nil && !IsNil(o.DiscordId) {
		return true
	}

	return false
}

// SetDiscordId gets a reference to the given int32 and assigns it to the DiscordId field.
func (o *CreateAlmanaxWebhookMentionsValueInner) SetDiscordId(v int32) {
	o.DiscordId = &v
}

// GetIsRole returns the IsRole field value if set, zero value otherwise.
func (o *CreateAlmanaxWebhookMentionsValueInner) GetIsRole() bool {
	if o == nil || IsNil(o.IsRole) {
		var ret bool
		return ret
	}
	return *o.IsRole
}

// GetIsRoleOk returns a tuple with the IsRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAlmanaxWebhookMentionsValueInner) GetIsRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRole) {
		return nil, false
	}
	return o.IsRole, true
}

// HasIsRole returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhookMentionsValueInner) HasIsRole() bool {
	if o != nil && !IsNil(o.IsRole) {
		return true
	}

	return false
}

// SetIsRole gets a reference to the given bool and assigns it to the IsRole field.
func (o *CreateAlmanaxWebhookMentionsValueInner) SetIsRole(v bool) {
	o.IsRole = &v
}

// GetPingDaysBefore returns the PingDaysBefore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAlmanaxWebhookMentionsValueInner) GetPingDaysBefore() int32 {
	if o == nil || IsNil(o.PingDaysBefore.Get()) {
		var ret int32
		return ret
	}
	return *o.PingDaysBefore.Get()
}

// GetPingDaysBeforeOk returns a tuple with the PingDaysBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAlmanaxWebhookMentionsValueInner) GetPingDaysBeforeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PingDaysBefore.Get(), o.PingDaysBefore.IsSet()
}

// HasPingDaysBefore returns a boolean if a field has been set.
func (o *CreateAlmanaxWebhookMentionsValueInner) HasPingDaysBefore() bool {
	if o != nil && o.PingDaysBefore.IsSet() {
		return true
	}

	return false
}

// SetPingDaysBefore gets a reference to the given NullableInt32 and assigns it to the PingDaysBefore field.
func (o *CreateAlmanaxWebhookMentionsValueInner) SetPingDaysBefore(v int32) {
	o.PingDaysBefore.Set(&v)
}
// SetPingDaysBeforeNil sets the value for PingDaysBefore to be an explicit nil
func (o *CreateAlmanaxWebhookMentionsValueInner) SetPingDaysBeforeNil() {
	o.PingDaysBefore.Set(nil)
}

// UnsetPingDaysBefore ensures that no value is present for PingDaysBefore, not even an explicit nil
func (o *CreateAlmanaxWebhookMentionsValueInner) UnsetPingDaysBefore() {
	o.PingDaysBefore.Unset()
}

func (o CreateAlmanaxWebhookMentionsValueInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAlmanaxWebhookMentionsValueInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscordId) {
		toSerialize["discord_id"] = o.DiscordId
	}
	if !IsNil(o.IsRole) {
		toSerialize["is_role"] = o.IsRole
	}
	if o.PingDaysBefore.IsSet() {
		toSerialize["ping_days_before"] = o.PingDaysBefore.Get()
	}
	return toSerialize, nil
}

type NullableCreateAlmanaxWebhookMentionsValueInner struct {
	value *CreateAlmanaxWebhookMentionsValueInner
	isSet bool
}

func (v NullableCreateAlmanaxWebhookMentionsValueInner) Get() *CreateAlmanaxWebhookMentionsValueInner {
	return v.value
}

func (v *NullableCreateAlmanaxWebhookMentionsValueInner) Set(val *CreateAlmanaxWebhookMentionsValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAlmanaxWebhookMentionsValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAlmanaxWebhookMentionsValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAlmanaxWebhookMentionsValueInner(val *CreateAlmanaxWebhookMentionsValueInner) *NullableCreateAlmanaxWebhookMentionsValueInner {
	return &NullableCreateAlmanaxWebhookMentionsValueInner{value: val, isSet: true}
}

func (v NullableCreateAlmanaxWebhookMentionsValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAlmanaxWebhookMentionsValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


