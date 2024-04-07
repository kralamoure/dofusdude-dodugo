/*
dofusdude

# A project for you - the developer. The all-in-one toolbelt for your next Ankama related project.  ## Client SDKs - [Javascript](https://github.com/dofusdude/dofusdude-js) `npm i dofusdude-js --save` - [Typescript](https://github.com/dofusdude/dofusdude-ts) `npm i dofusdude-ts --save` - [Go](https://github.com/dofusdude/dodugo) `go get -u github.com/dofusdude/dodugo` - [Python](https://github.com/dofusdude/dofusdude-py) `pip install dofusdude` - [PHP](https://github.com/dofusdude/dofusdude-php) - [Java](https://github.com/dofusdude/dofusdude-java) Maven with GitHub packages setup  Everything, including this site, is generated out of the [Docs Repo](https://github.com/dofusdude/api-docs). Consider it the Single Source of Truth. If there is a problem with the SDKs, create an issue there.  Your favorite language is missing? Please let me know!  # Main Features - 🥷 **Seamless Auto-Update** load data in the background when a new Dofus version is released and serving it within 10 minutes with atomic data source switching. No downtime and no effects for the user, just always up-to-date.  - ⚡ **Blazingly Fast** all data in-memory, aggressive caching over short time spans, HTTP/2 multiplexing, written in Go, optimized for low latency, hosted on bare metal in 🇩🇪.  - 📨 **Discord Integration** Ankama related RSS and Almanax feeds to post to Discord servers with advanced features like filters or mentions. Use the endpoints as a dev or the official [Web Client](https://discord.dofusdude.com) as a user.  - 🩸 **Dofus 2 Beta** from stable to bleeding edge by replacing /dofus2 with /dofus2beta.  - 🗣️ **Multilingual** supporting _en_, _fr_, _es_, _pt_ including the dropped languages from the Dofus website _de_ and _it_.  - 🧠 **Search by Relevance** allowing typos in name and description, handled by language specific text analysis and indexing.  - 🕵️ **Complete** actual data from the game including items invisible to the encyclopedia like quest items.  - 🖼️ **HD Images** rendering game assets to high-res images with up to 800x800 px.  ... and much more on the Roadmap on my [Discord](https://discord.gg/3EtHskZD8h). 

API version: 0.9.0
Contact: stelzo@steado.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dodugo

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateRSSWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRSSWebhook{}

// CreateRSSWebhook 
type CreateRSSWebhook struct {
	Whitelist []string `json:"whitelist,omitempty"`
	Blacklist []string `json:"blacklist,omitempty"`
	// Get the available subscriptions with /meta/webhooks/rss
	Subscriptions []string `json:"subscriptions"`
	Format string `json:"format"`
	PreviewLength NullableInt32 `json:"preview_length,omitempty"`
	// Discord Webhook URL
	Callback string `json:"callback"`
}

type _CreateRSSWebhook CreateRSSWebhook

// NewCreateRSSWebhook instantiates a new CreateRSSWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRSSWebhook(subscriptions []string, format string, callback string) *CreateRSSWebhook {
	this := CreateRSSWebhook{}
	this.Subscriptions = subscriptions
	this.Format = format
	this.Callback = callback
	return &this
}

// NewCreateRSSWebhookWithDefaults instantiates a new CreateRSSWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRSSWebhookWithDefaults() *CreateRSSWebhook {
	this := CreateRSSWebhook{}
	return &this
}

// GetWhitelist returns the Whitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRSSWebhook) GetWhitelist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Whitelist
}

// GetWhitelistOk returns a tuple with the Whitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRSSWebhook) GetWhitelistOk() ([]string, bool) {
	if o == nil || IsNil(o.Whitelist) {
		return nil, false
	}
	return o.Whitelist, true
}

// HasWhitelist returns a boolean if a field has been set.
func (o *CreateRSSWebhook) HasWhitelist() bool {
	if o != nil && !IsNil(o.Whitelist) {
		return true
	}

	return false
}

// SetWhitelist gets a reference to the given []string and assigns it to the Whitelist field.
func (o *CreateRSSWebhook) SetWhitelist(v []string) {
	o.Whitelist = v
}

// GetBlacklist returns the Blacklist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRSSWebhook) GetBlacklist() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Blacklist
}

// GetBlacklistOk returns a tuple with the Blacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRSSWebhook) GetBlacklistOk() ([]string, bool) {
	if o == nil || IsNil(o.Blacklist) {
		return nil, false
	}
	return o.Blacklist, true
}

// HasBlacklist returns a boolean if a field has been set.
func (o *CreateRSSWebhook) HasBlacklist() bool {
	if o != nil && !IsNil(o.Blacklist) {
		return true
	}

	return false
}

// SetBlacklist gets a reference to the given []string and assigns it to the Blacklist field.
func (o *CreateRSSWebhook) SetBlacklist(v []string) {
	o.Blacklist = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *CreateRSSWebhook) GetSubscriptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *CreateRSSWebhook) GetSubscriptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *CreateRSSWebhook) SetSubscriptions(v []string) {
	o.Subscriptions = v
}

// GetFormat returns the Format field value
func (o *CreateRSSWebhook) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *CreateRSSWebhook) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *CreateRSSWebhook) SetFormat(v string) {
	o.Format = v
}

// GetPreviewLength returns the PreviewLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRSSWebhook) GetPreviewLength() int32 {
	if o == nil || IsNil(o.PreviewLength.Get()) {
		var ret int32
		return ret
	}
	return *o.PreviewLength.Get()
}

// GetPreviewLengthOk returns a tuple with the PreviewLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRSSWebhook) GetPreviewLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviewLength.Get(), o.PreviewLength.IsSet()
}

// HasPreviewLength returns a boolean if a field has been set.
func (o *CreateRSSWebhook) HasPreviewLength() bool {
	if o != nil && o.PreviewLength.IsSet() {
		return true
	}

	return false
}

// SetPreviewLength gets a reference to the given NullableInt32 and assigns it to the PreviewLength field.
func (o *CreateRSSWebhook) SetPreviewLength(v int32) {
	o.PreviewLength.Set(&v)
}
// SetPreviewLengthNil sets the value for PreviewLength to be an explicit nil
func (o *CreateRSSWebhook) SetPreviewLengthNil() {
	o.PreviewLength.Set(nil)
}

// UnsetPreviewLength ensures that no value is present for PreviewLength, not even an explicit nil
func (o *CreateRSSWebhook) UnsetPreviewLength() {
	o.PreviewLength.Unset()
}

// GetCallback returns the Callback field value
func (o *CreateRSSWebhook) GetCallback() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value
// and a boolean to check if the value has been set.
func (o *CreateRSSWebhook) GetCallbackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Callback, true
}

// SetCallback sets field value
func (o *CreateRSSWebhook) SetCallback(v string) {
	o.Callback = v
}

func (o CreateRSSWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRSSWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Whitelist != nil {
		toSerialize["whitelist"] = o.Whitelist
	}
	if o.Blacklist != nil {
		toSerialize["blacklist"] = o.Blacklist
	}
	toSerialize["subscriptions"] = o.Subscriptions
	toSerialize["format"] = o.Format
	if o.PreviewLength.IsSet() {
		toSerialize["preview_length"] = o.PreviewLength.Get()
	}
	toSerialize["callback"] = o.Callback
	return toSerialize, nil
}

func (o *CreateRSSWebhook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptions",
		"format",
		"callback",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateRSSWebhook := _CreateRSSWebhook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRSSWebhook)

	if err != nil {
		return err
	}

	*o = CreateRSSWebhook(varCreateRSSWebhook)

	return err
}

type NullableCreateRSSWebhook struct {
	value *CreateRSSWebhook
	isSet bool
}

func (v NullableCreateRSSWebhook) Get() *CreateRSSWebhook {
	return v.value
}

func (v *NullableCreateRSSWebhook) Set(val *CreateRSSWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRSSWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRSSWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRSSWebhook(val *CreateRSSWebhook) *NullableCreateRSSWebhook {
	return &NullableCreateRSSWebhook{value: val, isSet: true}
}

func (v NullableCreateRSSWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRSSWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


