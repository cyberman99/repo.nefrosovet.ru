# Command line startup
 go run main.go --broker tcp://0.0.0.0:1883 --action pub --password {} --username datarouter --replaceWithId benchtest --topic services/mock/OUT --filepath ./example/datarouter.json --count 5 --clients 100
# API
    execOpts := tool.ExecOptions{}
    execOpts.Broker = brokerHostPort
    execOpts.Qos = byte(qos)
    execOpts.Retain = retain
    execOpts.Topic = topic
    execOpts.Username = username
    execOpts.Password = password
    execOpts.CertConfig = certConfig
    execOpts.ClientNum = clients
    execOpts.Count = count
    execOpts.UseDefaultHandler = useDefaultHandler
    execOpts.PreTime = preTime
    execOpts.IntervalTime = intervalTime
    execOpts.TargetMPS = float64(viper.GetInt("mps"))
    execOpts.ReplaceValueWithID = viper.GetString("replace")
    
    tool.Debug = debug
    
    switch method {
    case "pub":
        err = tool.Execute(tool.PublishAllClient, execOpts, message)
    case "sub":
        err =tool.Execute(tool.SubscribeAllClient, execOpts, message)
    }