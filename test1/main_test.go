package main

import "testing"
import "test1/help"


func TestHelloWorld(t *testing.T) {
    t.Log("hello world")
}

func TestAdd(t *testing.T){

if ans := help.Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}
}