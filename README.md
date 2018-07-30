# TIM SDK-腾讯“云通讯”服务服务端SDK(go语言) 
1.0.0 版发布了。常用的基础接口可用！

# USAGE

```
    // 初始化配置
    api = new(TimApp)
    api.AppID ="yourAppid"
    api.Identifiner = "yourIdentifiner"
    api.Sig = "yourSig"

    // 使用接口 eg:查询在线状态
    resp, err := api.QueryState([]string{"testAccount"})
    if err != nil {
        return err
    }
    if resp.ActionStatus == ResponseFail {
        fmt.Println(resp.ErrorInfo)
    } else {
        for _, acs := range resp.QueryResult {
            fmt.Printf("user `%s` is %s ", acs.ToAccount, acs.State)
        }
    }

```