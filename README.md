# CC98自动签到
主要是golang练手小玩具，每日自动签到98获取98币

修改`config.json`，没有的话初次运行会自动生成
```json
{
    "users": [
        {
            "username": "你的用户名",
            "password": "你的密码"
        },
    ]
}
```

运行
```bash
cc98sign
# 或者
go run .\cmd\cc98sign\ 
```

## TODO
- [ ] 定时脚本
- [ ] err检查