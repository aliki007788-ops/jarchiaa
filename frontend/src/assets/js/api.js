const API = {
baseURL: '/api/v1',
token: localStorage.getItem('jarchia_token'),

async request(endpoint, options = {}) {
const url = this.baseURL + endpoint;
const headers = {
'Content-Type': 'application/json',
...options.headers
};

text
if (this.token) {
  headers['Authorization'] = `Bearer ${this.token}`;
}

const response = await fetch(url, { ...options, headers });
return response.json();
},

async get(endpoint) {
return this.request(endpoint, { method: 'GET' });
},

async post(endpoint, data) {
return this.request(endpoint, {
method: 'POST',
body: JSON.stringify(data)
});
},

async put(endpoint, data) {
return this.request(endpoint, {
method: 'PUT',
body: JSON.stringify(data)
});
},

async delete(endpoint) {
return this.request(endpoint, { method: 'DELETE' });
}
};

window.API = API;


