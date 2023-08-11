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

# known issue

- ~~和vps连会过一段时间死掉，不知道为什么~~ 看上去解决了，也不知道为什么
- ~~copyIO can err, dunno~~ 已经不是问题了

# change log

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