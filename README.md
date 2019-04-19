[![pipeline status](https://api.travis-ci.org/bityuan/bityuan.svg?branch=master)](https://travis-ci.org/bityuan/bityuan/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bityuan/bityuan)](https://goreportcard.com/report/github.com/bityuan/bityuan)

# 基于 chain33 区块链开发 框架 开发的 TS公有链系统

#### 编译

```
git clone https://github.com/racalwang/tschain.git $GOPATH/src/github.com/racalwang/tschain
cd $GOPATH/src/github.com/racalwang/tschain
go build -i -o ts
go build -i -o ts-cli github.com/racalwang/tschain/cli
```

#### 运行

拷贝编译好的ts, ts-cli, tsChain.toml这三个文件置于同一个文件夹下，执行：

```
./ts -f tsChain.toml
```


