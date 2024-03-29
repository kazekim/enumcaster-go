// Code generated by "enumer -type=Weekday"; DO NOT EDIT.

//
package enum

import (
	"fmt"
)

const _WeekdayName = "SundayMondayTuesdayWednesdayThursdayFridaySaturday"

var _WeekdayIndex = [...]uint8{0, 6, 12, 19, 28, 36, 42, 50}

func (i Weekday) String() string {
	if i < 0 || i >= Weekday(len(_WeekdayIndex)-1) {
		return fmt.Sprintf("Weekday(%d)", i)
	}
	return _WeekdayName[_WeekdayIndex[i]:_WeekdayIndex[i+1]]
}

var _WeekdayValues = []Weekday{0, 1, 2, 3, 4, 5, 6}

var _WeekdayNameToValueMap = map[string]Weekday{
	_WeekdayName[0:6]:   0,
	_WeekdayName[6:12]:  1,
	_WeekdayName[12:19]: 2,
	_WeekdayName[19:28]: 3,
	_WeekdayName[28:36]: 4,
	_WeekdayName[36:42]: 5,
	_WeekdayName[42:50]: 6,
}

// WeekdayString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func WeekdayString(s string) (Weekday, error) {
	if val, ok := _WeekdayNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Weekday values", s)
}

// WeekdayValues returns all values of the enum
func WeekdayValues() []Weekday {
	return _WeekdayValues
}

// IsAWeekday returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Weekday) IsAWeekday() bool {
	for _, v := range _WeekdayValues {
		if i == v {
			return true
		}
	}
	return false
}
