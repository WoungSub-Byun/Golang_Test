# Go routine & Channel
![goroutine](https://miro.medium.com/max/5216/1*dD8qcmhpjUlKxZ1GHmX1AQ.png)

## Go routine
---
- **비동기 Cocurrent 처리**를 구현하기 위해 만든 것
- **Go 런타임상에서 관리하기 때문에 OS 쓰레드와 1대1로 대응되지않고 Multiflexing으로 훨씬 적은 OS 쓰레드를 사용한다.**
- Go 런타임은 Go 루틴을 관리하면서 *Go 채널*을 통해 Go 루틴간의 통신을 쉽게 할 수 있도록 한다.

> 사용방법
> - main 함수에서 사용할 함수 앞에 ```go``` 키워드를 적어주면 된다.
```go
func main() {
    go [func]
}
```
> 비동기 처리 이해
```go
func main() {
    
    //동기적으로 실행
    goroutine("sync")
    
    //비동기적으로 실행
    go goroutine("async1")
    go goroutine("async2")
    go goroutine("async3")
    
    time.Sleep(time.Second * 3)
}

func goroutine(text string) {
    for i := 0; i < 5; i++ {
        fmt.Println(text, "**", i)
    }
}
```

- 위의 예제에서는 main함수가 goroutine이라는 함수를 동기적으로 1번 실행하고 비동기적으로 3번 실행하고 있다.
- 이때 동기적으로 호출한 goroutine함수가 완전히 끝났을 떄 나머지 **go goroutine은 Go 루틴들에서 동작하기 떄문에 메인루틴에서는 time.sleep을 실행한다.** 그리고 go goroutine은 실행순서가 일정하지 않으므로 프로그램 실행시마다 출력결과(순서)가 다르다.

## Channel
---
- ***goroutine들 사이 데이터를 주고 는 데 사용됨.***
- **상대편이 준비될 때까지 채널에서 대기함**으로써 별도의 lock을 걸지 않고 **데이터를 동기화하는데 사용**된다.

> 사용방법
> - ``` make()```함수를 통해 미리 생성되어야함
> - ``` <- ``` : 채널 연산자를 통해 데이터를 송수신함

```go
func main() {
  c := make(chan bool)   
  fmt.Println(<-c)
}

func goroutine(channel chan bool) {
    fmt.Println('test')
    channel <- true
}
```
- 위의 예제에서는 bool형 채널을 생성하고, goroutine에서 채널에 ```true```라는 bool형 데이터를 보내고 메인 루틴에서 받아 출력하는 코드이다.
- 채널을 생성할 때에는**어떤 타입의 데이터를 주고 받을지 make()함수에 지정**해주어야한다.
- 채널로 데이터를 보낼때 **```채널명 <- 데이터```**
- 채널로 데이터를 받을때 **```<- 채널명```**