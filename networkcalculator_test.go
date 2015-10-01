package main

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestGetRxBytesPerSecond(t *testing.T) {
	// GIVEN
	var oldDate = time.Now()
	// 1000000000 nanoseconds = 1 second
	var newDate = oldDate.Add(time.Duration(1000000000))

	// old rxBytes = 10
	var oldData = NetworkData{oldDate, 10, 0, 0, 0, 0, 0, 0, 0}
	// new rxBytes = 110
	var newData = NetworkData{newDate, 110, 0, 0, 0, 0, 0, 0, 0}
	var calculator = NetworkCalculator{oldData, newData}

	// WHEN
	value := calculator.getRxBytesPerSecond()

	// THEN
	// value should be 100 bytes / second
	assert.Equal(t, float64(100), value)
}

func TestGetRxDroppedPerSecond(t *testing.T) {
	// GIVEN
	var oldDate = time.Now()
	// 2000000000 nanoseconds = 2 seconds
	var newDate = oldDate.Add(time.Duration(2000000000))

	// old rxDropped = 20
	var oldData = NetworkData{oldDate, 0, 20, 0, 0, 0, 0, 0, 0}
	// new rxDropped = 120
	var newData = NetworkData{newDate, 0, 120, 0, 0, 0, 0, 0, 0}
	var calculator = NetworkCalculator{oldData, newData}

	// WHEN
	value := calculator.getRxDroppedPerSecond()

	// THEN
	// value should be 50 dropped / second
	assert.Equal(t, float64(50), value)
}

func TestGetRxErrorsPerSecond(t *testing.T) {
	// GIVEN
	var oldDate = time.Now()
	// 10000000000 nanoseconds = 10 seconds
	var newDate = oldDate.Add(time.Duration(10000000000))

	// old rxErrors = 30
	var oldData = NetworkData{oldDate, 0, 0, 30, 0, 0, 0, 0, 0}
	// new rxErrors = 31
	var newData = NetworkData{newDate, 0, 0, 31, 0, 0, 0, 0, 0}
	var calculator = NetworkCalculator{oldData, newData}

	// WHEN
	value := calculator.getRxErrorsPerSecond()

	// THEN
	// value should be 0.1 error / second
	assert.Equal(t, float64(0.1), value)
}

func TestGetRxPacketsPerSecond(t *testing.T) {
	// GIVEN
	var oldDate = time.Now()
	// 10000000000 nanoseconds = 10 seconds
	var newDate = oldDate.Add(time.Duration(10000000000))

	// old rxPackets = 40
	var oldData = NetworkData{oldDate, 0, 0, 0, 40, 0, 0, 0, 0}
	// new rxErrors = 140
	var newData = NetworkData{newDate, 0, 0, 0, 140, 0, 0, 0, 0}
	var calculator = NetworkCalculator{oldData, newData}

	// WHEN
	value := calculator.getRxPacketsPerSecond()

	// THEN
	// value should be 10 packets / second
	assert.Equal(t, float64(10), value)
}

func TestGetTxBytesPerSecond(t *testing.T) {
	// GIVEN
	var oldDate = time.Now()
	// 1000000000 nanoseconds = 1 second
	var newDate = oldDate.Add(time.Duration(1000000000))

	// old txBytes = 10
	var oldData = NetworkData{oldDate, 0, 0, 0, 0, 10, 0, 0, 0}
	// new txBytes = 10
	var newData = NetworkData{newDate, 0, 0, 0, 0, 10, 0, 0, 0}
	var calculator = NetworkCalculator{oldData, newData}

	// WHEN
	value := calculator.getTxBytesPerSecond()

	// THEN
	// value should be 0 bytes / second
	assert.Equal(t, float64(0), value)
}

func TestGetTxDroppedPerSecond(t *testing.T) {
	// GIVEN
	var oldDate = time.Now()
	// 2000000000 nanoseconds = 2 seconds
	var newDate = oldDate.Add(time.Duration(2000000000))

	// old txDropped = 0
	var oldData = NetworkData{oldDate, 0, 0, 0, 0, 0, 0, 0, 0}
	// new txDropped = 0
	var newData = NetworkData{newDate, 0, 0, 0, 0, 0, 0, 0, 0}
	var calculator = NetworkCalculator{oldData, newData}

	// WHEN
	value := calculator.getTxDroppedPerSecond()

	// THEN
	// value should be 0 dropped / second
	assert.Equal(t, float64(0), value)
}

func TestGetTxErrorsPerSecond(t *testing.T) {
	// GIVEN
	var oldDate = time.Now()
	// 10000000000 nanoseconds = 10 seconds
	var newDate = oldDate.Add(time.Duration(10000000000))

	// old txErrors = 70
	var oldData = NetworkData{oldDate, 0, 0, 0, 0, 0, 0, 70, 0}
	// new txErrors = 170
	var newData = NetworkData{newDate, 0, 0, 0, 0, 0, 0, 170, 0}
	var calculator = NetworkCalculator{oldData, newData}

	// WHEN
	value := calculator.getTxErrorsPerSecond()

	// THEN
	// value should be 10 error / second
	assert.Equal(t, float64(10), value)
}

func TestGetTxPacketsPerSecond(t *testing.T) {
	// GIVEN
	var oldDate = time.Now()
	// 10000000000 nanoseconds = 10 seconds
	var newDate = oldDate.Add(time.Duration(10000000000))

	// old txPackets = 80
	var oldData = NetworkData{oldDate, 0, 0, 0, 0, 0, 0, 0, 80}
	// new txErrors = 92
	var newData = NetworkData{newDate, 0, 0, 0, 0, 0, 0, 0, 92}
	var calculator = NetworkCalculator{oldData, newData}

	// WHEN
	value := calculator.getTxPacketsPerSecond()

	// THEN
	// value should be 0 packets / second when old value > new value
	assert.Equal(t, float64(1.2), value)
}