<template>
  <div class="home">
    <h1>به جارچیا خوش آمدید</h1>
    <p>این صفحه اصلی است</p>
    <button @click="$router.push('/login')">ورود</button>
  </div>
</template>

<script>
export default {
  name: 'Home'
}
</script>

<style scoped>
.home {
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
