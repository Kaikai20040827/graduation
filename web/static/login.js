const API_BASE = '/api/v1'

function isSuccess(resp) {
    return resp && (resp.code === 0 || resp.code === 200)
}

function register() {
    const email = document.getElementById("email") && document.getElementById("email").value
    const username = document.getElementById("userid") && document.getElementById("userid").value
    const password = document.getElementById("password") && document.getElementById("password").value
    const confirmed_password = document.getElementById("confirm") && document.getElementById("confirm").value

    if (!email || !username || !password) {
        alert('请填写所有必填项');
        return;
    }
    if (password !== confirmed_password) {
        alert("两次密码输入不一致，请重新输入！");
        return;
    }

    const data = { email, username, password, confirmed_password }

    fetch(API_BASE + "/auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data)
    })
        .then(res => res.json())
        .then(result => {
            document.getElementById("regResult").innerText = JSON.stringify(result);
            if (isSuccess(result)) {
                // registration success
                window.location.href = "/register_result"
            } else {
                alert(result.message || '注册失败')
            }
        })
        .catch(err => {
            console.error(err)
            alert('注册请求失败')
        })

}

function login() {
    // 从表单获取数据
    const email = document.getElementById("email") && document.getElementById("email").value
    const password = document.getElementById("password") && document.getElementById("password").value

    if (!email || !password) {
        alert('请输入邮箱和密码');
        return false;
    }

    fetch(API_BASE + "/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password })
    })
        .then(res => res.json())
        .then(result => {
            document.getElementById("loginResult").innerText = JSON.stringify(result);
            if (isSuccess(result)) {
                // registration success
                window.location.href = "/index"
            } else {
                alert(result.message || '登录失败')
            }
        })
        .catch(err => {
            console.error(err)
            console.log(API_BASE + "/auth/login")
            alert('登录请求失败')
        })
}
