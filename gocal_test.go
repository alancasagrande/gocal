package gocal

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ics = `HELLO:WORLD
BEGIN:VEVENT
DTSTART;VALUE=DATE:20141217
DTEND;VALUE=DATE:20141219
DTSTAMP:20151116T133227Z
UID:0001@example.net
CREATED:20141110T150010Z
DESCRIPTION:Amazing description on t
 wo lines
LAST-MODIFIED:20141110T150010Z
ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=ACCEPTED;CN=Antoin
 e Popineau;X-NUM-GUESTS=0:mailto:antoine.popineau@example.net
ATTENDEE;CUTYPE=INDIVIDUAL;ROLE=REQ-PARTICIPANT;PARTSTAT=ACCEPTED;CN=John
  Connor;X-NUM-GUESTS=0:mailto:john.connor@example.net
LOCATION:My Place
SEQUENCE:0
STATUS:CONFIRMED
SUMMARY:Lorem Ipsum Dolor Sit Amet
TRANSP:TRANSPARENT
END:VEVENT
BEGIN:VEVENT
DTSTART:20141203T130000Z
DTEND:20141203T163000Z
DTSTAMP:20151116T133227Z
UID:0002@google.com
CREATED:20141110T145426Z
DESCRIPTION:
LAST-MODIFIED:20141110T150016Z
LOCATION:Over there
SEQUENCE:1
STATUS:CONFIRMED
SUMMARY:The quick brown fox jumps over the lazy dog
TRANSP:TRANSPARENT
END:VEVENT`

func Test_Parse(t *testing.T) {
	gc := NewParser(strings.NewReader(ics))
	gc.Parse()

	assert.Equal(t, 2, len(gc.Events))

	assert.Equal(t, "Lorem Ipsum Dolor Sit Amet", gc.Events[0].Summary)
	assert.Equal(t, "0001@example.net", gc.Events[0].Uid)
	assert.Equal(t, "Amazing description on two lines", gc.Events[0].Description)
	assert.Equal(t, 2, len(gc.Events[0].Attendees))
	assert.Equal(t, "John Connor", gc.Events[0].Attendees[1].Cn)
}

func Test_ParseLine(t *testing.T) {
	gc := NewParser(strings.NewReader("HELLO;KEY1=value1;KEY2=value2: world"))
	gc.scanner.Scan()
	l, err, done := gc.parseLine()

	assert.Equal(t, nil, err)
	assert.Equal(t, true, done)

	assert.Equal(t, "HELLO", l.Key)
	assert.Equal(t, "world", l.Value)
	assert.Equal(t, map[string]string{"KEY1": "value1", "KEY2": "value2"}, l.Params)
}
