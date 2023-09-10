//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabase() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgresSQL,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBar() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// salah
//func InitializedHelloService() *HelloService {
//	wire.Build(NewHelloService, NewSayHelloImpl)
//	return nil
//}

// bind SayHello to SayHelloImpl
var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}
