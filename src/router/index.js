import { createRouter, createWebHistory } from 'vue-router'
import ProductsView from '../views/Products.vue'
import AboutView from '../views/About.vue'
import HomeView from '../views/Home.vue'
import RegisterView from '../views/Register.vue'
import LoginView from '../views/Login.vue'
import ChatView from '../views/Chat.vue'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			name: 'home',
			component: HomeView
		},
		{
			path: '/about',
			name: 'about',
			component: AboutView
		},
		{
			path: '/products',
			name: 'products',
			component: ProductsView
		},

		{
			path: '/register',
			name: 'register',
			component: RegisterView
		},
		{
			path: '/login',
			name: 'login',
			component: LoginView
		},
		{
			path: '/chats',
			name: 'chats',
			component: ChatView
		},
	]
})

export default router