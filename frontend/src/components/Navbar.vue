<template>
  <nav class="navbar">
    <div class="navbar-left">
      <div class="navbar-brand">
        <router-link to="/">IT Order App</router-link>
      </div>
      <div class="nav-links">
        <template v-if="authStore.isAdmin">
          <router-link to="/admin">Admin Dashboard</router-link>
        </template>
        <template v-else>
          <router-link to="/dashboard">My Orders</router-link>
          <router-link to="/orders/create">New Order</router-link>
        </template>
      </div>
    </div>

    <div class="navbar-right">
      <div class="user-info">
        <span class="user-name">{{ authStore.userName }}</span>
        <span class="user-role">{{ authStore.userRole }}</span>
      </div>
      <button @click="handleLogout" class="logout-button">Logout</button>
    </div>
  </nav>
</template>

<script setup>
import { useAuthStore } from '../store/auth';
import { RouterLink } from 'vue-router';

const authStore = useAuthStore();

const handleLogout = () => {
  authStore.logout();
};
</script>

<style scoped>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 2rem;
  height: 60px;
  background-color: var(--color-background-card);
  color: var(--color-text-primary);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 1000;
}

.navbar-left,
.navbar-right {
  display: flex;
  align-items: center;
  gap: 2rem;
}

.navbar-brand a {
  font-size: 1.5rem;
  font-weight: bold;
  color: var(--color-primary);
  text-decoration: none;
}

.nav-links {
  display: flex;
  gap: 1.5rem;
}

.nav-links a {
  text-decoration: none;
  color: var(--color-text-secondary);
  font-weight: 500;
  padding: 8px 0;
  border-bottom: 2px solid transparent;
  transition: color 0.3s, border-color 0.3s;
}

.nav-links a:hover {
  color: var(--color-primary);
}

/* Gaya untuk link yang sedang aktif */
.nav-links a.router-link-exact-active {
  color: var(--color-primary);
  border-bottom-color: var(--color-primary);
}

.user-info {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  line-height: 1.2;
}

.user-name {
  font-weight: bold;
}

.user-role {
  font-size: 0.8rem;
  color: #6c757d;
  text-transform: capitalize;
}

.logout-button {
  background: none;
  border: 1px solid var(--color-danger);
  color: var(--color-danger);
  padding: 0.5rem 1rem;
  border-radius: 20px;
  cursor: pointer;
  font-weight: bold;
  transition: background-color 0.3s, color 0.3s;
}

.logout-button:hover {
  background-color: var(--color-danger);
  color: var(--color-text-light);
}
</style>