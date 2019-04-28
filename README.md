版本要求: go version go1.12.1
关闭程序: kill -3 pid

部署步骤:

设置环境变量
vim /etc/profile
添加
export DELIVERY_CONFIG="/Users/homework/data/work/gowork/queueup/conf"

修改install文件:
DELIVERY_ETC="配置文件目录"

运行:
./install 编译代码

./bin/message& 运行启动程序


测试：
    //用户获取进入游戏地址
    for i in {1..10};do 
        curl -d"post=5555" "http://127.0.0.1:8080/user/offonline?userId=s1_$i"
    done
    //用户离线
    curl -d"post=5555" 127.0.0.1:8080/user/offonline?userId=s1_1

