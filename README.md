# qml-go-iocomponent

Input Output component 

## About package

It is simple package for working with text files.

## How to use

### Initialization in Go

You should to register package as QML component.

> ```go
> func run() error {
>     iocomponent.Register("fullappname")
>     ui.SetEngine()
>
>     ui.InitModels()
>     err := ui.SetComponent()
>     if err != nil {
>         return err
>     }
>     ui.Win.Show()
>     ui.Win.Wait()
>
>     return nil
> }
> ```

"fullappname" is name of your program, for example "by.roker2.mega".

### 3 writable folders

UBports provide 3 writable folders for save and read your program data:

- /home/phablet/.config/fullappname/

- /home/phablet/.cache/fullappname/

- /home/phablet/.local/share/fullappname/

This package use "folderType" for choosing folder. "folderType" is int

- Config is 0
- Cache is 1
- AppData is 2

### Initialization in QML

```javascript
import GoIOComponent 0.1

IOComponent {
    id: io
}
```

### Functions

##### Write to file

```go
WriteToFile(ft folderType, fileName string, text string)
```

#### Read from file

```go
ReadFromFile(ft folderType, fileName string) string
```

#### File is exist

```go
FileIsExist(ft folderType, fileName string) bool
```

#### Create file

```go
CreateFile(ft folderType, fileName string)
```

#### Remove file

```go
RemoveFile(ft folderType, fileName string)
```

