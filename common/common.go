package common

import (
    "blog/config"
    "blog/models"
    "sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {
    w := sync.WaitGroup{}
    w.Add(1)
    go func() {
        // 耗时
        var err error
        Template, err = models.InitTemplate(config.Config.SystemConfig.CurrentDir + "/template/")
        if err != nil {
            panic(err)
        }
        w.Done()
    }()
    w.Wait()
}
