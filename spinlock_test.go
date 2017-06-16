package golocker

import "testing"

var spinLocker *SpinLocker

func init() {
	spinLocker = NewSpinLocker()
	spinLocker.MaxSpinTimes = 100000
}

func TestSpinLocker(t *testing.T) {

}
