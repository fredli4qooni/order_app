<template>
    <div class="register-container">
        <div class="register-box">
            <h2>Create Your Account</h2>
            <p>Join the IT Order App</p>
            <form @submit.prevent="handleRegister">
                <div class="input-group">
                    <label for="name">Full Name</label>
                    <input type="text" id="name" v-model="form.name" required />
                </div>
                <div class="input-group">
                    <label for="email">Email</label>
                    <input type="email" id="email" v-model="form.email" required />
                </div>
                <div class="input-group">
                    <label for="password">Password</label>
                    <input type="password" id="password" v-model="form.password" required />
                </div>
                <button type="submit" class="register-button" :disabled="isLoading">
                    {{ isLoading ? 'Registering...' : 'Register' }}
                </button>
            </form>
            <div class="login-link">
                <p>
                    Already have an account?
                    <router-link to="/login">Sign In</router-link>
                </p>
            </div>
        </div>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue';
import { useAuthStore } from '../store/auth';

const authStore = useAuthStore();
const isLoading = ref(false);

const form = reactive({
    name: '',
    email: '',
    password: '',
});

const handleRegister = async () => {
    isLoading.value = true;
    try {
        await authStore.register(form);
    } catch (error) {
    } finally {
        isLoading.value = false;
    }
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
  background-color: var(--color-primary);
  color: var(--color-text-light);
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.register-button:hover {
  background-color: var(--color-primary-hover); 
}

.register-button:disabled {
  background-color: #aaa;
  cursor: not-allowed;
}

.register-link, .login-link {
  margin-top: 20px;
  font-size: 0.9rem;
}

.register-link a, .login-link a {
  color: var(--color-primary);
  text-decoration: none;
  font-weight: bold;
}

.register-link a:hover, .login-link a:hover {
  color: var(--color-primary-hover);
}
</style>