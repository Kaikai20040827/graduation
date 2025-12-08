document.addEventListener('DOMContentLoaded', () => {
    const sideMenu = document.querySelector("aside");
    const profileBtn = document.querySelector("#profile-btn");
    const themeToggler = document.querySelector(".theme-toggler");
    const nextDay = document.getElementById('nextDay');
    const prevDay = document.getElementById('prevDay');

    // Profile button toggle for side menu
    profileBtn.onclick = function () {
        sideMenu.classList.toggle('active');
    }

    // Scroll event to remove side menu and add/remove header active class
    window.onscroll = () => {
        sideMenu.classList.remove('active');
        if (window.scrollY > 0) {
            document.querySelector('header').classList.add('active');
        } else {
            document.querySelector('header').classList.remove('active');
        }
    }

    // Theme toggle function
    const applySavedTheme = () => {
        const isDarkMode = localStorage.getItem('dark-theme') === 'true';
        if (isDarkMode) {
            document.body.classList.add('dark-theme');
            themeToggler.querySelector('span:nth-child(2)').classList.add('active');
            themeToggler.querySelector('span:nth-child(1)').classList.remove('active');
        } else {
            document.body.classList.remove('dark-theme');
            themeToggler.querySelector('span:nth-child(2)').classList.remove('active');
            themeToggler.querySelector('span:nth-child(1)').classList.add('active');
        }
    }

    // Set the initial theme based on localStorage
    applySavedTheme();

    // Toggle theme function
    themeToggler.onclick = function () {
        // Toggle dark theme class on body
        document.body.classList.toggle('dark-theme');

        // Toggle active class on the theme toggler spans
        themeToggler.querySelector('span:nth-child(1)').classList.toggle('active');
        themeToggler.querySelector('span:nth-child(2)').classList.toggle('active');

        // Save the theme preference in localStorage
        localStorage.setItem('dark-theme', document.body.classList.contains('dark-theme'));
    }

    // Function to set timetable data
    let setData = (day) => {
        document.querySelector('table tbody').innerHTML = '';  // Clear previous table data
        let daylist = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
        document.querySelector('.timetable div h2').innerHTML = daylist[day];

        // Define subjects for each day (you might need to update this with real data)
        let daySchedule = [];
        switch (day) {
            case 0: daySchedule = Sunday; break;
            case 1: daySchedule = Monday; break;
            case 2: daySchedule = Tuesday; break;
            case 3: daySchedule = Wednesday; break;
            case 4: daySchedule = Thursday; break;
            case 5: daySchedule = Friday; break;
            case 6: daySchedule = Saturday; break;
        }

        // Append timetable data to table
        daySchedule.forEach(sub => {
            const tr = document.createElement('tr');
            const trContent = `
                <td>${sub.time}</td>
                <td>${sub.roomNumber}</td>
                <td>${sub.subject}</td>
                <td>${sub.type}</td>
            `;
            tr.innerHTML = trContent;
            document.querySelector('table tbody').appendChild(tr);
        });
    }

    // Get current day and set timetable on page load
    let now = new Date();
    let today = now.getDay();  // Get current day (0 - 6)
    let day = today;  // To prevent today value from changing

    // Function to toggle timetable visibility
    function timeTableAll() {
        document.getElementById('timetable').classList.toggle('active');
        setData(today);
        document.querySelector('.timetable div h2').innerHTML = "Today's Timetable";
    }

    // Event listeners for next and previous day buttons
    nextDay.onclick = function () {
        day <= 5 ? day++ : day = 0;  // If-else one-liner
        setData(day);
    }

    prevDay.onclick = function () {
        day >= 1 ? day-- : day = 6;  // Move to previous day
        setData(day);
    }

    // Set data on page load
    setData(day);
    document.querySelector('.timetable div h2').innerHTML = "Today's Timetable";  // Set heading on load
});

// Base API URL (points to v1)
const API_BASE = 'http://127.0.0.1:8080/api/v1'
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
    const data = {
        username: document.getElementById("userid") && document.getElementById("userid").value,
        password: document.getElementById("password") && document.getElementById("password").value
    };

    fetch(API_BASE + "/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data)
    })
        .then(res => res.json())
        .then(result => {
            document.getElementById("loginResult").innerText = JSON.stringify(result);
            if (isSuccess(result) && result.data && result.data.token) {
                // save token for later use
                localStorage.setItem('token', result.data.token)
                window.location.href = '/index'
            } else {
                alert(result.message || '登录失败')
            }
        })
        .catch(err => {
            console.error(err)
            alert('登录请求失败')
        })
}

// Public file upload (uses public endpoint added to backend)
document.addEventListener('DOMContentLoaded', () => {
    const uploadBtn = document.getElementById('uploadBtn');
    const fileInput = document.getElementById('fileInput');
    const fileDesc = document.getElementById('fileDesc');
    const uploadResult = document.getElementById('uploadResult');

    if (!uploadBtn) return;

    uploadBtn.addEventListener('click', () => {
        uploadResult.innerText = '';
        if (!fileInput || fileInput.files.length === 0) {
            uploadResult.innerText = '请选择一个文件后再上传';
            return;
        }
        const fd = new FormData();
        fd.append('file', fileInput.files[0]);
        fd.append('description', fileDesc ? fileDesc.value : '');

        fetch(API_BASE + '/files/public/upload', {
            method: 'POST',
            body: fd,
        })
            .then(res => res.json())
            .then(data => {
                if (!isSuccess(data)) {
                    uploadResult.innerText = data.message || JSON.stringify(data);
                } else if (data && data.data) {
                    uploadResult.innerText = '上传成功：' + (data.data.filename || JSON.stringify(data.data));
                } else {
                    uploadResult.innerText = '上传成功';
                }
            })
            .catch(err => {
                uploadResult.innerText = '上传失败：' + err;
            });
    });
});
