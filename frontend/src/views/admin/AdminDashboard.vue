<template>
  <div>
    <Navbar />
    <main class="dashboard-container">
      <!-- Page Header -->
      <div class="page-header">
        <div class="header-content">
          <h1 class="page-title">
            <span class="title-icon">⚡</span>
            Admin Dashboard
          </h1>
          <p class="page-subtitle">Manage and monitor all orders in your system</p>
        </div>
        <div class="header-actions">
          <button class="refresh-btn" @click="refreshData" :disabled="orderStore.isLoading">
            <span class="btn-icon">🔄</span>
            Refresh
          </button>
        </div>
      </div>

      <!-- Dashboard Stats -->
      <DashboardStats />

      <!-- Orders Section -->
      <div class="orders-section">
        <div class="section-header">
          <h2 class="section-title">
            <span class="section-icon">📋</span>
            All Orders
          </h2>
          <div class="section-actions">
            <div class="search-box">
              <input 
                type="text" 
                placeholder="Search orders..." 
                v-model="searchQuery"
                class="search-input"
              >
              <span class="search-icon">🔍</span>
            </div>
            <select v-model="filterStatus" class="filter-select">
              <option value="">All Statuses</option>
              <option value="Pending">Pending</option>
              <option value="In Progress">In Progress</option>
              <option value="Done">Done</option>
            </select>
          </div>
        </div>

        <!-- Loading State -->
        <div v-if="orderStore.isLoading" class="loading-container">
          <div class="loading-spinner"></div>
          <p class="loading-text">Loading all orders...</p>
        </div>

        <!-- Error State -->
        <div v-else-if="orderStore.error" class="error-container">
          <div class="error-icon">⚠️</div>
          <p class="error-text">{{ orderStore.error }}</p>
          <button class="retry-btn" @click="refreshData">Try Again</button>
        </div>

        <!-- Orders Table -->
        <div v-else-if="filteredOrders.length > 0" class="table-container">
          <div class="table-wrapper">
            <table class="orders-table">
              <thead>
                <tr>
                  <th class="sortable" @click="sortBy('id')">
                    <div class="th-content">
                      ID
                      <span class="sort-indicator" :class="getSortClass('id')">↕️</span>
                    </div>
                  </th>
                  <th class="sortable" @click="sortBy('user')">
                    <div class="th-content">
                      User
                      <span class="sort-indicator" :class="getSortClass('user')">↕️</span>
                    </div>
                  </th>
                  <th class="sortable" @click="sortBy('title')">
                    <div class="th-content">
                      Title
                      <span class="sort-indicator" :class="getSortClass('title')">↕️</span>
                    </div>
                  </th>
                  <th class="sortable" @click="sortBy('category')">
                    <div class="th-content">
                      Category
                      <span class="sort-indicator" :class="getSortClass('category')">↕️</span>
                    </div>
                  </th>
                  <th class="sortable" @click="sortBy('created_at')">
                    <div class="th-content">
                      Created At
                      <span class="sort-indicator" :class="getSortClass('created_at')">↕️</span>
                    </div>
                  </th>
                  <th class="sortable" @click="sortBy('status')">
                    <div class="th-content">
                      Status
                      <span class="sort-indicator" :class="getSortClass('status')">↕️</span>
                    </div>
                  </th>
                  <th>Action</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="order in paginatedOrders" :key="order.id" class="table-row">
                  <td class="id-cell">
                    <span class="order-id">#{{ order.id }}</span>
                  </td>
                  <td class="user-cell">
                    <div class="user-info">
                      <div class="user-avatar">
                        {{ getUserInitials(order.User.name) }}
                      </div>
                      <div class="user-details">
                        <div class="user-name">{{ order.User.name }}</div>
                        <div class="user-email">{{ order.User.email }}</div>
                      </div>
                    </div>
                  </td>
                  <td class="title-cell">
                    <div class="order-title" :title="order.title">
                      {{ truncateText(order.title, 40) }}
                    </div>
                  </td>
                  <td class="category-cell">
                    <span class="category-badge">{{ order.category }}</span>
                  </td>
                  <td class="date-cell">
                    <div class="date-info">
                      <div class="date-primary">{{ formatDate(order.created_at) }}</div>
                      <div class="date-secondary">{{ formatTime(order.created_at) }}</div>
                    </div>
                  </td>
                  <td class="status-cell">
                    <span :class="['status-badge', getStatusClass(order.status)]">
                      <span class="status-dot"></span>
                      {{ order.status }}
                    </span>
                  </td>
                  <td class="action-cell">
                    <select 
                      :value="order.status" 
                      @change="updateStatus(order.id, $event.target.value)" 
                      class="status-select"
                      :class="getStatusClass(order.status)"
                    >
                      <option value="Pending">Pending</option>
                      <option value="In Progress">In Progress</option>
                      <option value="Done">Done</option>
                    </select>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <!-- Pagination -->
          <div class="pagination-container" v-if="totalPages > 1">
            <div class="pagination-info">
              Showing {{ ((currentPage - 1) * itemsPerPage) + 1 }} to {{ Math.min(currentPage * itemsPerPage, filteredOrders.length) }} of {{ filteredOrders.length }} orders
            </div>
            <div class="pagination-controls">
              <button 
                class="pagination-btn" 
                @click="currentPage--" 
                :disabled="currentPage === 1"
              >
                ← Previous
              </button>
              <span class="pagination-numbers">
                <button 
                  v-for="page in visiblePages" 
                  :key="page"
                  class="pagination-number"
                  :class="{ active: page === currentPage }"
                  @click="currentPage = page"
                >
                  {{ page }}
                </button>
              </span>
              <button 
                class="pagination-btn" 
                @click="currentPage++" 
                :disabled="currentPage === totalPages"
              >
                Next →
              </button>
            </div>
          </div>
        </div>

        <!-- Empty State -->
        <div v-else class="empty-state">
          <div class="empty-icon">📭</div>
          <h3 class="empty-title">No Orders Found</h3>
          <p class="empty-text">
            <span v-if="searchQuery || filterStatus">
              Try adjusting your search or filter criteria.
            </span>
            <span v-else>
              There are no orders in the system yet.
            </span>
          </p>
          <button v-if="searchQuery || filterStatus" class="clear-filters-btn" @click="clearFilters">
            Clear Filters
          </button>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { onMounted, ref, computed } from 'vue';
import Navbar from '../../components/Navbar.vue';
import { useOrderStore } from '../../store/order';
import DashboardStats from './DashboardStats.vue';

const orderStore = useOrderStore();

// Reactive data
const searchQuery = ref('');
const filterStatus = ref('');
const sortField = ref('created_at');
const sortDirection = ref('desc');
const currentPage = ref(1);
const itemsPerPage = ref(10);

onMounted(() => {
  orderStore.fetchAllOrders();
});

// Methods
const refreshData = () => {
  orderStore.fetchAllOrders();
};

const updateStatus = (orderId, newStatus) => {
  if (confirm(`Are you sure you want to change status to "${newStatus}"?`)) {
    orderStore.updateOrderStatus(orderId, newStatus);
  }
};

const getStatusClass = (status) => {
  const classes = {
    'Pending': 'status-pending',
    'In Progress': 'status-progress',
    'Done': 'status-done'
  };
  return classes[status] || '';
};

const getUserInitials = (name) => {
  return name
    .split(' ')
    .map(word => word.charAt(0))
    .join('')
    .toUpperCase()
    .substring(0, 2);
};

const truncateText = (text, maxLength) => {
  return text.length > maxLength ? text.substring(0, maxLength) + '...' : text;
};

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

const sortBy = (field) => {
  if (sortField.value === field) {
    sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc';
  } else {
    sortField.value = field;
    sortDirection.value = 'asc';
  }
  currentPage.value = 1;
};

const getSortClass = (field) => {
  if (sortField.value !== field) return '';
  return sortDirection.value === 'asc' ? 'sort-asc' : 'sort-desc';
};

const clearFilters = () => {
  searchQuery.value = '';
  filterStatus.value = '';
  currentPage.value = 1;
};

// Computed properties
const filteredOrders = computed(() => {
  let filtered = orderStore.orders || [];

  // Search filter
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase();
    filtered = filtered.filter(order => 
      order.title.toLowerCase().includes(query) ||
      order.User.name.toLowerCase().includes(query) ||
      order.User.email.toLowerCase().includes(query) ||
      order.category.toLowerCase().includes(query) ||
      order.id.toString().includes(query)
    );
  }

  // Status filter
  if (filterStatus.value) {
    filtered = filtered.filter(order => order.status === filterStatus.value);
  }

  // Sort
  filtered.sort((a, b) => {
    let aValue, bValue;
    
    switch (sortField.value) {
      case 'user':
        aValue = a.User.name;
        bValue = b.User.name;
        break;
      case 'created_at':
        aValue = new Date(a.created_at);
        bValue = new Date(b.created_at);
        break;
      default:
        aValue = a[sortField.value];
        bValue = b[sortField.value];
    }

    if (aValue < bValue) return sortDirection.value === 'asc' ? -1 : 1;
    if (aValue > bValue) return sortDirection.value === 'asc' ? 1 : -1;
    return 0;
  });

  return filtered;
});

const totalPages = computed(() => Math.ceil(filteredOrders.value.length / itemsPerPage.value));

const paginatedOrders = computed(() => {
  const start = (currentPage.value - 1) * itemsPerPage.value;
  const end = start + itemsPerPage.value;
  return filteredOrders.value.slice(start, end);
});

const visiblePages = computed(() => {
  const pages = [];
  const total = totalPages.value;
  const current = currentPage.value;
  
  if (total <= 7) {
    for (let i = 1; i <= total; i++) {
      pages.push(i);
    }
  } else {
    if (current <= 4) {
      for (let i = 1; i <= 5; i++) pages.push(i);
      pages.push('...');
      pages.push(total);
    } else if (current >= total - 3) {
      pages.push(1);
      pages.push('...');
      for (let i = total - 4; i <= total; i++) pages.push(i);
    } else {
      pages.push(1);
      pages.push('...');
      for (let i = current - 1; i <= current + 1; i++) pages.push(i);
      pages.push('...');
      pages.push(total);
    }
  }
  
  return pages;
});
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

.refresh-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1.5rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 10px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.refresh-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 15px rgba(102, 126, 234, 0.4);
}

.refresh-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
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
  flex-wrap: wrap;
  gap: 1rem;
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

.section-actions {
  display: flex;
  gap: 1rem;
  align-items: center;
}

.search-box {
  position: relative;
}

.search-input {
  padding: 0.75rem 1rem 0.75rem 2.5rem;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  font-size: 0.9rem;
  width: 250px;
  transition: all 0.3s ease;
}

.search-input:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.search-icon {
  position: absolute;
  left: 0.75rem;
  top: 50%;
  transform: translateY(-50%);
  color: #a0aec0;
}

.filter-select {
  padding: 0.75rem 1rem;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  font-size: 0.9rem;
  background: white;
  cursor: pointer;
  transition: all 0.3s ease;
}

.filter-select:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
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
  border-top: 4px solid #667eea;
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
}

.retry-btn {
  padding: 0.75rem 1.5rem;
  background: #e53e3e;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
}

/* Table */
.table-container {
  overflow: hidden;
}

.table-wrapper {
  overflow-x: auto;
}

.orders-table {
  width: 100%;
  border-collapse: collapse;
}

.orders-table th {
  background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
  padding: 1rem;
  text-align: left;
  font-weight: 600;
  color: #495057;
  border-bottom: 2px solid #dee2e6;
  position: sticky;
  top: 0;
  z-index: 10;
}

.sortable {
  cursor: pointer;
  user-select: none;
  transition: background-color 0.2s ease;
}

.sortable:hover {
  background: linear-gradient(135deg, #e9ecef 0%, #dee2e6 100%);
}

.th-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 0.5rem;
}

.sort-indicator {
  font-size: 0.8rem;
  opacity: 0.5;
  transition: all 0.2s ease;
}

.sort-asc .sort-indicator {
  opacity: 1;
  transform: rotate(180deg);
}

.sort-desc .sort-indicator {
  opacity: 1;
  transform: rotate(0deg);
}

.table-row {
  transition: all 0.2s ease;
}

.table-row:hover {
  background-color: #f8f9fa;
}

.orders-table td {
  padding: 1rem;
  border-bottom: 1px solid #e2e8f0;
  vertical-align: middle;
}

.id-cell .order-id {
  font-weight: 600;
  color: #667eea;
  background: #f0f4ff;
  padding: 0.25rem 0.5rem;
  border-radius: 6px;
  font-size: 0.9rem;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-weight: 600;
  font-size: 0.9rem;
}

.user-name {
  font-weight: 600;
  color: #2d3748;
  font-size: 0.9rem;
}

.user-email {
  color: #718096;
  font-size: 0.8rem;
}

.order-title {
  font-weight: 500;
  color: #2d3748;
  line-height: 1.4;
}

.category-badge {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
  color: white;
  padding: 0.35rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.date-info {
  text-align: left;
}

.date-primary {
  font-weight: 500;
  color: #2d3748;
  font-size: 0.9rem;
}

.date-secondary {
  color: #718096;
  font-size: 0.8rem;
  margin-top: 0.1rem;
}

.status-badge {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  border-radius: 20px;
  font-size: 0.8rem;
  font-weight: 600;
}

.status-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.status-pending {
  background: #fff3cd;
  color: #856404;
}

.status-pending .status-dot {
  background: #ffc107;
}

.status-progress {
  background: #d1ecf1;
  color: #0c5460;
}

.status-progress .status-dot {
  background: #17a2b8;
}

.status-done {
  background: #d4edda;
  color: #155724;
}

.status-done .status-dot {
  background: #28a745;
}

.status-select {
  padding: 0.5rem 0.75rem;
  border: 2px solid #e2e8f0;
  border-radius: 8px;
  font-size: 0.85rem;
  background: white;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s ease;
}

.status-select:focus {
  outline: none;
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.status-select.status-pending {
  border-color: #ffc107;
  background: #fff3cd;
}

.status-select.status-progress {
  border-color: #17a2b8;
  background: #d1ecf1;
}

.status-select.status-done {
  border-color: #28a745;
  background: #d4edda;
}

/* Pagination */
.pagination-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 2rem;
  border-top: 1px solid #e2e8f0;
  background: #f8f9fa;
}

.pagination-info {
  color: #718096;
  font-size: 0.9rem;
}

.pagination-controls {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.pagination-btn {
  padding: 0.5rem 1rem;
  border: 2px solid #e2e8f0;
  background: white;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s ease;
}

.pagination-btn:hover:not(:disabled) {
  border-color: #667eea;
  color: #667eea;
}

.pagination-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.pagination-numbers {
  display: flex;
  gap: 0.25rem;
}

.pagination-number {
  width: 40px;
  height: 40px;
  border: 2px solid transparent;
  background: white;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 500;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.pagination-number:hover {
  border-color: #667eea;
  color: #667eea;
}

.pagination-number.active {
  background: #667eea;
  color: white;
  border-color: #667eea;
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
}

.clear-filters-btn {
  padding: 0.75rem 1.5rem;
  background: #667eea;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.2s ease;
}

.clear-filters-btn:hover {
  background: #5a67d8;
  transform: translateY(-1px);
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
  
  .section-header {
    flex-direction: column;
    align-items: stretch;
  }
  
  .section-actions {
    justify-content: center;
    flex-wrap: wrap;
  }
}

@media (max-width: 768px) {
  .search-input {
    width: 200px;
  }
  
  .user-info {
    flex-direction: column;
    gap: 0.5rem;
    text-align: center;
  }
  
  .pagination-container {
    flex-direction: column;
    gap: 1rem;
  }
  
  .pagination-numbers {
    flex-wrap: wrap;
    justify-content: center;
  }
}

@media (max-width: 640px) {
  .orders-table th,
  .orders-table td {
    padding: 0.5rem;
    font-size: 0.8rem;
  }
  
  .search-input {
    width: 150px;
  }
  
  .section-actions {
    flex-direction: column;
    gap: 0.5rem;
  }
}
</style>