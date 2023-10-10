<script setup>
</script>

<template>
	<h1 class="text-center">Зарегистрироваться</h1>
	<form @submit.prevent="registerUser" class="my-form">
		<div class="form-group">
			<label for="numberPhone">Телефон:</label>
			<input type="text" id="numberPhone" v-model="numberPhone" class="form-control" required maxlength="11" />
		</div>

		<div class="form-group">
			<label for="password">Пароль:</label>
			<input type="password" id="password" v-model="password" class="form-control" required maxlength="20" />
		</div>

		<div class=" form-group">
			<label for="firstName">Имя:</label>
			<input type="text" id="firstName" v-model="firstName" class="form-control" required maxlength="15" />
		</div>

		<div class="form-group">
			<label for="lastName">Фамилия:</label>
			<input type="text" id="lastName" v-model="lastName" class="form-control" required maxlength="20" />
		</div>
		<div class="text-center"><button type="submit" class="btn btn-primary">Отправить</button></div>
	</form>
</template>
<script>
export default {
	data() {
		return {
			isLoggedIn: false,
		};
	},
	async mounted() {
		try {
			await this.checkAuth();

			if (this.isLoggedIn) {
				window.location.href = '/';
			}
		} catch (error) {
			console.error(error);
		}

	},
	methods: {
		// Метод для регистрации пользователя на сервере
		async registerUser() {
			const response = await fetch('/register_user', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify({
					firstName: this.firstName,
					lastName: this.lastName,
					numberPhone: this.numberPhone,
					password: this.password,
					rules: "user"
				}),
			});

			alert(this.firstName + "\n" + this.lastName + "\n" + this.numberPhone + "\n" + this.password)

			if (!response.ok) {
				const lol = await response.text()
				alert(lol)
				return;
			}

			alert(this.firstName + ", Вы были успешно зарегистрированы!")

			window.location.href = '/';
		},
	},
	redirectToLoginPage() {
		window.location.href = "/login";
	}
};
</script>
<style scoped>
.my-form {
	max-width: 400px;
	margin: 0 auto;
	padding: 20px;
	border: 1px solid #ccc;
	border-radius: 5px;
	background-color: #f9f9f9;
}

.form-group {
	margin-bottom: 20px;
}
</style>