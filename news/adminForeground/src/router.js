import Vue from "vue"
import Router from "vue-router"
import Login from './views/login/index.vue'
import Home from './views/layout.vue'
import Index from './views/index/index.vue'
import We from './views/welcome'
import Find from './views/rib/find'
import Create from './views/rib/create'
Vue.use(Router)

const router=new Router({
	routes:[
		{path:'/',redirect:'/login'},
		{path:'/login',component:Login},
		{path:'/home',component:Home},
		{
			path:'/index',component:Index,
			redirect:'/welcome',			//重定向路由
			children:[
				{path:'/welcome',component:We},
				{path:'/find',component:Find},
				{path:'/create',component:Create}
			]
		}
	]
})



router.beforeEach((to,from,next)=>{
	if(to.path==='/login') return next();
	
	const tokenStr=window.sessionStorage.getItem('token')
	
	
	//获取token
	if(!tokenStr) return next('/login')
	next()
})

export default router