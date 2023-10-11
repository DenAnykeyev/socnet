<!--
	Представление, использующееся для самого чатинга...
-->
<script setup>
</script>

<template>
	<div class="chat-app">
		<div class="sidebar">
			<input type="text" class="search-box" placeholder="Найти...">
			<div class="chat-list">
				<a href="#">Chat 1</a>
				<a href="#">Chat 2</a>
				<a href="#">Chat 3</a>
				<a href="#">Chat 4</a>
				<a href="#">Chat 5</a>
				<a href="#">Chat 6</a>
				<a href="#">Chat 7</a>
				<a href="#">Chat 8</a>
				<a href="#">Chat 8</a>
				<a href="#">Chat 8</a>
				<a href="#">Chat 8</a>
				<a href="#">Chat 8</a>
				<a href="#">Chat 8</a>
				<a href="#">Chat 8</a>
			</div>
		</div>

		<div class="content">
			<h2>Main Content</h2>
			<p>This is the main content area.</p>
		</div>
	</div>
</template>

<script>
export default {
	data() {
		return {
			isAuth: false
		}
	},
	mounted() {
		this.checkAuth();
	},
	methods: {
		async checkAuth() {
			try {
				const response = await fetch(`/get_user`, {
					method: 'GET',
				});

				if (!response.ok) {
					throw new Error('Не удалось получить данные о авторизации');
				}
				/*
					JSON ответ с backend сервера

					Должен содержать вид:
					{ 
						"id": 1, "firstName": "ДенисПупка", 
						"lastName": "ДенисПупкаЛупкаПупуп", 
						"numberPhone": "89963247891", 
						"password": "", 
						"rules": "user"
					  }
				*/
				const data = await response.json();

				if (data === null || data.Id === -1) {
					return;
				}
				this.isAuth = true
				this.name = data.firstName
			} catch (error) {
				alert(error)
			}
		},
		redirectToRegister() {
			window.location.href = '/register';

		},
		redirectToLogin() {
			window.location.href = '/login';

		},
		redirectToChats() {
			window.location.href = '/chats';
		}
	}
}
</script>
<style scoped>
body {
	font-family: Arial, sans-serif;
	margin: 0;
}

.chat-app {
	display: flex;
}

.sidebar {
	width: 25%;
	background-color: #8f9194;
	height: 100vh;
	overflow-y: auto;
	padding-top: 60px;
	position: fixed;
	left: 0;
}

.sidebar a {
	padding: 15px 25px;
	text-decoration: none;
	font-size: 18px;
	color: #f8f9fa;
	display: block;
}

.sidebar a:hover {
	background-color: #495057;
}


.search-box {
	padding: 10px 20px;
	margin: 20px 15px;
	background-color: #f8f9fa;
	border: none;
	border-radius: 5px;
	width: 80%;
	font-size: 16px;
}

.chat-list {
	padding: 20px 15px;
}

.content {
	flex: 1;
	padding: 20px;
	background-color: #fff;
	color: #000;
}
</style>