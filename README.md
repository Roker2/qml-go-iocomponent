# qml-go-iocomponent

Input Output component 

## About package

It is simple package for working with text files.

## How to use

### Initialization

You should to register package as QML component.

> ```go
> func run() error {
>    downloader.Register()
>    iocomponent.Register("by.roker2.mega")
>    ui.SetEngine()
> 
>    ui.InitModels()
>    err := ui.SetComponent()
>    if err != nil {
>       return err
>    }
>    ui.Win.Show()
>    ui.Win.Wait()
> 
>    return nil
> }
> ```