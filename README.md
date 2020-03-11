# 推送系统

## 一、WebSocket服务

### 数据传输

> 发送的数据：```[req1, req2, req3, ...]```<br>
> 接收的数据：```[rsp1, rsp2, rsp3, ...]```<br>

* JSON Over WebSocket
* 必定为数组，数组的每个元素都是一个独立的请求/回应
* 同一个请求与回应中的seq相同

---

### 命令一览

* login 登陆<br>
* enter 进入频道<br>
* exit 退出频道<br>
* snd2cli 发送至客户端<br>
* snd2usr 发送至用户<br>
* snd2chan 发送至频道<br>
* rcvdata 收到数据(仅客户端接收)<br>

---

### 协议框架

> 可根据具体需求进行灵活修订

#### 请求协议

```json
{
  "cmd": "login",
  "seq": 1,
  "immed": true,
  "data": {
  }
}
```

#### 回应协议

```json
{
  "cmd": "login",
  "seq": 1,
  "code": 0,
  "msg": "",
  "data": {
  }
}
```

#### 各协议data部分

##### login 登陆

```
发送数据
{
  "uid": 1001  // 用户标识
}

接收数据
{
  "id": 1  // 客户端标识
}
```

##### enter 进入频道

```
发送数据
{
  "chans": ["chan1", "chan2"]  // 进入的频道列表
}

接收数据
{
}
```

##### exit 离开频道

```
发送数据
{
  "chans": ["chan1", "chan2"]  // 离开频道列表
}

接收数据
{
}
```

##### snd2cli 发送至客户端

```
发送数据
{
  "ids": [1, 2, 5, 111]  // 客户端标识列表
  "data": {...}  // 数据体
}

接收数据
{
}
```

##### snd2usr 发送至用户

```
发送数据
{
  "uids": [1001, 1002]  // 用户标识列表
  "data": {...}  // 数据体
}

接收数据
{
}
```

##### snd2chan 发送至频道

```
发送数据
{
  "chans": ["world", "world/room1", "buy"]  // 频道标识列表
  "data": {...}  // 数据体
}

接收数据
{
}
```

##### rcvdata 收到数据

```
发送数据
{
}

接收数据
{
  "id": 3,  // 来源客户端标识
  "uid": 1001,  // 来源用户标识
  "chan": "world",  // 频道标识
  "data": {...}  // 数据体
}
```

---

## 二、HTTP服务

> 待补充
