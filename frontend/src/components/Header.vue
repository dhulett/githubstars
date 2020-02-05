<template>
	<div class="header-container">
		<div class="title">{{ applicationTitle }}</div>
		<div class="home" v-if="displayHomeButton">
			<router-link to="/">{{ homeButtonTitle }}</router-link>
		</div>
	</div>
</template>

<script>
export default {
	data() {
		return {
			displayHomeButton: true
		};
	},
	beforeRouteUpdate(to, from, next) {
        this.displayHomeButton = to != "/";
		next();
	},
	props: {
		applicationTitle: String,
		homeButtonTitle: String
	}
};
</script>

<style scoped>
div {
	position: relative;
	display: block;
}

.header-container {
    width: 100%;
    height: 50px;
}

.title {
	float: left;
}

.home {
	float: right;
}
</style>
