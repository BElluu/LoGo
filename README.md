# LoGo - Work In Progress
Small and simple logger for GoLang

## Usage
All you need is a setting where you want to use the logs (to console and/or to file) and the minimum log level.

```golang
LoggerConfiguration.
		WriteTo().
		Console().
		File("MyTestApp", "C:\PathForLogsLocation").
		MinimumLevel(LogLevel.DEBUG)
```
```golang
	Log.Information("This is information event.", "But I want to add time too: ", time.Now().Format("15:04:05"))
	Log.Warning("Just warning event :)")
	Log.Debug("I hope you will not use this in production")
	Log.Error("Something went wrong... :(")
	Log.Fatal("This should never, ever happen!", "Call to support: ", 123456789)
```
### Console
![image](https://user-images.githubusercontent.com/11333925/166064416-52ea8ad9-93cf-46c9-8cd8-ef8d5c3f9ddc.png)

### File
![image](https://user-images.githubusercontent.com/11333925/166064486-deb07b56-fd9a-4db4-b928-76fb594b5d8b.png)

## License
[MIT](https://choosealicense.com/licenses/mit/)
