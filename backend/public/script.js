let token = null;
let loginForm = document.getElementById('login-form');
let signupForm = document.getElementById('signup-form');
let dashboardForm = document.getElementById('dashboard-form');
let createBtn = document.getElementById('create-btn');
let dashboardsList = document.getElementById('dashboards-list');
let dashboardsHeader = document.getElementById('dashboards-header');

function showLogin() {
    loginForm.classList.remove('hidden');
    signupForm.classList.add('hidden');
    dashboardForm.classList.add('hidden');
}

function showSignup() {
    loginForm.classList.add('hidden');
    signupForm.classList.remove('hidden');
    dashboardForm.classList.add('hidden');
}

function login(e) {
    e.preventDefault();
    let username = document.getElementById('username').value;
    let password = document.getElementById('password').value;
    axios.post('http://localhost:8080/api/v1/auth/login', {username, password})
        .then(response => {
            token = response.data.token;
            localStorage.setItem('token', token);
            alert('Logged in successfully!');
            showDashboard();
        })
        .catch(error => {
            alert('Invalid credentials');
        });
}

function signup(e) {
    e.preventDefault();
    let username = document.getElementById('signup-username').value;
    let email = document.getElementById('signup-email').value;
    let name = document.getElementById('signup-name').value;
    let password = document.getElementById('signup-password').value;
    axios.post('http://localhost:8080/api/v1/auth/signup', {username, email, name, password})
        .then(response => {
            alert('Signed up successfully!');
            showDashboard();
        })
        .catch(error => {
            alert('Invalid data provided');
        });
}

function createDashboard() {
    axios.post('http://localhost:8080/api/v1/dashboards', {
        name: 'New Dashboard',
    }, {
        headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
        }
    })
        .then(response => {
            alert('Dashboard created successfully!');
            showDashboard();
        })
        .catch(error => {
            alert('Invalid data provided');
        });
}

function showDashboard() {
    loginForm.classList.add('hidden');
    signupForm.classList.add('hidden');
    dashboardForm.classList.remove('hidden');
    axios.get('http://localhost:8080/api/v1/dashboards')
        .then(response => {
            let dashboards = response.data.data;
            dashboards.forEach(dashboard => {
                axios.get(`http://localhost:8080/api/v1/dashboards/${dashboard.id}`)
                    .then(response => {
                        let dashboardDetails = response.data.data;
                        dashboardsList.innerHTML += `
                            <div class="flex mb-2 border-b border-gray-200">
                                <div class="w-1/2">${dashboardDetails.name}</div>
                                <div class="w-1/2 flex justify-end">
                                    <button class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded" onclick="deleteDashboard(${dashboard.id})">Delete</button>
                                </div>
                            </div>
                        `;
                    })
                    .catch(error => {
                        alert('Unauthorized access');
                    });
            });
        })
        .catch(error => {
            alert('Unauthorized access');
        });
}

function deleteDashboard(id) {
    axios.delete(`http://localhost:8080/api/v1/dashboards/${id}`, {
        headers: {
            Authorization: `Bearer ${localStorage.getItem('token')}`,
        }
    })
        .then(response => {
            alert('Dashboard deleted successfully!');
            showDashboard();
        })
        .catch(error => {
            alert('Unauthorized access');
        });
}