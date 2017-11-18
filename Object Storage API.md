## Swift Object Storage API 测试报告

## 普通用户
- 测试账号1
    > 账号名: 1 <br>
    账号密码: 123456 <br>
    关联项目: 11111111@qq.com
    
* 请求token curl
<pre class=”brush: bash; gutter: true;”>
curl -X POST \
-d '{
    "auth": {
        "tenantName": "11111111@qq.com",
        "passwordCredentials":{
            "username": "1",
            "password": "123456"
        }
    }
}' \
-H "Content-type: application/json" \
http://172.16.92.30:5000/v2.0/tokens | python -mjson.tool
</pre>
    
* 请求返回

<pre class=”brush: bash; gutter: true;”>
"token": {
            "audit_ids": [
                "In8lAbRzR_2347LrkaiQXA"
            ],
            "expires": "2017-11-18T06:36:37Z",
            "id": "ead1720959d54ac5acfc496f8e32455a",
            "issued_at": "2017-11-18T02:36:37.876520",
            "tenant": {
                "description": "11111111@qq.com",
                "enabled": true,
                "id": "6ff3043712604a2698aacacb62572bc5",
                "name": "11111111@qq.com"
            }
        },
        "user": {
            "id": "fdc054ab82e848958b716250a00c1711",
            "name": "1",
            "roles": [
                {
                    "name": "_member_"
                }
            ],
            "roles_links": [],
            "username": "1"
        }
    }
</pre>
* 获得token ead1720959d54ac5acfc496f8e32455a<br>
利用该token 请求swift接口

### 1 Account测试
    GET /v1/{account}
>原始请求
    
    curl -i $publicURL?format=json -X GET -H "X-Auth-Token: $token"
>实际测试请求
        
    curl -i http://172.16.92.140:7480/swift/v1?format=json -X GET -H "X-Auth-Token: ead1720959d54ac5acfc496f8e32455a"

>返回数据

    HTTP/1.1 200 OK
    Content-Length: 2
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 02:50:17 GMT
    Keep-Alive: timeout=38
    X-Account-Bytes-Used: 0
    X-Account-Bytes-Used-Actual: 0
    X-Account-Container-Count: 0
    X-Account-Object-Count: 0
    X-Account-Storage-Policy-Default-Placement-Bytes-Used: 0
    X-Account-Storage-Policy-Default-Placement-Bytes-Used-Actual: 0
    X-Account-Storage-Policy-Default-Placement-Container-Count: 0
    X-Account-Storage-Policy-Default-Placement-Object-Count: 0
    X-Openstack-Request-Id: tx0000000000000000000f5-005a0f9fe8-12331-default
    X-Timestamp: 1510973417.87202
    X-Trans-Id: tx0000000000000000000f5-005a0f9fe8-12331-default
    []
    
>未有数据 若测试有数据 goto 测试容器是否创建成功 object是否创建成功

### 2 Containers测试
    PUT /v1/{account}/{container}
>原始请求
    
    curl -i $publicURL/steven -X PUT -H "Content-Length: 0" -H "X-Auth-Token: $token"
>实际测试请求(创建已有命名)
        
    curl -i http://172.16.92.140:7480/swift/v1/teng?format=json -X PUT -H "X-Auth-Token: ead1720959d54ac5acfc496f8e32455a"

>返回结果

    HTTP/1.1 409 Conflict
    Content-Length: 146
    Accept-Ranges: bytes
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 03:00:41 GMT
    Keep-Alive: timeout=38
    X-Openstack-Request-Id: tx0000000000000000000fa-005a0fa259-12331-default
    X-Trans-Id: tx0000000000000000000fa-005a0fa259-12331-default
    {"Code":"BucketAlreadyExists","BucketName":"teng","RequestId":"tx0000000000000000000fa-005a0fa259-12331-default","HostId":"12331-default-default"}
    
>实际测试请求(创建未有命名)
        
    curl -i http://172.16.92.140:7480/swift/v1/teng1?format=json -X PUT -H "X-Auth-Token: ead1720959d54ac5acfc496f8e32455a"

>返回结果
    
    HTTP/1.1 201 Created
    Content-Length: 0
    Accept-Ranges: bytes
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 03:02:27 GMT
    Keep-Alive: timeout=38
    X-Openstack-Request-Id: tx0000000000000000000fb-005a0fa2c2-12331-default
    X-Trans-Id: tx0000000000000000000fb-005a0fa2c2-12331-default
    
>查看返回 

>测试请求

    curl -i http://172.16.92.140:7480/swift/v1?format=json -X GET -H "X-Auth-Token: ead1720959d54ac5acfc496f8e32455a"
    
>code:测试容器是否创建成功 

    HTTP/1.1 200 OK
    Content-Length: 38
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 03:05:40 GMT
    Keep-Alive: timeout=38
    X-Account-Bytes-Used: 0
    X-Account-Bytes-Used-Actual: 0
    X-Account-Container-Count: 1
    X-Account-Object-Count: 0
    X-Account-Storage-Policy-Default-Placement-Bytes-Used: 0
    X-Account-Storage-Policy-Default-Placement-Bytes-Used-Actual: 0
    X-Account-Storage-Policy-Default-Placement-Container-Count: 1
    X-Account-Storage-Policy-Default-Placement-Object-Count: 0
    X-Openstack-Request-Id: tx0000000000000000000fc-005a0fa384-12331-default
    X-Timestamp: 1510974340.42674
    X-Trans-Id: tx0000000000000000000fc-005a0fa384-12331-default
    [{"name":"teng1","count":0,"bytes":0}]
    
>修改读写权限的bucket 请求
    
    curl -i http://172.16.92.140:7480/swift/v1/teng1?format=json -X PUT -H "X-Auth-Token: ead1720959d54ac5acfc496f8e32455a" -H "X-Container-Read: .r:*"

>返回成功

    HTTP/1.1 202 Accepted
    Content-Length: 0
    Accept-Ranges: bytes
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 03:20:49 GMT
    Keep-Alive: timeout=38
    X-Openstack-Request-Id: tx0000000000000000000fd-005a0fa710-12331-default
    X-Trans-Id: tx0000000000000000000fd-005a0fa710-12331-default

>查看bucket 请求

    curl -i http://172.16.92.140:7480/swift/v1/teng1?format=json -X GET -H "X-Auth-Token: ead1720959d54ac5acfc496f8e32455a"
    
>返回

    HTTP/1.1 200 OK
    Content-Length: 2
    Accept-Ranges: bytes
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 03:59:47 GMT
    Keep-Alive: timeout=38
    X-Container-Bytes-Used: 0
    X-Container-Bytes-Used-Actual: 0
    X-Container-Object-Count: 0
    X-Openstack-Request-Id: tx000000000000000000100-005a0fb033-12331-default
    X-Storage-Policy: default-placement
    X-Timestamp: 1510974146.89111
    X-Trans-Id: tx000000000000000000100-005a0fb033-12331-default
    []
 
>删除bucket
 
>原始请求

    curl -i $publicURL/steven -X DELETE -H "X-Auth-Token: $token"
    
>实际curl

    curl -i http://172.16.92.140:7480/swift/v1/teng1?format=json -X DELETE -H "X-Auth-Token: ead1720959d54ac5acfc496f8e32455a"
    
> 返回

    HTTP/1.1 401 Unauthorized
    Content-Length: 119
    Accept-Ranges: bytes
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 04:14:30 GMT
    Keep-Alive: timeout=38
    X-Openstack-Request-Id: tx000000000000000000108-005a0fb3a6-12331-default
    X-Trans-Id: tx000000000000000000108-005a0fb3a6-12331-default
    {"Code":"AccessDenied","RequestId":"tx000000000000000000108-005a0fb3a6-12331-default","HostId":"12331-default-default"}
    

### 3 Object测试

>创建object

    curl -i http://172.16.92.140:7480/swift/v1/teng1/note?format=json -X PUT -d "hello"  -H "Content-Type: text/html; charset=UTF-8" -H "X-Auth-Token: ead1720959d54ac5acfc496f8e32455a"

>返回结果

    HTTP/1.1 201 Created
    Content-Length: 0
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 04:32:07 GMT
    Etag: 5d41402abc4b2a76b9719d911017c592
    Keep-Alive: timeout=38
    Last-Modified: Sat, 18 Nov 2017 04:32:07 GMT
    X-Openstack-Request-Id: tx00000000000000000010d-005a0fb7c6-12331-default
    X-Trans-Id: tx00000000000000000010d-005a0fb7c6-12331-default
    


>若有object则不能删除bucket

    Content-Length: 186
    Accept-Ranges: bytes
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 04:33:35 GMT
    Keep-Alive: timeout=38
    X-Openstack-Request-Id: tx00000000000000000010e-005a0fb81e-12331-default
    X-Trans-Id: tx00000000000000000010e-005a0fb81e-12331-default
    {"Code":"There was a conflict when trying to complete your request.","BucketName":"teng1","RequestId":"tx00000000000000000010e-005a0fb81e-12331-default","HostId":"12331-default-default"}

>code:object是否创建成功

    HTTP/1.1 200 OK
    Content-Length: 38
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 04:37:34 GMT
    Keep-Alive: timeout=38
    X-Account-Bytes-Used: 5
    X-Account-Bytes-Used-Actual: 4096
    X-Account-Container-Count: 1
    X-Account-Object-Count: 1
    X-Account-Storage-Policy-Default-Placement-Bytes-Used: 5
    X-Account-Storage-Policy-Default-Placement-Bytes-Used-Actual: 4096
    X-Account-Storage-Policy-Default-Placement-Container-Count: 1
    X-Account-Storage-Policy-Default-Placement-Object-Count: 1
    X-Openstack-Request-Id: tx00000000000000000010f-005a0fb90e-12331-default
    X-Timestamp: 1510979854.46708
    X-Trans-Id: tx00000000000000000010f-005a0fb90e-12331-default
    [{"name":"teng1","count":1,"bytes":5}]
    
>删除object 
    
    curl -i http://172.16.92.140:7480/swift/v1/teng1/note?format=json -X DELETE -H "X-Auth-Token: ead1720959d54ac5acfc496f8e32455a"
    
## 总结

>必须给该项目的用户添加角色_member_才可生成可用token

    keystone user-role-add --tenant 11111111@qq.com --user 11 --role _member_
    
>Account和租户一一对应 
>分租户不分用户，用户只是请求token时使用
>如下
>用户1和用户2是同一租户的不同用户

>用户1请求

    HTTP/1.1 200 OK
    Content-Length: 38
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 06:19:56 GMT
    Keep-Alive: timeout=38
    X-Account-Bytes-Used: 0
    X-Account-Bytes-Used-Actual: 0
    X-Account-Container-Count: 1
    X-Account-Object-Count: 1
    X-Account-Storage-Policy-Default-Placement-Bytes-Used: 0
    X-Account-Storage-Policy-Default-Placement-Bytes-Used-Actual: 0
    X-Account-Storage-Policy-Default-Placement-Container-Count: 1
    X-Account-Storage-Policy-Default-Placement-Object-Count: 1
    X-Openstack-Request-Id: tx000000000000000000132-005a0fd10b-12331-default
    X-Timestamp: 1510985996.01058
    X-Trans-Id: tx000000000000000000132-005a0fd10b-12331-default
    [{"name":"teng1","count":1,"bytes":0}]
    
>用户2请求
    
    HTTP/1.1 200 OK
    Content-Length: 38
    Content-Type: application/json; charset=utf-8
    Date: Sat, 18 Nov 2017 06:22:32 GMT
    Keep-Alive: timeout=38
    X-Account-Bytes-Used: 0
    X-Account-Bytes-Used-Actual: 0
    X-Account-Container-Count: 1
    X-Account-Object-Count: 1
    X-Account-Storage-Policy-Default-Placement-Bytes-Used: 0
    X-Account-Storage-Policy-Default-Placement-Bytes-Used-Actual: 0
    X-Account-Storage-Policy-Default-Placement-Container-Count: 1
    X-Account-Storage-Policy-Default-Placement-Object-Count: 1
    X-Openstack-Request-Id: tx000000000000000000133-005a0fd1a8-12331-default
    X-Timestamp: 1510986152.44960
    X-Trans-Id: tx000000000000000000133-005a0fd1a8-12331-default    
    [{"name":"teng1","count":1,"bytes":0}]
    
### 对应需求文档的接口

>计量和计费接口没有

#### 官网API列表都可调用
>bucket增删接口可以调用

>创建object接口可以调用

>修改读写权限的bucket权限可以调用

>get account接口可以调用

>get bucket接口可以调用

#### 不确定
>metadata 可以调用