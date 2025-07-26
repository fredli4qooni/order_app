<template>
  <div class="not-found-container">
    <div class="not-found-content">
      <h1>404</h1>
      <h2>Page Not Found</h2>
      <p>Sorry, the page you are looking for does not exist.</p>
      <router-link to="/" class="home-button">Go to Homepage</router-link>
    </div>
  </div>
</template>

<script setup>
</script>

<style scoped>
.not-found-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  text-align: center;
  background-color: #f8f9fa;
  color: #343a40;
  font-family: Arial, sans-serif;
}

.not-found-content h1 {
  font-size: 10rem;
  margin: 0;
  font-weight: 900;
  color: #dc3545;
}

.not-found-content h2 {
  font-size: 2rem;
  margin-top: -20px;
  margin-bottom: 20px;
}

.not-found-content p {
  font-size: 1.2rem;
  margin-bottom: 30px;
}

.home-button {
  display: inline-block;
  padding: 10px 20px;
  background-color: #007bff;
  color: white;
  text-decoration: none;
  border-radius: 5px;
  transition: background-color 0.3s;
}

.home-button:hover {
  background-color: #0056b3;
}
</style>