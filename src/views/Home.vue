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
		<h1 class="text-center">Авторизирован</h1>
	</div>
	<div v-else>
		<!--
			Перекидываем в /register||login
		-->
		<h1 class="text-center">Не авторизирован</h1>
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
	}
}
</script>
