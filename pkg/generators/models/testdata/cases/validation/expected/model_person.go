// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Test
//     Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"regexp"

	"time"
)

// PersonCronPatternError is the error message returned for pattern validation errors on Person.Cron
var PersonCronPatternError = validation.NewError("validation_Cron_pattern_invalid", "must be a valid cron value")

// personCronPattern is the validation pattern for Person.Cron
var personCronPattern = regexp.MustCompile(`(@(annually|yearly|monthly|weekly|daily|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7})`)

// personPomodoroPattern is the validation pattern for Person.Pomodoro
var personPomodoroPattern = regexp.MustCompile(`^\d{1,2}m$`)

// Person is an object.
type Person struct {
	// Address:
	Address Address `json:"address,omitempty"`
	// Age:
	Age float32 `json:"age,omitempty"`
	// Base64:
	Base64 string `json:"base64,omitempty"`
	// Cron:
	Cron string `json:"cron"`
	// Date:
	Date string `json:"date,omitempty"`
	// Datetime:
	Datetime time.Time `json:"datetime,omitempty"`
	// Email:
	Email string `json:"email,omitempty"`
	// FavoriteColors:
	FavoriteColors []Color `json:"favoriteColors"`
	// Gender:
	Gender Gender `json:"gender"`
	// Hostname:
	Hostname string `json:"hostname,omitempty"`
	// Ip:
	Ip string `json:"ip,omitempty"`
	// Ipv4:
	Ipv4 string `json:"ipv4,omitempty"`
	// Ipv6:
	Ipv6 string `json:"ipv6,omitempty"`
	// Name:
	Name string `json:"name"`
	// Pomodoro:
	Pomodoro string `json:"pomodoro,omitempty"`
	// RequestURI:
	RequestURI string `json:"requestURI,omitempty"`
	// SecondGender:
	SecondGender *Gender `json:"secondGender,omitempty"`
	// Uri:
	Uri string `json:"uri,omitempty"`
	// Url:
	Url string `json:"url,omitempty"`
	// Uuid:
	Uuid string `json:"uuid,omitempty"`
}

// Validate implements basic validation for this model
func (m Person) Validate() error {
	return validation.Errors{
		"address": validation.Validate(
			m.Address,
		),
		"age": validation.Validate(
			m.Age, validation.Min(float32(18)), validation.Max(float32(120)),
		),
		"base64": validation.Validate(
			m.Base64, is.Base64,
		),
		"cron": validation.Validate(
			m.Cron, validation.Match(personCronPattern).ErrorObject(PersonCronPatternError),
		),
		"email": validation.Validate(
			m.Email, is.EmailFormat,
		),
		"favoriteColors": validation.Validate(
			m.FavoriteColors, validation.Required, validation.Length(1, 0),
		),
		"gender": validation.Validate(
			m.Gender, validation.Required,
		),
		"hostname": validation.Validate(
			m.Hostname, is.Host,
		),
		"ip": validation.Validate(
			m.Ip, is.IP,
		),
		"ipv4": validation.Validate(
			m.Ipv4, is.IPv4,
		),
		"ipv6": validation.Validate(
			m.Ipv6, is.IPv6,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(2, 32),
		),
		"pomodoro": validation.Validate(
			m.Pomodoro, validation.Match(personPomodoroPattern),
		),
		"requestURI": validation.Validate(
			m.RequestURI, is.RequestURL.Error("must be valid URI with scheme"),
		),
		"secondGender": validation.Validate(
			m.SecondGender,
		),
		"uri": validation.Validate(
			m.Uri, is.RequestURI,
		),
		"url": validation.Validate(
			m.Url, is.URL.Error("must be a valid URL with HTTP or HTTPS scheme"),
		),
		"uuid": validation.Validate(
			m.Uuid, is.UUID,
		),
	}.Filter()
}
