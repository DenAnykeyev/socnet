<!--
	Представление, использующееся для редиректа!
	Если пользователь есть в списке сессий - /messanger
	Если пользователя нет в списке сессий - /auth
-->
<script setup>
</script>

<template>
	<div v-if="isAuth === true">
		<!--
			Перекидываем на /chats
		-->
		<div class="text-center">
			<h1>Авторизирован</h1>
			<button class="btn btn-primary ms-auto" @click="redirectToChats"> Перейти к чатам</button>
		</div>
	</div>
	<div v-else>
		<!--
			Перекидываем в /register||login
		-->
		<div class="text-center">
			<h1>Не авторизирован</h1>
			<button class="btn btn-primary ms-auto" @click="redirectToRegister"> Перейти к регистрации</button>
			<button class="btn btn-primary ms-auto" @click="redirectToLogin"> Перейти к авторизации</button>
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
