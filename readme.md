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

和vps连会过一段时间死掉，不知道为什么