<template>
    <div class="stats-container">
        <h3>My Order Summary</h3>
        <div v-if="orderStore.isLoading">Loading summary...</div>
        <div v-else-if="orderStore.orders.length > 0" class="stats-grid">

            <div class="stat-card">
                <h4>Total Requests</h4>
                <p class="stat-number">{{ orderStore.totalUserOrders }}</p>
            </div>

            <div class="stat-card">
                <h4>Pending</h4>
                <p class="stat-number status-pending-text">
                    {{ orderStore.userOrdersByStatus.Pending }}
                </p>
            </div>

            <div class="stat-card">
                <h4>In Progress</h4>
                <p class="stat-number status-progress-text">
                    {{ orderStore.userOrdersByStatus['In Progress'] }}
                </p>
            </div>

            <div class="stat-card">
                <h4>Completed</h4>
                <p class="stat-number status-done-text">
                    {{ orderStore.userOrdersByStatus.Done }}
                </p>
            </div>

        </div>
        <div v-else>
            <p>You have no orders to summarize yet.</p>
        </div>
    </div>
</template>

<script setup>
import { useOrderStore } from '../../store/order';

const orderStore = useOrderStore();
</script>

<style scoped>
.stats-container {
    margin-bottom: 2rem;
    background-color: #f8f9fa;
    padding: 1.5rem;
    border-radius: 8px;
}

.stats-container h3 {
    margin-top: 0;
}

.stats-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
    gap: 1rem;
}

.stat-card {
    background-color: #fff;
    text-align: center;
    padding: 1rem;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.stat-card h4 {
    margin: 0 0 0.5rem 0;
    font-size: 0.9rem;
    color: #6c757d;
    font-weight: normal;
}

.stat-number {
    font-size: 2rem;
    font-weight: bold;
    margin: 0;
}

.status-pending-text {
    color: #f0ad4e;
}

.status-progress-text {
    color: #17a2b8;
}

.status-done-text {
    color: #28a745;
}
</style>