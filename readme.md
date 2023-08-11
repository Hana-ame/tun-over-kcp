# build

```sh
./build.sh
```

# usage

server

```sh
./tun-over-kcp.exe --server -laddr "localhost:22" -url "https://helper.moonchan.xyz/node"
```

client

```sh
./tun-over-kcp.exe --client -laddr ":9000" -url "https://helper.moonchan.xyz/node"
```

# binary

github的下载吃屎吧。

- [proxy.bin](https://moonchan.xyz/tun/proxy.bin)
- [helper.bin](https://moonchan.xyz/tun/helper.bin)
- [tun-over-kcp.bin](https://moonchan.xyz/tun/tun-over-kcp.bin)
- [tun-over-kcp.exe](https://moonchan.xyz/tun/tun-over-kcp.exe)


# known issue

- ~~和vps连会过一段时间死掉，不知道为什么~~ 看上去解决了，也不知道为什么
- ~~copyIO can err, dunno~~ 已经不是问题了
- quite slow... dunno why.
- 果然还是要用mux，在同一个tunnel当中传输
- 我不会debug别人写的东西，还是要耐心看才行吧。

# change log

- **go** newConn(url)
- var wg sync.WaitGroup
- N : 4 -> 16


## v0.1.1

- 修改了一些log提示
- 尝试给copyIO函数加上断开连接的提示信息，~~但失败了（所以server会不会因为conn关不掉溢出）~~ 成功了
- client端好像关了的，server端应该也会关吧。（有待验证）

worked about 2 hours.
### 测试

在google的cloud shell上，配合http proxy工作得很好

## v0.1.0

- it works on local but no further test.
- vps -> local will fail after a while

通了