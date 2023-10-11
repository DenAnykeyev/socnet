<script setup>
</script>

<template>
	<h1 class="text-center">Войти</h1>
	<form @submit.prevent="loginUser" class="my-form">
		<div class="form-group">
			<label for="numberPhone">Номер телефона:</label>
			<input type="text" id="numberPhone" v-model="numberPhone" class="form-control" required />
		</div>

		<div class="form-group">
			<label for="password">Пароль:</label>
			<input type="password" id="password" v-model="password" class="form-control" required />
		</div>
		<div class="text-center"><button type="submit" class="btn btn-primary">Отправить</button></div>

	</form>
</template>

<script>
export default {
	data() {
		return {

		};
	},
	mounted() {
	},
	methods: {
		async loginUser() {
			try {
				const response = await fetch('/login_user', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
					},
					body: JSON.stringify({
						numberPhone: this.numberPhone,
						password: this.password
					}),
				});

				if (!response.ok) {
					const errorText = await response.text();
					throw new Error(errorText);
				}
				const lol = await response.text()
				alert(lol)

				window.location.href = '/';
			} catch (error) {
				alert(error);

				throw error;
			}


		},
	},
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