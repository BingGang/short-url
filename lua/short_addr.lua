local redis = require "redis-util"

-- 指定 redis 环境变量 设定默认值
local redisIp = "127.0.0.1"
local redisPort = 6379
local redisPass = nil
local redisDbIndex = 0
local redisKeyWhiteIp = "unite:app:white:ips"

local reidsKey = ngx.var.uri



-- 连接 redis
local red = redis:new({
    host= redisIp,
    port= redisPort,
    db_index= redisDbIndex,
    password= redisPass,
    timeout=1000,
    keepalive=60000,
    pool_size=100
});

-- 验证 redis 是否正常
if red == nil then
    ngx.status = 403
    ngx.say("403 Request is limited (redis is not ok) ")
    ngx.exit(403)
end

-- 验证当前 redis-key 是否存在
local hasKey = red:exists(reidsKey);
if tonumber(hasKey) ~= 1 then
    ngx.say(" ")
    ngx.exit(400)
end
-- 获取 redis 的值
local resp, err = red:get(reidsKey) 
if  err  then
    ngx.say(" ")
    ngx.exit(400)
end
ngx.redirect(resp, 301)

