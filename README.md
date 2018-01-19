# fbftp

## 一、简介

支持本地磁盘和亚马逊的S3文件系统存储的ftp服务。（当前ftp只支持被动模式）

## 二、功能

已支持的ftp命令

|  命令  |          功能           | 存储类型支持 |
| :--: | :-------------------: | :----: |
| CDUP |        回到上层目录         | 本地、S3  |
| CWD  |        转换工作目录         | 本地、S3  |
| DELE |         删除文件          | 本地、S3  |
| FEAT |    列出所有的扩展命令与扩展功能     | 本地、S3  |
| HELP |     列出服务所支持的所有命令      | 本地、S3  |
| LIST |     列出当前目录的所有文件信息     | 本地、S3  |
| MKD  |      在当前目录下创建目录       | 本地、S3  |
| PASS |      用户密码（登陆命令）       | 本地、S3  |
| PASV |         被动模式          | 本地、S3  |
| PWD  |         当前路径          | 本地、S3  |
| QUIT |         退出服务          | 本地、S3  |
| RETR |         下载文件          | 本地、S3  |
| RNFR |    把xxx重命名（配合RNTO）    | 本地、S3  |
| RNTO |    重命名为xxx(配合RNFR)    | 本地、S3  |
| RMD  |      删除目录（目录应为空）      | 本地、S3  |
| SIZE |        获取文件大小         | 本地、S3  |
| STAT |        获取文件信息         | 本地、S3  |
| STOR |         上传文件          | 本地、S3  |
| SYST |     查明服务器上操作系统的类型     | 本地、S3  |
| TYPE | 确定数据的传输方式（目前只支持二进制传输） | 本地、S3  |
| USER |       用户名（登陆命令）       | 本地、S3  |



