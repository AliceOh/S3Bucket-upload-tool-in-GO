# s3backuptool
s3 backup tool in GO

# how to build
```
> go build 
```
this should generate an execution file named `s3backuptool.exe` on Windows or `s3backuptool` on Linux.

# how to run

```
> .\s3backuptool.exe version
s3backuptool.exe [Iress Content Team S3 Backup Tool] 1.0.0

> .\s3backuptool.exe uploads3
expected path argument
usage: D:\Iress\s3backuptool\s3backuptool.exe [global options...] uploadse [command options...] <file-to-upload> <s3uri> 

arguments:

  file        file name with path to push to s3 bucket
  s3uri       s3 uri to push zip to (s3://<bucketname>/<bucketpath>)


global options:

  --help          display help


> .\s3backuptool.exe --help  
s3backuptool.exe [global options...] <command> [command options...]

commands:

  uploads3         upload an environment configuration file
  version           display version

global options:

  --help          display help
```