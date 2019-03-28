# MessageHub

## 使用方法：

### 服务器端
1. 本地开启redis服务器 `redis-server /usr/local/etc/redis.conf`
2. 开启后端服务 `cd  notification_normal_go; sh startup.sh`

### 客户端（Web3.js）端
0. Prepare `npm install ethereumjs-tx` `npm install web3`
1. 开启本地客户端 `cd web3_client; sh startup.sh`
2. 登陆网站 `http://localhost:3000/user/login` 目前可用的账户是dcd，密码是123456
3. 相关服务器配置在 web3_client/app.js 和  web3_client/public/static/js/constant.js中


## 数据方接口
- **/dataclient/adddata**  
数据方上传你数据到区块链接口(交易形式,已实现)
- **/dataclient/pushdatatocomputing**  
数据方将原数据发送给计算方(交易形式,已实现)
- **/dataclient/aggreemodelclient**  
数据方同意模型方的交易(交易形式，已实现)
- **/dataclient/askcomputing**  
数据方请求运算方的运算资源(交易形式，已实现)
- **/dataclient/deletedata**  
数据方删除存储在消息服务器数据库上的MetadataIpfsHash(交易形式，未实现)
- **/dataclient/monitormetadata**  
数据方监听模型方的metadata请求(已实现)
- **/data/client/monitorcomputingaggree**  
数据方监听运算方的同意交易(已实现)



## 模型方接口



## 运算方接口