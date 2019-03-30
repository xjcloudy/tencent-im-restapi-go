# TIM SDK-腾讯“云通讯”服务服务端SDK(go语言) 
<a href="https://996.icu"><img src="https://img.shields.io/badge/link-996.icu-red.svg"></a>
1.0.0-beta版发布了目前项目中在使用，等项目上线测试通过后。发正式版。

# USAGE

```golang
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
