package main

import "fmt"

// "github.com/gch2008net/common-go/greetings"
// "github.com/gch2008net/common-go/greetings/book/api"
// "github.com/gch2008net/common-go/hello"

// "example.com/greetings"

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	// api.GetInitInfo()

	// msg, err := greetings.Hello("安妮")
	// if err == nil {

	// 	fmt.Println("msg:" + msg)
	// }

	// hello.Say()

	// message, err := greetings.Hello("Gladys")
	// if err != nil {
	// 	fmt.Println("err:" + err.Error())
	// }
	// fmt.Println(message)

	// message, err := greetings.Hello("")
	// if err != nil {
	// 	fmt.Println("err:" + err.Error())
	// }
	// fmt.Println("msg:" + message)

	// // Request a greeting message.
	// message, err := greetings.Hello("李刚")
	// // If an error was returned, print it to the console and
	// // exit the program.
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println(err)
	// }

	// // If no error was returned, print the returned message
	// // to the console.
	// fmt.Println(message)

	// // A slice of names.
	// names := []string{"张三", "李四", "老六"}

	// // Request greeting messages for the names.
	// messages, err := greetings.Hellos(names)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // If no error was returned, print the returned map of
	// // messages to the console.
	// fmt.Println(messages)
	// test77.Hello222()

	// 创建一个Dog对象
	dog := Dog{}

	// 调用子类的方法，实际调用的是重写后的方法
	dog.Sound()

}

// 定义父类
type Animal struct {
	Name string
}

// 父类的方法
func (a *Animal) Sound() {
	fmt.Println("Animal makes a sound")
	a.Name = "这个是父类Animal的Name属性"

	fmt.Println(a.Name)
}

// 定义子类
type Dog struct {
	Animal // 嵌入类型，相当于子类继承了父类的字段和方法
}

// 子类重写父类的方法
func (d *Dog) Sound() {
	fmt.Println("Dog barks")

	d.Name = "这个是子类Dog的Name属性"

	fmt.Println(d.Name)
}
