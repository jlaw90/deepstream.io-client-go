package deepstreamio

type EventListener func(eventName string, args interface{})
