<template>
  <div class="login">
    <h1>ورود به حساب</h1>
    <button @click="$router.push('/')">بازگشت</button>
  </div>
</template>

<script>
export default {
  name: 'Login'
}
</script>

<style scoped>
.login {
  padding: 20px;
  text-align: center;
}
h1 {
  color: #42b983;
}
button {
  padding: 10px 20px;
  background: #42b983;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}
</style>
