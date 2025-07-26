<template>
    <div class="status-tracker">
        <div class="step" :class="{ 'active': isPending, 'completed': isProgress || isDone }">
            <div class="icon">✔</div>
            <div class="label">Pending</div>
        </div>

        <div class="line" :class="{ 'completed': isProgress || isDone }"></div>

        <div class="step" :class="{ 'active': isProgress, 'completed': isDone }">
            <div class="icon">🔧</div>
            <div class="label">In Progress</div>
        </div>

        <div class="line" :class="{ 'completed': isDone }"></div>

        <div class="step" :class="{ 'active': isDone, 'completed': isDone }">
            <div class="icon">✅</div>
            <div class="label">Done</div>
        </div>
    </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
    status: {
        type: String,
        required: true,
        validator: (value) => ['Pending', 'In Progress', 'Done'].includes(value),
    },
});

const isPending = computed(() => props.status === 'Pending');
const isProgress = computed(() => props.status === 'In Progress');
const isDone = computed(() => props.status === 'Done');
</script>

<style scoped>
.status-tracker {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
    min-width: 300px;
}

.step {
    display: flex;
    flex-direction: column;
    align-items: center;
    color: #ccc;
    transition: color 0.4s ease;
    position: relative;
}

.icon {
    width: 30px;
    height: 30px;
    border-radius: 50%;
    background-color: #ccc;
    color: white;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 16px;
    border: 2px solid #ccc;
    transition: background-color 0.4s ease, border-color 0.4s ease;
}

.label {
    margin-top: 8px;
    font-size: 12px;
    font-weight: 500;
}

.line {
    flex-grow: 1;
    height: 4px;
    background-color: #ccc;
    margin: 0 10px;
    margin-bottom: 25px;
    transition: background-color 0.4s ease;
}

.step.completed {
    color: var(--color-primary);
}

.step.completed .icon {
    background-color: var(--color-primary);
    border-color: var(--color-primary);
}

.line.completed {
    background-color: var(--color-primary);
}

.step.active {
    color: var(--color-primary);
}

.step.active .icon {
    background-color: var(--color-background-card);
    border-color: var(--color-primary);
    color: var(--color-primary);
    transform: scale(1.1);
    transform: scale(1.1);
}
</style>