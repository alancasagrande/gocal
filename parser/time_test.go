package parser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ParseTimeUTC(t *testing.T) {
	ti, err := ParseTime("20150910T135212Z", map[string]string{}, TimeStart)

	assert.Equal(t, nil, err)

	assert.Equal(t, 10, ti.Day())
	assert.Equal(t, time.September, ti.Month())
	assert.Equal(t, 2015, ti.Year())
	assert.Equal(t, 13, ti.Hour())
	assert.Equal(t, 52, ti.Minute())
	assert.Equal(t, 12, ti.Second())
	assert.Equal(t, time.UTC, ti.Location())
}

func Test_ParseTimeTZID(t *testing.T) {
	ti, err := ParseTime("20150910T135212", map[string]string{"TZID": "Europe/Paris"}, TimeStart)
	tz, _ := time.LoadLocation("Europe/Paris")

	assert.Equal(t, nil, err)

	assert.Equal(t, 10, ti.Day())
	assert.Equal(t, time.September, ti.Month())
	assert.Equal(t, 2015, ti.Year())
	assert.Equal(t, 13, ti.Hour())
	assert.Equal(t, 52, ti.Minute())
	assert.Equal(t, 12, ti.Second())
	assert.Equal(t, tz, ti.Location())
}

func Test_ParseTimeAllDayStart(t *testing.T) {
	ti, err := ParseTime("20150910", map[string]string{"VALUE": "DATE"}, TimeStart)

	assert.Equal(t, nil, err)

	assert.Equal(t, 10, ti.Day())
	assert.Equal(t, time.September, ti.Month())
	assert.Equal(t, 2015, ti.Year())
	assert.Equal(t, 0, ti.Hour())
	assert.Equal(t, 0, ti.Minute())
	assert.Equal(t, 0, ti.Second())
	assert.Equal(t, time.UTC, ti.Location())
}

func Test_ParseTimeAllDayEnd(t *testing.T) {
	ti, err := ParseTime("20150910", map[string]string{"VALUE": "DATE"}, TimeEnd)

	assert.Equal(t, nil, err)

	assert.Equal(t, 10, ti.Day())
	assert.Equal(t, time.September, ti.Month())
	assert.Equal(t, 2015, ti.Year())
	assert.Equal(t, 23, ti.Hour())
	assert.Equal(t, 59, ti.Minute())
	assert.Equal(t, 59, ti.Second())
	assert.Equal(t, time.UTC, ti.Location())
}
