package quickfixgo

import (
	. "launchpad.net/gocheck"
)

var _ = Suite(&FieldMapTests{})

type FieldMapTests struct{ fieldMap FieldMap }

func (s *FieldMapTests) SetUpTest(c *C) {
	s.fieldMap.init()
}

func (s *FieldMapTests) TestLength(c *C) {
	s.fieldMap.append(1, newFieldBytes(1, []byte("hello")))
	s.fieldMap.append(2, newFieldBytes(2, []byte("world")))
	s.fieldMap.append(8, newFieldBytes(8, []byte("FIX.4.4")))
	s.fieldMap.append(9, newFieldBytes(9, []byte("100")))
	s.fieldMap.append(10, newFieldBytes(10, []byte("100")))

	c.Check(s.fieldMap.length(), Equals, 16,
		Commentf("Length includes all fields but beginString, bodyLength, and checkSum"))
}

func (s *FieldMapTests) TestLengthDupTags(c *C) {
	s.fieldMap.append(1, newFieldBytes(1, []byte("hello")))
	s.fieldMap.append(1, newFieldBytes(1, []byte("dup")))
	s.fieldMap.append(2, newFieldBytes(2, []byte("world")))

	c.Check(s.fieldMap.length(), Equals, 22)
}