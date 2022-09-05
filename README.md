> 基于区块链的高校助学金系统。本项目使用Hyperledger Fabric构建区块链网络, go编写智能合约，应用层使用gin+fabric-sdk-go调用合约。前端展示使用vue+element。前后端分离。


> 本项目借鉴了https://github.com/togettoyou/fabric-realty.git ，项目原型是房地产系统，感谢有大神的指引，得以完成该项目。

## 技术栈

- Hyperledger Fabric
- Go
- Vue

## 部署环境

- Docker
- Docker Compose

## 前提

Mac，要求安装了 Docker 和 Docker Compose

## 运行

### 1、克隆本项目放在任意目录下，例：`/root/blockchain-real-estate`

### 2、给予项目权限，执行 `sudo chmod -R +x /root/blockchain-real-estate/`

### 3、进入 `deploy` 目录，执行 `./start.sh` 启动区块链网络

### 4、进入 `vue` 目录，执行 `./build.sh` 编译前端

### 5、进入 `application` 目录，执行 `./build.sh` 编译后端

### 6、在 `application` 目录下，执行 `./start.sh` 启动应用

### 7、浏览器访问 [http://localhost:8000/web](http://localhost:8000/web)

### 8、进入 `deploy/explorer` 目录，执行 `./start.sh` 启动区块链浏览器

报错Failed to create wallet, please check the configuration, and valid file paths:
需要cd 进入```blockchain-real-estate/deploy/crypto-config/peerOrganizations/org1.blockchainrealestate.com/users/Admin@org1.blockchainrealestate.com/msp/keystore/ce9a70db5c4e5437234ba47c579e7bf5bf11efa7b21f501c00c5a1bcd0b4e980_sk```，将sk文件的名字复制下来，替换掉network_temp.json与network.json文件内path的priv_sk，再执行`./start.sh`即可。

### 浏览器访问 [http://localhost:8080](http://localhost:8080)，用户名 admin，密码 123456

## 目录结构

`application` : go gin + fabric-sdk-go 调用链码，提供外部访问接口，前端编译后静态资源放在`dist`目录下

`chaincode` : go 编写的智能合约

`deploy` : 区块链网络配置

`vue` : vue + element的前端展示页面

## 模块展示


### 注册模块（新建用户写入账本），127.0.0.1:8000/api/v1/createWyuUser

post方法，raw-json格式

```
{"wyuUserId":"4444433333","wyuUserName":"测试员","wyuPasswd":"zp43"}
```

### 登录模块（将账本数据读出并按条件查询），127.0.0.1:8000/api/v1/queryWyuUser

post方法，raw-json格式

```
key:wyuUserId;value:4444433333
key:wyuPasswd;value:zp43
```

### 助学金评定模块（用户写入各项数据后合约算法计算得出助学金评分，并写入账本），127.0.0.1:8000/api/v1/createProofMaterial

post方法，raw-json格式
内部测试

```
{"stipendId":"3118002204","annualHouseholdIncome":3500,"comprehensiveEvaluation":100,"volunteerTime":100}
```

```
{"stipendId":"1234567890","annualHouseholdIncome":3700,"comprehensiveEvaluation":110,"volunteerTime":100}
```

演示测试

```
{"stipendId":"4444433333","annualHouseholdIncome":4000,"comprehensiveEvaluation":100,"volunteerTime":100}
```

### 助学金申请信息查询模块（读出账本数据，查询所有助学金申请人的各项填入信息），127.0.0.1:8000/api/v1/queryProofMaterial

不需要输入

### 助学金申请信息修改模块 （读出账本中匹配的用户数据，删除账本原数据，新写入数据）127.0.0.1:8000/api/v1/updateProofMaterial

```
{"stipendId":"4444433333","annualHouseholdIncome":2333,"comprehensiveEvaluation":4,"volunteerTime":50}
```

### 助学金申请资料（文件）模块 (传入ipfs系统中) 127.0.0.1:8000/api/v1/createPhotoMaterial

post方法，data-from格式

### 助学金排序模块 127.0.0.1:8000/api/v1/createStipendRanking

不需要输入

```
{"stipendId":"1234567890","stipendScore":70,"ranking":2}
```

queryStipendRanking
