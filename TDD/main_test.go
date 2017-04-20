package TDD

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestInInterval_exclude1_AND_Include6_input3(t *testing.T) {
	//arrange
	// "( 1, 6]"
	scope := Interval{
		Begin:        1,
		IncludeBegin: false,
		End:          6,
		IncludeEnd:   true,
	}
	var target = 3

	//act
	ok := InInterval(scope, target)

	//assert
	assert := assert.New(t)
	assert.True(ok)
}

func TestInInterval_excludeMinus2_AND_Include7_input3(t *testing.T) {
	//arrange
	// "( 1, 6]"
	scope := Interval{
		Begin:        -2,
		IncludeBegin: false,
		End:          7,
		IncludeEnd:   true,
	}
	var target = 3

	//act
	ok := InInterval(scope, target)

	//assert
	assert := assert.New(t)
	assert.True(ok)
}

func TestInInterval_exclude1_AND_Include6_input1(t *testing.T) {
	//arrange
	//"( 1, 6]"
	scope := Interval{
		Begin:        1,
		IncludeBegin: false,
		End:          6,
		IncludeEnd:   true,
	}
	var target = 1

	//act
	ok := InInterval(scope, target)

	//assert
	assert := assert.New(t)
	assert.False(ok)
}

func TestInInterval_exclude1_AND_Include6_input6(t *testing.T) {
	//arrange
	//"( 1, 6]"
	scope := Interval{
		Begin:        1,
		IncludeBegin: false,
		End:          6,
		IncludeEnd:   true,
	}
	var target = 6

	//act
	ok := InInterval(scope, target)

	//assert
	assert := assert.New(t)
	assert.True(ok)
}

func TestInInterval_exclude1_AND_Exclude6_input1(t *testing.T) {
	//arrange
	//"( 1, 6)"
	scope := Interval{
		Begin:        1,
		IncludeBegin: false,
		End:          6,
		IncludeEnd:   false,
	}
	var target = 1

	//act
	ok := InInterval(scope, target)

	//assert
	assert := assert.New(t)
	assert.False(ok)
}

func TestInInterval_exclude1_AND_Exclude6_input6(t *testing.T) {
	//arrange
	//"( 1, 6)"
	scope := Interval{
		Begin:        1,
		IncludeBegin: false,
		End:          6,
		IncludeEnd:   false,
	}
	var target = 6

	//act
	ok := InInterval(scope, target)

	//assert
	assert := assert.New(t)
	assert.False(ok)
}

func TestGetNowString(t *testing.T) {
	//arange
	beforeTime := time.Now()

	//act
	var resultStr = GetNowString()

	afterTime := beforeTime.Add(1)
	//assert

	before := beforeTime.UnixNano()
	after := afterTime.UnixNano()
	resultTime, err := time.Parse(time.RFC3339Nano, resultStr)
	assert := assert.New(t)
	if err != nil {
		assert.FailNow("not valid time string")
	}
	result := resultTime.UnixNano()
	assert.True(result <= after)
	assert.True(result >= before)

}
