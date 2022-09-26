const install = (Vue, vm) => {
	// 是否登录
	const isLogin=()=>{
		const token=vm.vuex_token;
		// 没有token，跳转登录页面
		if(!token){
			uni.showToast({
				title:"请先登录",
				icon:"error",
				mask:true
			});
			setTimeout(()=>{
				// 跳转登录页
				vm.$u.route({
					type: "reLaunch",
					url: "/pages/certific/login"
				})
			},1500);
			return false;
		}
		return true;
	}
	
	// 更新用户信息
	const setUserInfo= async()=>{
		// 更新vuex用户信息
		const userInfo = await vm.$u.api.getUserInfo();
		vm.$u.vuex('vuex_user',userInfo);
	};
	
	vm.$u.utils={
		isLogin,
		setUserInfo,
	};
}

export default {
	install
}