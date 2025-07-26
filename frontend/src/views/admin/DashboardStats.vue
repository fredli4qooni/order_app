<template>
    <div class="stats-container">
        <h2>Dashboard Overview</h2>
        <div v-if="dashboardStore.isLoading">Loading stats...</div>
        <div v-else-if="dashboardStore.error" class="error">{{ dashboardStore.error }}</div>
        <div v-else-if="dashboardStore.stats" class="stats-grid">

            <div class="stat-card">
                <h3>Total Orders</h3>
                <p class="stat-number">{{ dashboardStore.stats.total_orders }}</p>
            </div>

            <div class="stat-card">
                <h3>Orders by Status</h3>
                <ul class="stat-list">
                    <li>Pending: <span>{{ dashboardStore.stats.orders_by_status.Pending || 0 }}</span></li>
                    <li>In Progress: <span>{{ dashboardStore.stats.orders_by_status['In Progress'] || 0 }}</span></li>
                    <li>Done: <span>{{ dashboardStore.stats.orders_by_status.Done || 0 }}</span></li>
                </ul>
            </div>

            <div class="stat-card">
                <h3>Top Requester</h3>
                <p v-if="dashboardStore.stats.top_user.user_name" class="stat-text">
                    {{ dashboardStore.stats.top_user.user_name }}
                    <span class="sub-text">with {{ dashboardStore.stats.top_user.order_count }} orders</span>
                </p>
                <p v-else>No orders yet.</p>
            </div>

            <div class="stat-card">
                <h3>Avg. Completion Time</h3>
                <p class="stat-number">{{ dashboardStore.stats.avg_completion_time }}</p>
            </div>

        </div>
    </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { useDashboardStore } from '../../store/dashboard';

const dashboardStore = useDashboardStore();

onMounted(() => {
    dashboardStore.fetchStats();
});
</script>

<style scoped>
.stats-container {
    margin-bottom: 2rem;
}

.stats-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 1.5rem;
}

.stat-card {
    background-color: #fff;
    padding: 1.5rem;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
    border: 1px solid #e9ecef;
}

.stat-card h3 {
    margin-top: 0;
    font-size: 1rem;
    color: #6c757d;
    text-transform: uppercase;
}

.stat-number {
    font-size: 2.5rem;
    font-weight: bold;
    color: #343a40;
}

.stat-text {
    font-size: 1.2rem;
    font-weight: bold;
}

.sub-text {
    display: block;
    font-size: 0.9rem;
    color: #6c757d;
    font-weight: normal;
}

.stat-list {
    list-style: none;
    padding: 0;
}

.stat-list li {
    display: flex;
    justify-content: space-between;
    padding: 0.25rem 0;
}

.stat-list li span {
    font-weight: bold;
}

.error {
    color: #dc3545;
}
</style>