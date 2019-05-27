// Make sure to import from 'blessed-vue' instead of 'vue'!
import Vue from 'blessed-vue';
import App from './App.vue';

// Create a new fake root element. (Since blessed has no real DOM)
const el = Vue.dom.createElement();
Vue.dom.append(el);

new Vue({
  render: h => h(App)
})
// Mount to the fake element.
.$mount(el);