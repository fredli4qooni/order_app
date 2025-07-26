<template>
  <div class="login-container">
    <div class="login-box">
      <h2>IT Order App</h2>
      <p>Please sign in to continue</p>
      <form @submit.prevent="handleLogin">
        <div class="input-group">
          <label for="email">Email</label>
          <input type="email" id="email" v-model="email" required />
        </div>
        <div class="input-group">
          <label for="password">Password</label>
          <input type="password" id="password" v-model="password" required />
        </div>
        <button type="submit" class="login-button">Login</button>
      </form>
      <div class="register-link">
        <p>
          Don't have an account?
          <router-link to="/register">Register here</router-link>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useAuthStore } from '../store/auth';

const authStore = useAuthStore();
const email = ref('');
const password = ref('');

const handleLogin = async () => {
  await authStore.login(email.value, password.value);
};
</script>

<style scoped>
.login-container, .register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: var(--color-background); 
}

.login-box, .register-box {
  padding: 40px;
  background-color: var(--color-background-card); 
  border-radius: 8px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
  text-align: center;
}

.input-group {
  margin-bottom: 20px;
  text-align: left;
}

label {
  display: block;
  margin-bottom: 5px;
  color: #555;
}

input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  box-sizing: border-box;
}

.login-button {
  width: 100%;
  padding: 12px;
  border: none;
  border-radius: 4px;
  background-color: var(--color-primary); 
  color: var(--color-text-light);
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.login-button:hover {
  background-color: var(--color-primary-hover); 
}

.register-button {
  width: 100%;
  padding: 12px;
  border: none;
  border-radius: 4px;
  background-color: var(--color-success); 
  color: var(--color-text-light);
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.register-button:hover {
  background-color: #218838;
}

.register-link {
  margin-top: 20px;
  font-size: 0.9rem;
}

.register-link a {
  color: var(--color-primary); 
  text-decoration: none;
  font-weight: bold;
}

.register-link a:hover {
  color: var(--color-primary-hover);
}
</style>