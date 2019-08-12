package main

import (
    "testing"
)

func Testmain(t *testing.T) {
    t.Run("Test something", func(t *testing.T) {
        if err := main("user.json"); (err != nil) != false {
            t.Errorf("analyze() error = %v", err)
        }
    })
}