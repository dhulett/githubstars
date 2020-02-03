<template>
  <div></div>
</template>

<script>
export default {
  data() {
    return {};
  },
  async mounted() {
    const isAuthenticated = await this.$auth.isAuthenticated();
    isAuthenticated && this.$router.push('/me');
  },
  methods: {
    login () {
      this.$auth.loginRedirect('/me')
    }
  }
}
</script>
