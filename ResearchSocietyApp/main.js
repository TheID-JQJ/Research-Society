import App from './App'

import Vue from 'vue'

//引入uview
import uView from '@/uni_modules/uview-ui'
Vue.use(uView)

//引入vuex
import store from '@/store'
Vue.prototype.store = store
import vuexStore from "@/store/$u.mixin.js"
Vue.mixin(vuexStore)

Vue.config.productionTip = false
App.mpType = 'app'
const app = new Vue({
    ...App,
	store
})
app.$mount()