<template>
<div>
  <form>
    <input type="text" v-model="user" placeholder="search by tag" @keyup.enter="requestAuthToken"/>
  </form>
</div>
</template>

<script>
export default {
    data() {
      return {
        query: null,
      };
    },
    methods: {
      onSearchSubmition() {
        this.$emit('search-submitted', this.query);
      },
      async logout () {
        this.$router.push('/')
    }
  }
}
</script>
