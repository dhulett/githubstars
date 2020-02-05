<template>
	<div>
		<form v-on:submit="this.$router.push({ path: `/${user}` })">
        <label class="busca" for="busca">https://github.com/</label>
        <input class="busca-user" type="text" placeholder="username" >
        <button type="submit">get repositories</button>
		</form>
	</div>
</template>

<script>
export default {
	data() {
		return {
			user: ""
		};
	},
	methods: {
		requestAuthToken(router, user) {
			// Check if user is authenticated with github and get authentication token
			router.push(`/${user}/repos`);
		}
	}
};
</script>

<style scoped>
 input, button {
     padding: 10px;
    }

.busca{
    color: black;
}


 input {
     width: 30%;
    }

.busca-user{
    color: black;
    margin-top: 10px;

}

img{
    width: 10px;
}

input, button{
    padding: 5px;
    font-size: 17px;
}

form{
    width: 300px;
    margin: 0 auto;
    margin-top: 200px;
}

button{
    margin-top: 30px;
    margin-left: 40px;
    background-color:white;
    border-bottom-color:black;
    border-top-color:black;
    border-left-color:black;
}

</style>
