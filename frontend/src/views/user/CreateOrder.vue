<template>
  <div>
    <Navbar />
    <main class="container">
      <h1>Create New IT Order</h1>
      <form @submit.prevent="handleSubmit" class="order-form">
        <div class="form-group">
          <label for="title">Title</label>
          <input type="text" id="title" v-model="form.title" placeholder="e.g., Fix Laptop Keyboard" required />
        </div>
        <div class="form-group">
          <label for="category">Category</label>
          <select id="category" v-model="form.category" required>
            <option disabled value="">Please select one</option>
            <option>Hardware</option>
            <option>Software</option>
            <option>Network</option>
            <option>Other</option>
          </select>
        </div>
        <div class="form-group">
          <label for="description">Description</label>
          <textarea id="description" v-model="form.description" rows="5" placeholder="Provide a detailed description of your request..."></textarea>
        </div>

        <!-- Menampilkan pesan error dari store -->
        <div v-if="orderStore.error" class="error-message">
          {{ orderStore.error }}
        </div>

        <div class="form-actions">
          <button type="submit" :disabled="orderStore.isLoading" class="submit-button">
            {{ orderStore.isLoading ? 'Submitting...' : 'Submit Order' }}
          </button>
          <router-link to="/dashboard" class="cancel-button">Cancel</router-link>
        </div>
      </form>
    </main>
  </div>
</template>

<script setup>
import { reactive } from 'vue';
import Navbar from '../../components/Navbar.vue';
import { useOrderStore } from '../../store/order';

const orderStore = useOrderStore();

const form = reactive({
  title: '',
  description: '',
  category: '',
});

const handleSubmit = async () => {
  try {
    await orderStore.createOrder(form);
    alert('Order created successfully!');
  } catch (error) {
    console.log('Component caught an error during order creation.');
  }
};
</script>

<style scoped>
.container {
  padding: 2rem;
  max-width: 700px;
  margin: 0 auto;
}
.order-form {
  margin-top: 1.5rem;
  padding: 2rem;
  background-color: #f9f9f9;
  border: 1px solid #eee;
  border-radius: 8px;
}
.form-group {
  margin-bottom: 1.5rem;
}
.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: bold;
}
.form-group input,
.form-group select,
.form-group textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}
.form-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1.5rem;
}
.submit-button, .cancel-button {
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  text-decoration: none;
  font-size: 1rem;
  text-align: center;
}
.submit-button {
  background-color: #28a745;
  color: white;
}
.submit-button:disabled {
  background-color: #aaa;
  cursor: not-allowed;
}
.cancel-button {
  background-color: #6c757d;
  color: white;
}
.error-message {
  color: #dc3545;
  margin-top: 1rem;
  padding: 0.75rem;
  background-color: #f8d7da;
  border: 1px solid #f5c6cb;
  border-radius: 4px;
}
</style>