import { defineStore } from 'pinia';
import api from '../services/api';

export const useOrderStore = defineStore('order', {
    state: () => ({
        orders: [],
        isLoading: false,
        error: null,
    }),

    getters: {
        // Menghitung total order
        totalUserOrders: (state) => state.orders.length,

        // Menghitung order berdasarkan status
        userOrdersByStatus: (state) => {
            const stats = {
                Pending: 0,
                'In Progress': 0,
                Done: 0,
            };

            for (const order of state.orders) {
                if (stats.hasOwnProperty(order.status)) {
                    stats[order.status]++;
                }
            }
            return stats;
        },
    },

    actions: {
        /**
         * Mengambil order HANYA untuk user yang sedang login.
         * Digunakan di UserDashboard.
         */
        async fetchUserOrders() {
            this.isLoading = true;
            this.error = null;
            try {
                const response = await api.get('/orders');
                this.orders = response.data;
            } catch (error) {
                console.error('Failed to fetch user orders:', error);
                this.error = 'Failed to fetch your orders. Please try again later.';
            } finally {
                this.isLoading = false;
            }
        },

        /**
         * Membuat order baru.
         * Digunakan di halaman CreateOrder.
         * @param {object} orderData - Data order dari form { title, description, category }.
         */
        async createOrder(orderData) {
            this.isLoading = true;
            this.error = null;
            try {
                await api.post('/orders', orderData);
                // Arahkan kembali ke dashboard setelah berhasil
                this.router.push('/dashboard');
            } catch (error) {
                console.error('Failed to create order:', error.response?.data || error.message);
                this.error = 'Failed to create order. Please check your input and try again.';
                throw error; // Lemparkan error agar komponen bisa menangkapnya
            } finally {
                this.isLoading = false;
            }
        },

        /**
         * [ADMIN] Mengambil SEMUA order dari semua user.
         * Digunakan di AdminDashboard.
         */
        async fetchAllOrders() {
            this.isLoading = true;
            this.error = null;
            try {
                const response = await api.get('/admin/orders');
                this.orders = response.data;
            } catch (error) {
                console.error('Failed to fetch all orders:', error);
                this.error = 'Failed to fetch all orders.';
            } finally {
                this.isLoading = false;
            }
        },

        /**
         * [ADMIN] Mengubah status sebuah order.
         * Digunakan di AdminDashboard.
         * @param {number} orderId - ID dari order yang akan diubah.
         * @param {string} status - Status baru ('Pending', 'In Progress', 'Done').
         */
        async updateOrderStatus(orderId, status) {
            try {
                await api.put(`/admin/orders/${orderId}/status`, { status });

                // Optimistic UI update: langsung ubah data di state
                // agar UI terasa responsif tanpa perlu fetch ulang.
                const index = this.orders.findIndex(order => order.id === orderId);
                if (index !== -1) {
                    this.orders[index].status = status;
                }
            } catch (error) {
                console.error('Failed to update order status:', error);
                // Tampilkan pesan error ke admin
                alert('Failed to update status. Please check the console for details.');
            }
        },
    },
});