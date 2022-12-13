# Quantum Trial From Microsoft 

## Run & Build

`dotnet build ; dotnet run`

## Q & A

Problem we have met: 

`MSBUILD : error MSB1003: Specify a project or solution file. The current working directory does not contain a project or solution file.`

The reason is there is a need for a .csproj file to fetch info to build, but we don't have the file. So we recreated a Q# project using VS Code, and now we have this file. After then, you will see follows log message, that's means you can build it successful:

```
quantum-trial/quantum [master●] » dotnet build # dotnet version is 5.0.100
Microsoft (R) Build Engine version 16.8.0+126527ff1 for .NET
Copyright (C) Microsoft Corporation. All rights reserved.

  Determining projects to restore...
  All projects are up-to-date for restore.
  quantum -> /xxx/quantum-trial/quantum/bin/Debug/netcoreapp3.1/quantum.dll

Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:01.22
```
So, do not create .qs file only, also you need to create .csproj file.

## Reference

- [https://docs.microsoft.com/en-us/quantum/](https://docs.microsoft.com/en-us/quantum/) ~~机翻可耻！~~
