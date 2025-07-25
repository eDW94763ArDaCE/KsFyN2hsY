以下是优化后的Dockerfile代码片段：

```Dockerfile
FROM alpine:3.13

# Install required packages
RUN apk --no-cache add ca-certificates curl bash jq py3-pip \
    && pip install awscli

COPY ecs-deploy /usr/local/bin/ecs-deploy
RUN chmod a+x /usr/local/bin/ecs-deploy
RUN ln -s /usr/local/bin/ecs-deploy /ecs-deploy

COPY test.bats /test.bats
COPY run-tests.sh /run-tests.sh
RUN chmod a+x /run-tests.sh

ENTRYPOINT ["ecs-deploy"]
```

优化点：
1. 合并了RUN命令，减少层数，提高构建速度。
2. 保持原有功能不变。

以下是实现登录流程的伪代码：

```javascript
// 登录流程伪代码
function login(username, password) {
  // 校验用户名和密码
  if (!username || !password) {
    return "用户名和密码不能为空";
  }

  // 调用后端接口进行登录校验
  fetch('/api/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ username, password })
  })
  .then(response => response.json())
  .then(data => {
    if (data.success) {
      // 登录成功，保存登录状态
      saveLoginState(data.token);
    } else {
      // 登录失败，提示错误信息
      return "登录失败：" + data.message;
    }
  })
  .catch(error => {
    return "登录请求失败：" + error.message;
  });
}

// 保存登录状态
function saveLoginState(token) {
  // 将token保存到localStorage或cookie中
  localStorage.setItem('token', token);
}

// 校验是否为管理员
function checkAdmin() {
  // 从localStorage或cookie中获取token
  const token = localStorage.getItem('token');
  
  // 调用后端接口校验是否为管理员
  fetch('/api/check-admin', {
    method: 'GET',
    headers: {
      'Authorization': 'Bearer ' + token
    }
  })
  .then(response => response.json())
  .then(data => {
    if (data.isAdmin) {
      return true;
    } else {
      return false;
    }
  })
  .catch(error => {
    return false;
  });
}
```

这个伪代码实现了登录流程和校验是否为管理员的功能。可以根据实际需求进行调整和完善。