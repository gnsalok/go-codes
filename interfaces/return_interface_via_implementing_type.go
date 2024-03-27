package main

import "fmt"

type EngineImpl struct{}

type interfaceEngine interface {
	startEngine(engineName string) string
}

type Config struct{}

func New(config *Config) interfaceEngine {

	return &EngineImpl{}

}

func (engimpl *EngineImpl) startEngine(en string) string {
	return "Engine started..."

}

func main() {
	fmt.Println("Testing interface return type")
}
