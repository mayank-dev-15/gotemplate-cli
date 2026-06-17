package gotemplate_cli
import "testing"
func TestEngine(t *testing.T){ e:=NewEngine(); e.Init(); if !e.ready{ t.Fatal("not ready") }}
