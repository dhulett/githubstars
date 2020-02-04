<template>
<div>

  <form>
    https://github.com/
    <input type="text" v-model="user" placeholder="username" @keyup.enter="requestAuthToken"/>
    <button onclick="requestAuthToken">get repositories</button>
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
  async mounted() {
    },
  methods: {
    requestAuthToken () {
      // Check if user is authenticated with github and get authentication token
     this.$router.push(`/${this.user}/repos`);
    }
  }
}
</script>
