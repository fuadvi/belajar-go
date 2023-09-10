package simple

type SayHello interface {
	Hello(Name string) string
}

type HelloService struct {
	SayHello
}

type SayHelloImpl struct {
}

func (s SayHelloImpl) Hello(Name string) string {
	return "Hello " + Name
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}
