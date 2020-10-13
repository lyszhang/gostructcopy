package gostructcopy

import (
	"testing"
)

func BenchmarkStructCopy(bench *testing.B) {
	a := struct {
		Id     int
		Name   string
		Weight int
		a      int
	}{100, "Dog", 200, 9}
	b := struct {
		Id   int
		Name string
		Desc string
		b    int
	}{}
	for i := 0; i < bench.N; i++ {
		StructCopyReflect(&a, &b)
	}
	for i := 0; i < bench.N; i++ {
		StructCopyGob(&a, &b)
	}
}

func TestStructCopy(t *testing.T) {
	a := struct {
		Id     int
		Name   string
		Weight int
		a      int
	}{100, "Dog", 200, 9}
	b := struct {
		Id   int
		Name string
		Desc string
		b    int
	}{}
	err := StructCopyReflect(&a, &b)
	if err != nil {
		t.Fatal(err)
	} else if a.Id == b.Id && a.Name == b.Name {
		t.Log("Success")
	} else {
		t.Fatal("Copy Fail")
	}
}
