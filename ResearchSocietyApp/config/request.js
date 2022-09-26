module.exports = (vm) => {
	
	
    // 初始化请求配置
    uni.$u.http.setConfig((config) => {
        /* config 为默认全局配置*/
		// config.baseURL = 'http://hkc.ink:8081/'
        config.baseURL = 'http://localhost:8081'
		config.dataType = 'json'
		config.showLoading =  true // 显示请求中的loading
		config.loadingText =  '请求中...' // 请求loading中的文字提示
		config.loadingTime =  80000 // 在此时间内，请求还没回来的话，就显示加载中动画，单位ms
		config.originalData =  true // 在拦截器中返回服务端的原始数据
		config.loadingMask =  true // 展示loading的时候，给一个透明的蒙层，防止触摸穿透
        return config
    })
	
	// 请求拦截
	uni.$u.http.interceptors.request.use((config) => { // 可使用async await 做异步操作
		config.header.Authorization = vm.$store.state.vuex_token
	    return config 
	})
	
	// 响应拦截
	uni.$u.http.interceptors.response.use((res) => { /* 对响应成功做点什么 可使用async await 做异步操作*/
		const {
			code,
			data,
			message
		} = res.data
		
		console.log(res)
		
		if (code < 400) {
			//正常则返回数据
			return data
		} else if (code == 400) {
			// 错误的请求
			vm.$u.toast(message)
			return false
		} else if (code == 401) {
			// token过期，未登录权限
			uni.showToast({
				title:"请重新登录",
				icon:"error",
				mask:true
			})
			setTimeout(()=>{
				// 跳转登录页
				vm.$u.route({
					url:'/pages/login/login'
				})
			},1500)
			return false
		} else {
			vm.$u.toast("请求超时")
			return false
		}
	}, (err) => {
		const {
			code,
			data,
			message
		} = err.data
		
		// console.log(code)
		
		if (code == 400) {
			// 错误的请求
			vm.$u.toast(message)
			return false
		} else if (code == 401) {
			// token过期，未登录权限
			uni.showToast({
				title:"请重新登录",
				icon:"error",
				mask:true
			})
			setTimeout(()=>{
				// 跳转登录页
				vm.$u.route({
					url:'/pages/login/login'
				})
			},1500)
			return false
		} else {
			vm.$u.toast("请求错误")
			return false
		}
	})
}