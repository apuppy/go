1.代码目录结构(目录在/root/go/下边)

├── build_project_calc
│   ├── bin
│   ├── pkg
│   └── src
│       ├── calc
│       │   └── calc.go
│       └── simplemath
│           ├── add.go
│           ├── add_test.go
│           ├── sqrt.go
│           └── sqrt_test.go

2.编辑~/.bashrc
export GOPATH=/root/go/build_project_calc/

3.编译
cd /root/go/build_project_calc/bin
go build calc

4.运行程序
./calc
./calc add 1 4
./calc sqrt 9

5.测试
go test simplemath
