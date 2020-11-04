local redis = require "redis-util"

-- 指定 redis 环境变量 设定默认值
local redisIp = "127.0.0.1"
local redisPort = 6379
local redisPass = nil
local redisDbIndex = 0

local reidsUri = ngx.var.uri
local reidsKey = string.sub(reidsUri,2,string.len(reidsUri))
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
--ngx.say(resp)
if  err  then
    ngx.say("11 ")
    ngx.exit(400)
end
--ngx.exit(400)
ngx.redirect(resp, 301)