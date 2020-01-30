package main

import (
	"log"

	"golang.org/x/net/context"

	coprocess "github.com/TykTechnologies/tyk-protobuf/bindings/go"
)

// Dispatcher implementation
type Dispatcher struct{}

// Dispatch will be called on every request:
func (d *Dispatcher) Dispatch(ctx context.Context, object *coprocess.Object) (*coprocess.Object, error) {
	log.Print("Dispatcher is called")
	// We dispatch the object based on the hook name (as specified in the manifest file), these functions are in hooks.go:
	switch object.HookName {
	case "MyPreHook1":
		log.Println("MyPreHook1 is called!")
		return MyPreHook1(object)
	case "MyPreHook2":
		log.Println("MyPreHook2 is called!")
		return MyPreHook2(object)
	}

	log.Println("Unknown hook: ", object.HookName)

	return object, nil
}

// DispatchEvent will be called when a Tyk event is triggered:
func (d *Dispatcher) DispatchEvent(ctx context.Context, event *coprocess.Event) (*coprocess.EventReply, error) {
	return &coprocess.EventReply{}, nil
}
