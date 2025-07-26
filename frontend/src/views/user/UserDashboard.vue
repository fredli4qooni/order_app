<template>
  <div>
    <Navbar />
    <main class="dashboard-container">
      <!-- Page Header -->
      <div class="page-header">
        <div class="header-content">
          <h1 class="page-title">
            <span class="title-icon">📋</span>
            My IT Orders
          </h1>
          <p class="page-subtitle">Track and manage your IT service requests</p>
        </div>
        <div class="header-actions">
          <router-link to="/orders/create" class="create-order-btn">
            <span class="btn-icon">✨</span>
            New Order
          </router-link>
        </div>
      </div>

      <!-- User Stats -->
      <UserStats />

      <!-- Orders Section -->
      <div class="orders-section">
        <div class="section-header">
          <h2 class="section-title">
            <span class="section-icon">📝</span>
            Your Orders
          </h2>
          <div class="orders-count">
            <span class="count-number">{{ orderStore.orders?.length || 0 }}</span>
            <span class="count-label">Total Orders</span>
          </div>
        </div>

        <!-- Loading State -->
        <div v-if="orderStore.isLoading" class="loading-container">
          <div class="loading-spinner"></div>
          <p class="loading-text">Loading your orders...</p>
        </div>

        <!-- Error State -->
        <div v-else-if="orderStore.error" class="error-container">
          <div class="error-icon">⚠️</div>
          <p class="error-text">{{ orderStore.error }}</p>
          <button class="retry-btn" @click="orderStore.fetchUserOrders()">Try Again</button>
        </div>

        <!-- Orders Grid -->
        <div v-else-if="orderStore.orders && orderStore.orders.length > 0" class="orders-grid">
          <div v-for="order in orderStore.orders" :key="order.id" class="order-card">
            <div class="order-header">
              <div class="order-id">
                <span class="id-badge">#{{ order.id }}</span>
              </div>
              <div class="order-date">
                <div class="date-primary">{{ formatDate(order.created_at) }}</div>
                <div class="date-secondary">{{ formatTime(order.created_at) }}</div>
              </div>
            </div>
            
            <div class="order-content">
              <h3 class="order-title">{{ order.title }}</h3>
              <div class="order-category">
                <span class="category-badge">{{ order.category }}</span>
              </div>
            </div>

            <div class="order-status">
              <div class="status-header">
                <span class="status-label">Status Progress</span>
              </div>
              <OrderStatusTracker :status="order.status" />
            </div>

            <div class="order-actions">
              <button class="view-btn" @click="viewOrderDetails(order.id)">
                <span class="btn-icon">👁️</span>
                View Details
              </button>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else class="empty-state">
          <div class="empty-icon">📭</div>
          <h3 class="empty-title">No Orders Yet</h3>
          <p class="empty-text">
            You haven't created any IT service requests yet. Get started by creating your first order!
          </p>
          <router-link to="/orders/create" class="empty-action-btn">
            <span class="btn-icon">➕</span>
            Create Your First Order
          </router-link>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { onMounted } from 'vue';
import Navbar from '../../components/Navbar.vue';
import { useOrderStore } from '../../store/order';
import UserStats from './UserStats.vue';
import OrderStatusTracker from '../../components/OrderStatusTracker.vue';

const orderStore = useOrderStore();

onMounted(() => {
  orderStore.fetchUserOrders();
});

// Methods
const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  });
};

const formatTime = (dateString) => {
  return new Date(dateString).toLocaleTimeString('en-US', {
    hour: '2-digit',
    minute: '2-digit'
  });
};

const viewOrderDetails = (orderId) => {
  console.log('View order details:', orderId);
};
</script>

<style scoped>
.dashboard-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
  padding: 2rem;
}

/* Page Header */
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  background: white;
  padding: 2rem;
  border-radius: 16px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
}

.header-content h1 {
  margin: 0;
  font-size: 2.5rem;
  font-weight: 700;
  color: #2d3748;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.title-icon {
  font-size: 2.2rem;
}

.page-subtitle {
  color: #718096;
  font-size: 1.1rem;
  margin: 0.5rem 0 0 0;
}

.create-order-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-weight: 600;
  text-decoration: none;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 1rem;
}

.create-order-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(79, 172, 254, 0.4);
  color: white;
}

.btn-icon {
  font-size: 1.1rem;
}

/* Orders Section */
.orders-section {
  background: white;
  border-radius: 16px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.section-header {
  padding: 2rem;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.section-title {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 600;
  color: #2d3748;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.orders-count {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.25rem;
}

.count-number {
  font-size: 2rem;
  font-weight: 700;
  color: #4facfe;
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.count-label {
  font-size: 0.8rem;
  color: #718096;
  font-weight: 500;
}

/* Loading State */
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 4rem;
}

.loading-spinner {
  width: 50px;
  height: 50px;
  border: 4px solid #e2e8f0;
  border-top: 4px solid #4facfe;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 1rem;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.loading-text {
  color: #718096;
  font-size: 1.1rem;
}

/* Error State */
.error-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 4rem;
}

.error-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
}

.error-text {
  color: #e53e3e;
  font-size: 1.1rem;
  margin-bottom: 1rem;
  text-align: center;
}

.retry-btn {
  padding: 0.75rem 1.5rem;
  background: #e53e3e;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.2s ease;
}

.retry-btn:hover {
  background: #c53030;
  transform: translateY(-1px);
}

/* Orders Grid */
.orders-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
  gap: 1.5rem;
  padding: 2rem;
}

.order-card {
  background: white;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 1.5rem;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.order-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.order-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 12px 28px rgba(0, 0, 0, 0.12);
  border-color: #4facfe;
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.id-badge {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 0.25rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.date-primary {
  font-weight: 500;
  color: #2d3748;
  font-size: 0.9rem;
  text-align: right;
}

.date-secondary {
  color: #718096;
  font-size: 0.8rem;
  text-align: right;
  margin-top: 0.1rem;
}

.order-content {
  margin-bottom: 1.5rem;
}

.order-title {
  font-size: 1.1rem;
  font-weight: 600;
  color: #2d3748;
  margin: 0 0 0.75rem 0;
  line-height: 1.4;
}

.category-badge {
  background: linear-gradient(135deg, #a8edea 0%, #fed6e3 100%);
  color: #2d3748;
  padding: 0.35rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.order-status {
  margin-bottom: 1.5rem;
}

.status-header {
  margin-bottom: 0.75rem;
}

.status-label {
  font-size: 0.9rem;
  font-weight: 500;
  color: #4a5568;
}

.order-actions {
  display: flex;
  gap: 0.75rem;
}

.view-btn {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 8px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.view-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

/* Empty State */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 4rem;
  text-align: center;
}

.empty-icon {
  font-size: 5rem;
  margin-bottom: 1rem;
}

.empty-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #2d3748;
  margin-bottom: 0.5rem;
}

.empty-text {
  color: #718096;
  font-size: 1rem;
  margin-bottom: 2rem;
  max-width: 400px;
  line-height: 1.5;
}

.empty-action-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 1rem 2rem;
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-weight: 600;
  text-decoration: none;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 1rem;
}

.empty-action-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(79, 172, 254, 0.4);
  color: white;
}

/* Responsive Design */
@media (max-width: 1024px) {
  .dashboard-container {
    padding: 1rem;
  }
  
  .page-header {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }
  
  .orders-grid {
    grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
    gap: 1rem;
    padding: 1rem;
  }
}

@media (max-width: 768px) {
  .section-header {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }
  
  .orders-grid {
    grid-template-columns: 1fr;
    padding: 1rem;
  }
  
  .order-header {
    flex-direction: column;
    gap: 0.5rem;
    align-items: flex-start;
  }
  
  .date-primary,
  .date-secondary {
    text-align: left;
  }
}

@media (max-width: 480px) {
  .page-header {
    padding: 1rem;
  }
  
  .header-content h1 {
    font-size: 2rem;
  }
  
  .orders-grid {
    padding: 0.5rem;
  }
  
  .order-card {
    padding: 1rem;
  }
}
</style>