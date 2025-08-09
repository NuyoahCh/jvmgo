# jvm.go

# Introduction

jvm.go is a toy JVM (which is far from complete) programmed in Go. The main purpose of this project is learning Go and the JVM. So the number one goal of the project is readability of code. The basic idea is to just implement the core JVM, and use rt.jar (from OpenJDK) as its class library. The garbage collector is implemented by directly using Go’s GC.

# Build jvm.go
```bash
git clone https://github.com/NuyoahCh/jvmgo.git
cd jvm.go
go build github.com/NuyoahCh/jvmgo/cmd/java
```

# Material & Book 
https://www.yuque.com/codereview1024/uboc4y/tfrw4723yg6d37ul