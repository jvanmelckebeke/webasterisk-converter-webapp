import 'primevue/resources/themes/lara-dark-teal/theme.css';

import {createApp} from 'vue'
import App from './App.vue'
import PrimeVue from 'primevue/config';

const app = createApp(App)
app.use(PrimeVue, {ripple: true})

app.mount('#app')
