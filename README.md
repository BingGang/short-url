# short_url
##基于golang +openresty+lua +redis 短链接服务

run.sh 测试启动脚本 
curl 测试：
    curl -d 'req_url=url&ex=123' -X POST http://127.0.0.1:17681/api/1/set/short/url
    
nginx.d nginx配置文件

 修改nginx.conf lua文件路径

nginx 配置 文件引入配置文件
include './nginx.d/nginx.conf'



