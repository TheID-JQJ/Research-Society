const install = (Vue, vm) => {
	const http = uni.$u.http
	
//菜谱相关
	//请求所有
	const recipes = (currentPage) => http.get("/recipes/all/"+currentPage);
	//按id查询
	const getRecipe = (rid) => http.get("/recipes/one/"+rid);
	//按关键词查询
	const searchRecipe = (keyword) => http.get("/recipes/search/"+keyword);

//商品相关
	//请求所有
	const goods = (currentPage) => http.get("/goods/all/"+currentPage);
	//按id查询
	const getOneGoods = (gid) => http.get("/goods/one/"+gid);
	//按关键词查询
	const searchGoods = (keyword) => http.get("/goods/search/"+keyword);
	//按种类查询
	const kindGoods = (kid) => http.get("/goods/kind/"+kid);
	//请求所有种类
	const kinds = (currentPage) => http.get("/kinds/all/");

//购物车相关
	//请求所有
	const carts = (currentPage) => http.post("/cart/all/"+currentPage);
	//添加到购物车
	const addCart = (gid) => http.post("/cart/add/"+gid);
	//删除
	const deleteCart = (params) => http.post("/cart/delete", params);

//订单相关
	//请求所有
	const orders = (currentPage) => http.post("/userorder/all/"+currentPage);
	//按id查询
	const order = (oid) => http.post("/orders/one/"+oid);
	//查询待付款
	const pendingPayment = () => http.post("/orders/pendingPayment");
	//查询待收货
	const received = () => http.post("/orders/received");
	//查询待评价
	const evaluated = () => http.post("/orders/evaluated");
	//创建订单
	const createOrder = (order) => http.post("/orders/create", order);
	//支付
	const pay = (oid) => http.post("/orders/pay", oid);
	//收货
	const delivery = (oid) => http.post("/orders/delivery", oid);
	//完成
	const success = (oid) => http.post("/orders/success", oid);
	
//用户相关
	//用户信息
	const getUserInfo = () => http.post("/users/getInfo");
	//修改用户信息
	const setUserInfo = (user) => http.post("/users/setInfo", user);
	//修改账号信息
	const setAccountInfo = (account) => http.post("/accounts/setInfo", account);
	//收藏的菜谱
	const likes = (currentPage) => http.post("/like/all/"+currentPage);
	//收藏菜谱
	const toLike = (rid) => http.post("/like/toLike/"+rid);
	//取消收藏菜谱
	const disLike = (rid) => http.post("/like/disLike/"+rid);
	//是否收藏菜谱
	const isLike = (rid) => http.post("/like/isLike/"+rid);
	//收藏的商品
	const stars = (currentPage) => http.post("/star/all/"+currentPage);
	//收藏商品
	const toStar = (gid) => http.post("/star/toStar/"+gid);
	//取消收藏商品
	const disStar = (gid) => http.post("/star/disStar/"+gid);
	//是否收藏商品
	const isStar = (gid) => http.post("/star/isStar/"+gid);

//VIP相关
	//查询是否是vip
	const isVip = () => http.post("/vip/isVip");
	//开通/续费vip
	const buyVip = (day) => http.post("/vip/buyVip/"+day);
	
//收货地址相关
	//请求所有
	const addresses = () => http.post("/address/all");
	//根据id查询
	const address = (aid) => http.post("/address/one/"+aid);
	//增
	const addAddress = (address) => http.post("/address/add", address);
	//删
	const deleteAddress = (oid) => http.post("/address/delete"+oid);
	//改
	const updateAddress = (address) => http.post("/address/update", address);
	
//认证相关
	//注册
	const register = (params ={}) => http.post("/accounts/register", params);
	//登录
	const login = (params ={}) => http.post("/accounts/login", params);
	//注销
	const logout = () => http.get("/accounts/logout");
	
//文件相关
	//获取图片
	// const picture = (name) => http.get("/picture/one/"+name);
	
	//挂载
	vm.$u.api={
		//菜谱相关
		recipes,
		getRecipe,
		searchRecipe,
		//商品相关
		goods,
		getOneGoods,
		searchGoods,
		kindGoods,
		kinds,
		//购物车相关
		carts,
		addCart,
		deleteCart,
		//订单相关
		orders,
		order,
		pendingPayment,
		received,
		evaluated,
		createOrder,
		pay,
		delivery,
		success,
		//用户相关
		getUserInfo,
		setUserInfo,
		setAccountInfo,
		likes,
		isLike,
		toLike,
		disLike,
		stars,
		toStar,
		disStar,
		isStar,
		//VIP相关
		isVip,
		buyVip,
		//收货地址相关
		addresses,
		address,
		addAddress,
		deleteAddress,
		updateAddress,
		//认证相关
		register,
		login,
		logout,	
	};
}

export default {
	install
}
