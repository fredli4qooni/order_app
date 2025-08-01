import { createApp } from 'vue';
import { createPinia } from 'pinia';
import router from './router'; 
import App from './App.vue';
import './assets/main.css'

const app = createApp(App);
const pinia = createPinia();

pinia.use(({ store }) => {
  store.router = router;
});

app.use(pinia);
app.use(router); 

app.mount('#app');