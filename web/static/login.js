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
    console.log('login() function called');

    // 从表单获取数据
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;
    const resultElement = document.getElementById("loginResult");

    console.log('Email:', email);
    console.log('Password:', password);

    if (!email && !password) {
        alert('请输入邮箱和密码');
        return false;
    } else if(!email) {
        alert('请输入邮箱');
        return false;
    } else if (!password) {
        alert('请输入密码');
        return false;
    }

    // 构建请求数据
    const requestData = {
        email: email,
        password: password
    };

    console.log('Sending request to:', API_BASE + "/auth/login");
    console.log('Request data:', requestData);

    // 发送POST请求
    fetch(API_BASE + "/auth/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json"
        },
        body: JSON.stringify(requestData)
    })
        .then(response => {
            console.log('Response status:', response.status);
            console.log('Response headers:', response.headers);

            // 首先检查HTTP状态码
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }

            // 尝试解析JSON
            const contentType = response.headers.get('content-type');
            if (contentType && contentType.includes('application/json')) {
                return response.json();
            } else {
                // 如果不是JSON，获取文本
                return response.text().then(text => {
                    console.log('Non-JSON response:', text);
                    throw new Error('Server returned non-JSON response');
                });
            }
        })
        .then(result => {
            console.log('Login response:', result);

            if (isSuccess(result)) {
                // 登录成功，存储token（如果有的话）
                if (result.data && result.data.token) {
                    localStorage.setItem('authToken', result.data.token);
                    localStorage.setItem('userEmail', email);
                }
                // 跳转到首页
                console.log('Redirecting to /index');
                window.location.href = "/index";
            } else {
                // 显示错误信息
                const errorMsg = result.message || result.msg || '登录失败';
                resultElement.textContent = errorMsg;
                alert(errorMsg);
            }
        })
        .catch(err => {
            console.error('Login error:', err);

            // 更详细的错误信息
            let errorMsg = '登录请求失败';

            if (err.message.includes('Failed to fetch')) {
                errorMsg = '无法连接到服务器，请检查服务器是否运行';
            } else if (err.message.includes('non-JSON')) {
                errorMsg = '服务器返回了非JSON响应，请检查API';
            } else {
                errorMsg = err.message;
            }

            resultElement.textContent = errorMsg;
            alert(errorMsg);
        });

    return false; // 阻止表单提交
}

function clearLoginForm() {
    const emailInput = document.getElementById('email');
    const passwordInput = document.getElementById('password');
    
    if (emailInput) {
        emailInput.value = '';
        // 设置空值后再设置一次以确保清空
        emailInput.setAttribute('value', '');
    }
    
    if (passwordInput) {
        passwordInput.value = '';
        // 设置空值后再设置一次以确保清空
        passwordInput.setAttribute('value', '');
    }
    
    // 移除登出标志
    localStorage.removeItem('justLoggedOut');
}

// 页面加载完成后绑定事件
document.addEventListener('DOMContentLoaded', function () {
     // 检查是否是从登出跳转过来的
    const justLoggedOut = localStorage.getItem('justLoggedOut');
    if (justLoggedOut === 'true') {
        clearLoginForm();
    }
    
    // 额外的：如果用户已经登录，直接跳转到首页
    const logout_token = localStorage.getItem('authToken');
    if (logout_token && window.location.pathname === '/login') {
        window.location.href = '/index';
    }

    console.log('DOM loaded, initializing login form');

    const loginForm = document.getElementById('loginForm');
    if (loginForm) {
        console.log('Found login form');
        loginForm.addEventListener('submit', function (e) {
            console.log('Form submit event triggered');
            e.preventDefault(); // 阻止默认表单提交
            login();
        });
    } else {
        console.error('Login form not found! Check the HTML.');
    }

    // 为按钮添加点击事件作为备用
    const loginButton = document.querySelector('.btn-login');
    if (loginButton) {
        console.log('Found login button');
        loginButton.addEventListener('click', function (e) {
            console.log('Button click event triggered');
            e.preventDefault();
            login();
        });
    }

    // 检查用户是否已登录
    const token = localStorage.getItem('authToken');
    if (token) {
        console.log('User is already logged in with token');
        // 可以选择自动跳转或显示已登录状态
        // window.location.href = "/index";
    }
});
