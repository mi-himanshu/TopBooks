package tests

import (
    "testing"
)

func Test_Setup(t *testing.T) {
    var status = true
    if status != true {
        t.Errorf("Setup Initialization failed!")
    }
}