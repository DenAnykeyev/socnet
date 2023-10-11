<!--
	Тут есть только шапка приложения

	Опишем тут только UX/UI плюшку, показывывающую, что пользователь
	вошел в аккаунт и может выйти по нажатию на эту кнопку...
-->
<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<template>
	<header>
		<nav class="navbar navbar-expand-lg navbar-light bg-light">
			<div class="container">
				<a class="navbar-brand fs-2 fw-bold" href="#">MyMessanger</a>
				<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
					aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
					<i class="fa fa-bars"></i>
				</button>
				<!--
				<div class="d-flex align-items-center ms-auto">
					<a class="nav-link" href="register.html">
						<i class="bi bi-person fs-4"></i>
					</a>
				</div>
				-->
				<div v-if="isAuth === true">
					<button class="btn btn-primary ms-auto bi bi-door-closed" @click="logout"> {{ this.name }}</button>
				</div>
			</div>
		</nav>
	</header>
	<main class="container">
		<RouterView />
	</main>
	<footer>
	</footer>
</template>
<script>
export default {
	data() {
		return {
			// Флаг регистрации пользователя
			isAuth: false,
			// Имя пользователя, если он зарегистрирован
			name: null
		};
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

		async logout() {
			const confirmDelete = confirm('Вы уверены, что хотите выйти из аккаунта?');
			if (confirmDelete) {
				try {
					const response = await fetch("/logout_user", {
						method: "POST",
					});

					if (!response.ok) {
						new Error("!!!")
					}
					this.isAuth = false;
					this.name = null;

					window.location.href = '/';
				} catch (error) {
					alert("Ошибка выхода из аккаунта", error);
				}
			}
		},
	},
};
</script>
<style scoped></style>